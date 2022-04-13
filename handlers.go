package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/r3labs/sse/v2"
)

var (
	sseClient = sse.New()
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
	if err := AddMagnet(magnet); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	MainPage(w, r)
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
	}
	MainPage(w, r)
}

func TorrentsServe(w http.ResponseWriter, r *http.Request) {
	Torrents := GetActiveTorrents()
	d, _ := json.Marshal(Torrents)
	w.Write(d)
}

func GetHTMLDir(f map[string]os.FileInfo, IP string, rootDir string) string {
	rootDir = strings.Replace(rootDir, "\\", "/", -1)
	TorrDir := strings.Replace(rootDir, root, "", 1)
	TorrDir = strings.Replace(TorrDir, "\\", "/", -1)
	var html = downloads
	var files = ""
	tbl := `<tr> <td>{{name}}</td> <td>{{size}}</td> <td>{{type}}</td> <td>{{date}}</td> </tr>`
	for _, v := range f {
		var TorrDirV = TorrDir
		if !v.IsDir() {
			TorrDirV = TorrDir + "/" + v.Name()
		} else {
			TorrDirV = TorrDir + "/" + v.Name() + "/"
		}
		files += tbl
		Size := ByteCountSI(int64(v.Size()))
		FileType := GetFileType(v.Name())
		if v.IsDir() {
			FileType = "Folder"
			Size = "-"
		}
		files = strings.Replace(files, "{{name}}", fmt.Sprintf(`<a href="%s">`+strings.Title(GetFileName(v.Name()))+`</a>`, TorrDirV), -1)
		files = strings.Replace(files, "{{size}}", Size, -1)
		files = strings.Replace(files, "{{type}}", FileType, -1)
		files = strings.Replace(files, "{{date}}", strings.ReplaceAll(v.ModTime().String(), "+0000 GMT", ""), -1)
	}
	html = strings.Replace(html, "{{files}}", files, -1)
	html = strings.Replace(html, "{{ip}}", IP, -1)
	return html
}

func File(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(root, r.URL.Path)
	if filepath.Ext(path) == "" {
		serveDir(w, r, path)
		return
	}
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
	w.Write([]byte(GetHTMLDir(list, IP, path)))
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	Disk := DiskUsage(root)
	torr := torrents
	torrs := GetActiveTorrents()
	tbl := `<tr><th class="id">{{id}}</th><th class="name"><a href="/torrents/details?uid={{uid}}">{{name}}</a></th><th class="size">{{size}}</th><th class="status">{{status}}</th><th class="status">{{percent}}</th><th class="status">{{eta}}</th><th class="magnet">{{magnet}}</th><th class="action"><a href="%s" class="download">Download</a><a href="%s" class="delete">Delete</a></th></tr>`
	data := ""
	for i, v := range torrs {
		data += fmt.Sprintf(tbl, "/torrents/details?uid="+v.UID, "/torrents/delete?uid="+v.UID)
		data = strings.Replace(data, "{{id}}", strconv.Itoa(i+1), -1)
		data = strings.Replace(data, "{{name}}", v.Name, -1)
		data = strings.Replace(data, "{{size}}", v.Size, -1)
		data = strings.Replace(data, "{{status}}", v.Status, -1)
		data = strings.Replace(data, "{{percent}}", v.Perc, -1)
		data = strings.Replace(data, "{{eta}}", v.Eta, -1)
		data = strings.Replace(data, "{{magnet}}", v.Magnet, -1)
		data = strings.Replace(data, "{{uid}}", v.UID, -1)
	}
	torr = strings.Replace(torr, "{{#each torrents}}", data, -1)
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	torr = strings.Replace(torr, "{{cpu}}", strconv.Itoa(runtime.NumCPU()), -1)
	torr = strings.Replace(torr, "{{memory}}", ByteCountSI(int64(mem.Alloc)), -1)
	torr = strings.Replace(torr, "{{goroutines}}", strconv.Itoa(runtime.NumGoroutine()), -1)
	torr = strings.Replace(torr, "{{torrents_len}}", strconv.Itoa(len(torrs)), -1)
	torr = strings.Replace(torr, "{{disk}}", fmt.Sprintf("%s/%s", Disk.Used, Disk.All), -1)
	torr = strings.Replace(torr, "{{ip}}", GetIP(r), -1)
	torr = strings.Replace(torr, "{{hash}}", "hie", -1)
	torr = strings.Replace(torr, "{{progress}}", "0", -1)
	w.Write([]byte(torr))
}

func SendSSE(data string, event string) {
	sseClient.Headers["Content-Type"] = "text/event-stream"
	sseClient.Publish("messages", &sse.Event{
		Data: []byte("ping"),
	})
}

func GetTorrDir(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	r.ParseForm()
	uid := r.Form.Get("uid")
	if uid == "" {
		http.Error(w, "No UID", http.StatusBadRequest)
		return
	}

	path := GetTorrentPath(uid)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	serveDir(w, r, path)
}

func ListAllDirFiles(dir string) {
	defer func() {
		if err, ok := recover().(error); ok {
			fmt.Println(err)
		}
	}()
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		if file.IsDir() {
			ListAllDirFiles(filepath.Join(dir, file.Name()))
		} else {
			fmt.Println(file.Name())
		}
	}
}
