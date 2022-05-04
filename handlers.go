package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/sse"
)

var (
	SSEFeed = sse.New()
)

func AddTorrent(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	r.ParseForm()
	magnet := r.FormValue("magnet")
	if magnet == "" {
		http.Error(w, "No magnet provided", http.StatusBadRequest)
		return
	}
	if ok, err := AddTorrentByMagnet(magnet); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if !ok {
		http.Error(w, "Torrent already exists", http.StatusBadRequest)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func DeleteTorrent(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	r.ParseForm()
	id := r.FormValue("uid")
	if id == "" {
		http.Error(w, "No uid provided", http.StatusBadRequest)
		return
	}
	if ok, err := DeleteTorrentByID(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if !ok {
		http.Error(w, "Torrent not found", http.StatusNotFound)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func PauseTorrent(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	r.ParseForm()
	id := r.FormValue("uid")
	if id == "" {
		http.Error(w, "No uid provided", http.StatusBadRequest)
		return
	}
	PauseTorrentByID(id)
	w.WriteHeader(http.StatusOK)
}

func ResumeTorrent(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	r.ParseForm()
	id := r.FormValue("uid")
	if id == "" {
		http.Error(w, "No uid provided", http.StatusBadRequest)
		return
	}
	ResumeTorrentByID(id)
	w.WriteHeader(http.StatusOK)
}

func DropAll(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	DropAllTorrents()
	w.WriteHeader(http.StatusOK)
}

func StartAllHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	StartAll()
	w.WriteHeader(http.StatusOK)
}

func StopAllHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	StopAll()
	w.WriteHeader(http.StatusOK)
}

func SystemStats(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	Disk := DiskUsage(Root)
	SystemStat := "<p><b>IP:</b> " + GetIP(r) + " " + "<b>OS:</b> " + runtime.GOOS + " " + "<b>Arch:</b> " + runtime.GOARCH + " " + "<b>CPU:</b> " + fmt.Sprint(runtime.NumCPU()) + " " + "<b>RAM:</b> " + MemUsage() + " " + "<b>Disk:</b> " + fmt.Sprintf("%s/%s", Disk.Used, Disk.All) + " " + "<b>Downloads:</b> " + strconv.Itoa(GetLenTorrents()) + "</p>"
	w.Write([]byte(SystemStat))
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	path := strings.Replace(AbsPath(filepath.Join(Root, r.URL.Path)), "/delete", "", 1)
	fmt.Println(path)
	if err := os.Remove(path); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	log.Printf("Uploaded file: %+v\n", handler.Filename)
	log.Printf("File size: %+v\n", handler.Size)
	log.Printf("MIME header: %+v\n", handler.Header)
	DirPath := strings.Replace(filepath.Join(Root, r.FormValue("path")), "/downloads", "", 1)
	f, err := os.OpenFile(filepath.Join(DirPath, handler.Filename), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	w.WriteHeader(http.StatusOK)
}

func CreateFolderHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	r.ParseForm()
	DirPath := strings.Replace(filepath.Join(Root, r.URL.Path), "/api/create/downloads", "", 1)
	if err := os.MkdirAll(DirPath, 0777); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func streamTorrentUpdate() {
	fmt.Println("Streaming Torrents started")
	for range time.Tick(time.Millisecond * 600) {
		TORRENTS := GetAllTorrents()
		d, _ := json.Marshal(TORRENTS)
		SSEFeed.SendString("", "torrents", string(d))
	}
}

func ActiveTorrents(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	TORRENTS := GetAllTorrents()
	d, _ := json.Marshal(TORRENTS)
	w.Write(d)
}

func GetDirContents(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	path := strings.Replace(AbsPath(filepath.Join(Root, r.URL.Path)), "/dir", "", 1)
	if IsDir, err := isDirectory(path); err == nil && IsDir {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			http.Error(w, "Directory not found", http.StatusNotFound)
			return
		}
		files, err := GetDirContentsMap(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(files) == 0 {
			http.Error(w, "Directory is empty", http.StatusNotFound)
			return
		}
		d, _ := json.Marshal(files)
		w.Write(d)
	} else {
		http.ServeFile(w, r, path)
	}
}

func AutoComplete(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	r.ParseForm()
	q := r.Form.Get("q")
	if q == "" {
		http.Error(w, "No query", http.StatusBadRequest)
		return
	}
	var client = http.DefaultClient
	resp, err := client.Get("https://streamm4u.ws/searchJS?term=" + url.QueryEscape(q))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	var data []string
	json.NewDecoder(resp.Body).Decode(&data)
	b, _ := json.Marshal(data)
	w.Write(b)
}

func SearchTorrents(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	r.ParseForm()
	q := r.Form.Get("q")
	if q == "" {
		http.Error(w, "No query", http.StatusBadRequest)
		return
	}
	var b []byte
	if q == "top100" {
		Search := PretifyResult(Top100Torrents())
		b, _ = json.Marshal(Search)
	} else {
		Search := PretifyResult(SearchTorrentReq(q))
		b, _ = json.Marshal(Search)
	}
	w.Write(b)
}
