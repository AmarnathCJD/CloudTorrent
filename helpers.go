package main

import (
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
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Size       string `json:"size,omitempty"`
	Type       string `json:"type,omitempty"`
	Color      string `json:"color,omitempty"`
	Path       string `json:"path,omitempty"`
	IsDir      string `json:"is_dir,omitempty"`
	Ext        string `json:"ext,omitempty"`
	StreamLink string `json:"stream,omitempty"`
	Class      string `json:"class,omitempty"`
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

func GetFileType(f string) (string, string, string) {
	f = strings.ToLower(f)
	if strings.HasSuffix(f, ".mp4") || strings.HasSuffix(f, ".avi") || strings.HasSuffix(f, ".mkv") || strings.HasSuffix(f, ".webm") {
		return "Video", "bi bi-file-earmark-play", "blue"
	} else if strings.HasSuffix(f, ".mp3") || strings.HasSuffix(f, ".wav") || strings.HasSuffix(f, ".flac") {
		return "Audio", "bi bi-file-earmark-music", "green"
	} else if strings.HasSuffix(f, ".jpg") || strings.HasSuffix(f, ".png") || strings.HasSuffix(f, ".gif") || strings.HasSuffix(f, ".webp") {
		return "Image", "bi bi-image", "orange"
	} else if strings.HasSuffix(f, ".pdf") {
		return "Pdf", "bi bi-filetype-pdf", "red"
	} else if strings.HasSuffix(f, ".txt") {
		return "Text", "bi bi-journal-text", "purple"
	} else if strings.HasSuffix(f, ".zip") || strings.HasSuffix(f, ".rar") || strings.HasSuffix(f, ".7z") {
		return "Archive", "bi bi-file-earmark-zip", "brown"
	} else if strings.HasSuffix(f, ".iso") {
		return "Iso", "bi bi-disc", "brown"
	} else if strings.HasSuffix(f, ".exe") {
		return "Exe", "bi bi-filetype-exe", "red"
	} else if strings.HasSuffix(f, ".doc") || strings.HasSuffix(f, ".docx") {
		return "Doc", "bi bi-file-word", "red"
	} else if strings.HasSuffix(f, ".xls") || strings.HasSuffix(f, ".xlsx") {
		return "Xls", "bi bi-file-earmark-excel", "green"
	} else if strings.HasSuffix(f, ".ppt") || strings.HasSuffix(f, ".pptx") {
		return "Ppt", "bi bi-filetype-pptx", "orange"
	} else if strings.HasSuffix(f, ".torrent") {
		return "Torrent", "bi bi-magnet", "green"
	} else if strings.HasSuffix(f, ".py") {
		return "Python", "bi bi-filetype-py", "blue"
	} else if strings.HasSuffix(f, ".go") {
		return "Go", "bi bi-filetype-go", "blue"
	} else if strings.HasSuffix(f, ".js") {
		return "Js", "bi bi-filetype-js", "blue"
	} else if strings.HasSuffix(f, ".json") {
		return "Json", "bi bi-filetype-json", "blue"
	} else if strings.HasSuffix(f, ".html") {
		return "Html", "bi bi-filetype-html", "green"
	} else if strings.HasSuffix(f, ".css") {
		return "Css", "bi bi-filetype-css", "blue"
	} else if strings.HasSuffix(f, ".db") {
		return "Db", "bi bi-box2", "blue"
	} else {
		return "Other", "bi bi-question", "black"
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

func SortAlpha(sortData []TorrentMeta) []TorrentMeta {
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
		var Size, Type, Icon, Color, Ext, StreamURL string
		var Path = filepath.Join(path, file.Name())
		Path = "/downloads" + ServerPath(path+"/"+file.Name())
		if file.IsDir() {
			Type = "Folder"
			StreamURL = ServerPath(path + "/" + file.Name())
			DirSize, _ := DirSize(path + "/" + file.Name())
			Size = ByteCountSI(DirSize)
		} else {
			Size = ByteCountSI(file.Size())
			Type, Icon, Color = GetFileType(file.Name())
			Ext = filepath.Ext(file.Name())
			StreamURL = AbsPath("/stream" + strings.Replace(Path, "downloads", "dir", 1))
		}
		files = append(files, FileInfo{
			ID:         strconv.Itoa(i),
			Name:       GetName(file.Name()),
			Size:       Size,
			Type:       Type,
			Path:       Path,
			Color:      Color,
			IsDir:      strconv.FormatBool(file.IsDir()),
			Ext:        Ext,
			StreamLink: StreamURL,
			Class:      Icon,
		})
	}
	return files, nil

}

func AbsPath(path string) string {
	if runtime.GOOS == "windows" {
		return filepath.ToSlash(path)
	}
	return path
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
