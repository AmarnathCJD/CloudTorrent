package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	root = "C:/Users/rj/Downloads"
	// dir to server
)

func main() {
	http.HandleFunc("/downloads/", File)
	http.ListenAndServe(":80", nil)
}

func File(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(root, r.URL.Path)
	stat, err := os.Stat(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if stat.IsDir() {
		serveDir(w, r, path)
		return
	}
	http.ServeFile(w, r, path)
}
func serveDir(w http.ResponseWriter, r *http.Request, path string) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	files, err := file.Readdir(-1)
	if err != nil {
		fmt.Println(err)
	}
	var list = make(map[string]os.FileInfo)
	for _, file := range files {
		list[file.Name()] = file
	}
	IP := GetIP(r)
	w.Write([]byte(GetHTMLDir(list, IP)))
}

func GetHTMLDir(f map[string]os.FileInfo, IP string) string {
	var html = downloads
	var files = ""
	tbl := `<tr> <td>{{name}}</td> <td>{{size}}</td> <td>{{type}}</td> <td>{{date}}</td> </tr>`
	for _, v := range f {
		files += tbl
		Size := ByteCountSI(int64(v.Size()))
		FileType := GetFileType(v.Name())
		if v.IsDir() {
			FileType = "Folder"
			Size = "-"
		}
		files = strings.Replace(files, "{{name}}", "<a href=\"/downloads/"+v.Name()+"\">"+strings.Title(v.Name())+"</a>", -1)
		files = strings.Replace(files, "{{size}}", Size, -1)
		files = strings.Replace(files, "{{type}}", FileType, -1)
		files = strings.Replace(files, "{{date}}", strings.ReplaceAll(v.ModTime().String(), "+0000 GMT", ""), -1)
	}
	html = strings.Replace(html, "{{files}}", files, -1)
	html = strings.Replace(html, "{{ip}}", IP, -1)
	return html
}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
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

func GetFileType(f string) string {
	if strings.HasSuffix(f, ".pdf") {
		return "PDF/Document"
	} else if strings.HasSuffix(f, ".doc") || strings.HasSuffix(f, ".docx") {
		return "Word/Document"
	} else if strings.HasSuffix(f, ".xls") || strings.HasSuffix(f, ".xlsx") {
		return "Excel/Document"
	} else if strings.HasSuffix(f, ".ppt") || strings.HasSuffix(f, ".pptx") {
		return "PowerPoint/Document"
	} else if strings.HasSuffix(f, ".zip") || strings.HasSuffix(f, ".rar") {
		return "Archive/Compressed"
	} else if strings.HasSuffix(f, ".txt") {
		return "Text/Document"
	} else if strings.HasSuffix(f, ".mp3") || strings.HasSuffix(f, ".wav") || strings.HasSuffix(f, ".ogg") {
		return "Audio/Music"
	} else if strings.HasSuffix(f, ".mp4") || strings.HasSuffix(f, ".avi") || strings.HasSuffix(f, ".mkv") {
		return "Video/Movie"
	} else if strings.HasSuffix(f, ".png") || strings.HasSuffix(f, ".jpg") || strings.HasSuffix(f, ".jpeg") || strings.HasSuffix(f, ".gif") {
		return "Image/Photo"
	} else if strings.HasSuffix(f, ".exe") {
		return "Executable/Program"
	} else if strings.HasSuffix(f, ".iso") {
		return "Disk Image/ISO"
	} else if strings.HasSuffix(f, ".apk") {
		return "Android App/Program"
	} else if strings.HasSuffix(f, ".py") {
		return "Python/Program"
	} else if strings.HasSuffix(f, ".go") {
		return "Go/Program"
	} else if strings.HasSuffix(f, ".cpp") {
		return "C++/Program"
	} else if strings.HasSuffix(f, ".java") {
		return "Java/Program"
	} else if strings.HasSuffix(f, ".c") {
		return "C/Program"
	} else if strings.HasSuffix(f, ".html") || strings.HasSuffix(f, ".htm") {
		return "HTML/Document"
	} else {
		return "Unknown"
	}
}
