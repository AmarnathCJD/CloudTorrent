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
	"strings"
	"time"

	"github.com/julienschmidt/sse"
)

var (
	SSEFeed = sse.New()
)

func GetTorrent(w http.ResponseWriter, r *http.Request) {
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
	torrent := GetTorrentByID(id)
	if torrent.Status == "" {
		http.Error(w, "Torrent not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(torrent)
}

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
	if ok, err := PauseTorrentByID(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if !ok {
		http.Error(w, "Torrent not found", http.StatusNotFound)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
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
	if ok, err := ResumeTorrentByID(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if !ok {
		http.Error(w, "Torrent not found", http.StatusNotFound)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func DropAll(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	if err := DropAllTorrents(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
	Details := SysInfo{
		IP:        GetIP(r),
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		CPU:       fmt.Sprint(runtime.NumCPU()),
		Mem:       MemUsage(),
		Disk:      fmt.Sprintf("%s/%s", Disk.Used, Disk.All),
		Downloads: fmt.Sprint(GetLenTorrents()),
	}
	b, _ := json.Marshal(Details)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func DeleteFileHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	path := strings.Replace(AbsPath(filepath.Join(Root, r.URL.Path)), "api/deletefile/", "", 1)
	if strings.Contains(path, "/downloads/downloads") {
		path = strings.Replace(path, "/downloads", "", 1)
	}
	if strings.Contains(path, "torrents.db") || r.URL.Path == "/api/deletefile/downloads/torrents" {
		http.Error(w, "Protected path, cant delete!", http.StatusBadRequest)
		return
	}
	if err := DeleteFile(path); err != nil {
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
	DirPath := AbsPath(strings.Replace(AbsPath(filepath.Join(Root, r.FormValue("path"))), "/downloads", "", 1))
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
	DirPath := AbsPath(strings.Replace(AbsPath(filepath.Join(Root, r.URL.Path)), "/api/create/downloads", "", 1))
	if err := os.MkdirAll(DirPath, 0777); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func streamTorrentUpdate() {
	fmt.Println("Streaming Torrents started...")
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
			w.Write([]byte("[]"))
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
	resp, err := client.Get("https://streamm4u.ws/searchJS?term=" + url.QueryEscape(q)) // Will improve
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
	w.Write(GatherSearchResults(q))
}

func ZipFolderHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {

			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	r.ParseForm()
	path := strings.Replace(AbsPath(filepath.Join(Root, r.URL.Path)), "api/zip/", "", 1)
	if strings.Contains(path, "/downloads/downloads") {
		path = strings.Replace(path, "/downloads", "", 1)
	}
	folderName := filepath.Base(path)
	_, err := ZipDir(path, folderName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	filePath := "/dir/torrents/" + folderName + ".zip"
	var data = map[string]string{
		"file": filePath,
		"name": folderName + ".zip",
	}
	d, _ := json.Marshal(data)
	w.Write(d)
}

func init() {
	var API = []Handle{
		{"/api/add", AddTorrent},
		{"/api/torrents", ActiveTorrents},
		{"/api/torrent", GetTorrent},
		{"/api/status", SystemStats},
		{"/api/remove", DeleteTorrent},
		{"/api/pause", PauseTorrent},
		{"/api/resume", ResumeTorrent},
		{"/api/search/", SearchTorrents},
		{"/api/autocomplete", AutoComplete},
		{"/api/removeall", DropAll},
		{"/api/stopall", StopAllHandler},
		{"/api/startall", StartAllHandler},
		{"/api/upload", UploadFileHandler},
		{"/api/create/", CreateFolderHandler},
		{"/api/deletefile/", DeleteFileHandler},
		{"/api/zip/", ZipFolderHandler},
	}
	for _, api := range API {
		http.HandleFunc(api.Path, api.Func)
	}
	// update Server Events
	http.Handle("/torrents/update", SSEFeed)
	// Serve files
	http.HandleFunc("/dir/", GetDirContents)
}
