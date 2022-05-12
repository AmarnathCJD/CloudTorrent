package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/disk"
)

type DiskStatus struct {
	All, Used, Free string
}

type FileInfo struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Size  string `json:"size,omitempty"`
	Type  string `json:"type,omitempty"`
	Path  string `json:"path,omitempty"`
	IsDir string `json:"is_dir,omitempty"`
	Ext   string `json:"ext,omitempty"`
	Icon  string `json:"icon,omitempty"`
}

type TopTorr struct {
	Name     string  `json:"name,omitempty"`
	Size     float64 `json:"size,omitempty"`
	Seeders  float64 `json:"seeders,omitempty"`
	Leechers float64 `json:"leechers,omitempty"`
	InfoHash string  `json:"info_hash,omitempty"`
}

type SearchReq struct {
	Name     string `json:"name"`
	InfoHash string `json:"info_hash"`
	Leechers string `json:"leechers"`
	Seeders  string `json:"seeders"`
	Size     string `json:"size"`
	Magnet   string `json:"magnet"`
}

type Handle struct {
	Path string
	Func func(http.ResponseWriter, *http.Request)
}

func DiskUsage(path string) DiskStatus {
	data, _ := disk.Usage(path)
	fs := DiskStatus{
		All:  ByteCountSI(int64(data.Total)),
		Used: ByteCountSI(int64(data.Used)),
		Free: ByteCountSI(int64(data.Free)),
	}
	return fs
}

func MemUsage() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return ByteCountSI(int64(m.Alloc))
}

func GetName(f string) string {
	return strings.TrimSuffix(f, filepath.Ext(f))
}

func GetFileType(f string) (string, string) {
	f = strings.ToLower(f)
	if strings.HasSuffix(f, ".mp4") || strings.HasSuffix(f, ".avi") || strings.HasSuffix(f, ".mkv") || strings.HasSuffix(f, ".webm") {
		return "Video", "fluency/48/000000/video.png"
	} else if strings.HasSuffix(f, ".mp3") || strings.HasSuffix(f, ".wav") || strings.HasSuffix(f, ".flac") {
		return "Audio", "nolan/64/musical.png"
	} else if strings.HasSuffix(f, ".jpg") || strings.HasSuffix(f, ".png") || strings.HasSuffix(f, ".gif") || strings.HasSuffix(f, ".webp") {
		return "Image", "color-glass/48/000000/image.png"
	} else if strings.HasSuffix(f, ".pdf") {
		return "Pdf", "color/48/000000/pdf.png"
	} else if strings.HasSuffix(f, ".txt") {
		return "Text", "external-prettycons-flat-prettycons/47/000000/external-text-text-formatting-prettycons-flat-prettycons-1.png"
	} else if strings.HasSuffix(f, ".zip") || strings.HasSuffix(f, ".rar") || strings.HasSuffix(f, ".7z") {
		return "Archive", "external-gradients-pongsakorn-tan/64/000000/external-archive-file-and-document-gradients-pongsakorn-tan-4.png"
	} else if strings.HasSuffix(f, ".iso") {
		return "ISO", "external-justicon-lineal-color-justicon/64/000000/external-iso-file-file-type-justicon-lineal-color-justicon.png"
	} else if strings.HasSuffix(f, ".exe") {
		return "Exe", "bi bi-filetype-exe"
	} else if strings.HasSuffix(f, ".doc") || strings.HasSuffix(f, ".docx") {
		return "Doc", "bi bi-file-word"
	} else if strings.HasSuffix(f, ".xls") || strings.HasSuffix(f, ".xlsx") {
		return "Xls", "bi bi-file-earmark-excel"
	} else if strings.HasSuffix(f, ".ppt") || strings.HasSuffix(f, ".pptx") {
		return "Ppt", "bi bi-filetype-pptx"
	} else if strings.HasSuffix(f, ".torrent") {
		return "Torrent", "fluency/48/000000/utorrent.png"
	} else if strings.HasSuffix(f, ".py") {
		return "Python", "color/48/000000/python--v1.png"
	} else if strings.HasSuffix(f, ".go") {
		return "Go", "color/48/000000/golang.png"
	} else if strings.HasSuffix(f, ".js") {
		return "Js", "color/48/000000/javascript--v1.png"
	} else if strings.HasSuffix(f, ".json") {
		return "JSON", "bi bi-filetype-json"
	} else if strings.HasSuffix(f, ".html") {
		return "HTML", "color/48/000000/html-5--v1.png"
	} else if strings.HasSuffix(f, ".css") {
		return "CSS", "external-flaticons-flat-flat-icons/64/000000/external-css-web-development-flaticons-flat-flat-icons.png"
	} else if strings.HasSuffix(f, ".db") {
		return "Database", "color/48/000000/data-configuration.png"
	} else {
		return "Unknown", "bi bi-file-earmark"
	}
}

