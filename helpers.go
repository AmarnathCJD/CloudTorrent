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
	if len(name) > 45 {
		name = name[:45] + "..."
	}
	return name
}

func GetFileType(f string) (string, string, string) {
	if strings.HasSuffix(f, ".mp4") || strings.HasSuffix(f, ".avi") || strings.HasSuffix(f, ".mkv") {
		return "video", "bi bi-filetype-mov", "blue"
	} else if strings.HasSuffix(f, ".mp3") || strings.HasSuffix(f, ".wav") || strings.HasSuffix(f, ".flac") {
		return "audio", "fa fa-file-audio-o", "green"
	} else if strings.HasSuffix(f, ".jpg") || strings.HasSuffix(f, ".png") || strings.HasSuffix(f, ".gif") {
		return "image", "fa fa-file-picture-o", "orange"
	} else if strings.HasSuffix(f, ".pdf") {
		return "pdf", "bi bi-filetype-pdf", "red"
	} else if strings.HasSuffix(f, ".txt") {
		return "text", "fa fa-file-text-o", "purple"
	} else if strings.HasSuffix(f, ".zip") || strings.HasSuffix(f, ".rar") || strings.HasSuffix(f, ".7z") {
		return "archive", "fa fa-file-archive-o", "brown"
	} else if strings.HasSuffix(f, ".iso") {
		return "iso", "fa fa-file-archive-o", "brown"
	} else if strings.HasSuffix(f, ".exe") {
		return "exe", "fa fa-file-code-o", "red"
	} else if strings.HasSuffix(f, ".doc") || strings.HasSuffix(f, ".docx") {
		return "doc", "fa fa-file-word-o", "red"
	} else if strings.HasSuffix(f, ".xls") || strings.HasSuffix(f, ".xlsx") {
		return "xls", "fa fa-file-excel-o", "green"
	} else if strings.HasSuffix(f, ".ppt") || strings.HasSuffix(f, ".pptx") {
		return "ppt", "fa fa-file-powerpoint-o", "orange"
	} else if strings.HasSuffix(f, ".torrent") {
		return "torrent", "fa fa-file-archive-o", "green"
	} else {
		return "other", "fa fa-file-o", "blue"
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
