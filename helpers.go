package main

import (
	"fmt"
	"net/http"
	"path/filepath"
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

func GetFileName(f string) string {
	name := strings.TrimSuffix(f, filepath.Ext(f))
	if len(name) > 55 {
		name = name[:55] + "..."
	}
	return name
}

func GetFileType(f string) string {
	if strings.HasSuffix(f, ".pdf") {
		return "PDF/" + filepath.Ext(f)
	} else if strings.HasSuffix(f, ".doc") || strings.HasSuffix(f, ".docx") {
		return "Wordt/" + filepath.Ext(f)
	} else if strings.HasSuffix(f, ".xls") || strings.HasSuffix(f, ".xlsx") {
		return "Excel/" + filepath.Ext(f)
	} else if strings.HasSuffix(f, ".ppt") || strings.HasSuffix(f, ".pptx") {
		return "PowerPoint/" + filepath.Ext(f)
	} else if strings.HasSuffix(f, ".zip") || strings.HasSuffix(f, ".rar") {
		return "Archive/" + filepath.Ext(f)
	} else if strings.HasSuffix(f, ".txt") {
		return "Text/" + filepath.Ext(f)
	} else if strings.HasSuffix(f, ".mp3") || strings.HasSuffix(f, ".wav") || strings.HasSuffix(f, ".ogg") {
		return "Audio/" + filepath.Ext(f)
	} else if strings.HasSuffix(f, ".mp4") || strings.HasSuffix(f, ".avi") || strings.HasSuffix(f, ".mkv") {
		return "Video/" + filepath.Ext(f)
	} else if strings.HasSuffix(f, ".png") || strings.HasSuffix(f, ".jpg") || strings.HasSuffix(f, ".jpeg") || strings.HasSuffix(f, ".gif") {
		return "Image/" + filepath.Ext(f)
	} else if strings.HasSuffix(f, ".exe") {
		return "Executable/" + filepath.Ext(f)
	} else if strings.HasSuffix(f, ".iso") {
		return "Disk Image/" + filepath.Ext(f)
	} else if strings.HasSuffix(f, ".apk") {
		return "Android App/" + filepath.Ext(f)
	} else if strings.HasSuffix(f, ".py") {
		return "Python/" + filepath.Ext(f)
	} else if strings.HasSuffix(f, ".go") {
		return "Go/" + filepath.Ext(f)
	} else if strings.HasSuffix(f, ".cpp") {
		return "C++/" + filepath.Ext(f)
	} else if strings.HasSuffix(f, ".java") {
		return "Java/" + filepath.Ext(f)
	} else if strings.HasSuffix(f, ".c") {
		return "C/" + filepath.Ext(f)
	} else if strings.HasSuffix(f, ".html") || strings.HasSuffix(f, ".htm") {
		return "HTML/" + filepath.Ext(f)
	} else {
		return "Unknown/" + filepath.Ext(f)
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