func DeleteFile(path string) error {
	if f, err := os.Stat(path); os.IsNotExist(err) {
		return err
	} else if f.IsDir() {
		return os.RemoveAll(path)
	} else {
		return os.Remove(path)
	}
}

func ByteCountSI(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func SortAlpha(sortData []TorrentData) []TorrentData {
	var data = sortData
	sort.Slice(data, func(p, q int) bool {
		return data[p].Name < data[q].Name
	})
	return data
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func StringToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func GetOutboundPort() string {
	if p := os.Getenv("PORT"); p != "" {
		if !strings.HasPrefix(p, ":") {
			return ":" + p
		} else {
			return p
		}
	}
	return ":80"
}

func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}

func GetDirContentsMap(path string) ([]FileInfo, error) {
	var files []FileInfo
	var DirWalk, err = ioutil.ReadDir(path)
	if err != nil {
		return files, err
	}
	for i, file := range DirWalk {
		var Size, Type, Ext, Icon string
		if file.IsDir() {
			Type = "Folder"
			DirSize, _ := DirSize(path + "/" + file.Name())
			Size = ByteCountSI(DirSize)
		} else {
			Size = ByteCountSI(file.Size())
			Type, Icon = GetFileType(file.Name())
			Ext = filepath.Ext(file.Name())
		}
		files = append(files, FileInfo{
			ID:    strconv.Itoa(i),
			Name:  GetName(file.Name()),
			Size:  Size,
			Type:  Type,
			Path:  GetPath(path, file),
			IsDir: strconv.FormatBool(file.IsDir()),
			Ext:   Ext,
			Icon:  Icon,
		})
	}
	return files, nil

}

func AbsPath(path string) string {
	return filepath.ToSlash(path)
}

func ServerPath(path string) string {
	return strings.Replace(AbsPath(path), AbsPath(Root), "", 1)
}

func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

func PrepareWD() {
	if _, err := os.Stat(filepath.Join(Root, "torrents")); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Join(Root, "torrents"), 0755); err != nil {
			panic(err)
		}
	}
}

func GetPath(path string, file os.FileInfo) string {
	if file.IsDir() {
		return "/downloads" + ServerPath(path+"/"+file.Name())
	} else {
		return "/dir" + ServerPath(path+"/"+file.Name())
	}
}

func ZipDir(path string, torrName string) (string, error) {
	path = path + "/"
	var zipPath = filepath.Join(Root, "torrents", torrName+".zip")
	if _, err := os.Stat(zipPath); !os.IsNotExist(err) {
		return zipPath, err
	}
	var zipFile, _ = os.Create(zipPath)
	defer zipFile.Close()
	var zipWriter = zip.NewWriter(zipFile)
	defer zipWriter.Close()
	addFiles(zipWriter, path, "")
	return zipPath, nil
}

func addFiles(w *zip.Writer, basePath, baseInZip string) {
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			dat, err := ioutil.ReadFile(basePath + file.Name())
			if err != nil {
				fmt.Println(err)
			}

			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				fmt.Println(err)
			}
			_, err = f.Write(dat)
			if err != nil {
				fmt.Println(err)
			}
		} else if file.IsDir() {

			newBase := basePath + file.Name() + "/"
			addFiles(w, newBase, baseInZip+file.Name()+"/")
		}
	}
}
