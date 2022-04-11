package main

import (
	"fmt"
	"net/http"
	"path/filepath"
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
	if len(name) > 20 {
		name = name[:20] + "..."
	}
	return name
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

type SortBy func(p1, p2 *TorrentMeta) bool

func (b SortBy) Sort (t []TorrentMeta) {

}

type TorrentSorter struct {
Torr []TorrentMeta
by SortBy
}

func (s *TorrentSorter) Len() int {
	return len(s.Torr)
}

func (s *TorrentSorter) Swap(i, j int) {
	s.Torr[i], s.Torr[j] = s.Torr[j], s.Torr[i]
}

func (s *TorrentSorter) Less(i, j int) bool {
	return s.by(&s.Torr[i], &s.Torr[j])
}

func SortTorrent(t TorrentsResponse) TorrentsResponse {
name := func(p1, p2 *TorrentMeta) bool {
return p1.Name < p2.Name
}
data := t.Torrents
SortBy(name).Sort(data)
return TorrentsResponse{Torrents: data}
}

//xo
