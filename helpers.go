package main

import (
	"fmt"
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

type Torrent struct {
	Name   string
	Size   string
	Date   string
	Magnet string
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

func GetFileName(f string) string {
	name := strings.TrimSuffix(f, filepath.Ext(f))
	if len(name) > 45 {
		name = name[:45] + "..."
	}
	return name
}

func GetDirName(name string) string {
	if len(name) > 45 {
		name = name[:45] + "..."
	}
	return name
}

func GetFileType(f string) (string, string, string) {
	f = strings.ToLower(f)
	if strings.HasSuffix(f, ".mp4") || strings.HasSuffix(f, ".avi") || strings.HasSuffix(f, ".mkv") {
		return "video", "bi bi-file-earmark-play", "blue"
	} else if strings.HasSuffix(f, ".mp3") || strings.HasSuffix(f, ".wav") || strings.HasSuffix(f, ".flac") {
		return "audio", "bi bi-file-earmark-music", "green"
	} else if strings.HasSuffix(f, ".jpg") || strings.HasSuffix(f, ".png") || strings.HasSuffix(f, ".gif") {
		return "image", "bi bi-image", "orange"
	} else if strings.HasSuffix(f, ".pdf") {
		return "pdf", "bi bi-filetype-pdf", "red"
	} else if strings.HasSuffix(f, ".txt") {
		return "text", "bi bi-journal-text", "purple"
	} else if strings.HasSuffix(f, ".zip") || strings.HasSuffix(f, ".rar") || strings.HasSuffix(f, ".7z") {
		return "archive", "bi bi-file-earmark-zip", "brown"
	} else if strings.HasSuffix(f, ".iso") {
		return "iso", "bi bi-disc", "brown"
	} else if strings.HasSuffix(f, ".exe") {
		return "exe", "bi bi-filetype-exe", "red"
	} else if strings.HasSuffix(f, ".doc") || strings.HasSuffix(f, ".docx") {
		return "doc", "bi bi-file-word", "red"
	} else if strings.HasSuffix(f, ".xls") || strings.HasSuffix(f, ".xlsx") {
		return "xls", "bi bi-file-earmark-excel", "green"
	} else if strings.HasSuffix(f, ".ppt") || strings.HasSuffix(f, ".pptx") {
		return "ppt", "bi bi-filetype-pptx", "orange"
	} else if strings.HasSuffix(f, ".torrent") {
		return "torrent", "bi bi-magnet", "green"
	} else if strings.HasSuffix(f, ".py") {
		return "python", "bi bi-filetype-py", "blue"
	} else if strings.HasSuffix(f, ".go") {
		return "go", "bi bi-filetype-go", "blue"
	} else if strings.HasSuffix(f, ".js") {
		return "js", "bi bi-filetype-js", "blue"
	} else if strings.HasSuffix(f, ".json") {
		return "json", "bi bi-filetype-json", "blue"
	} else if strings.HasSuffix(f, ".html") {
		return "html", "bi bi-filetype-html", "green"
	} else if strings.HasSuffix(f, ".css") {
		return "css", "bi bi-filetype-css", "blue"
	} else {
		return "other", "bi bi-question", "black"
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

func SortAlpha(t TorrentsResponse) TorrentsResponse {
	data := t.Torrents
	sort.Slice(data, func(p, q int) bool {
		return data[p].Name < data[q].Name
	})
	return TorrentsResponse{Torrents: data}
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

func PORT() string {
	if p := os.Getenv("PORT"); p != "" {
		return p
	}
	return "80"
}

func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}
