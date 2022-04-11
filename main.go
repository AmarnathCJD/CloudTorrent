package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

const (
	root = "C:/Users/rj/Downloads"
	// dir to serve
)

func main() {
	fmt.Println("Server started on port 8080")
	http.HandleFunc("/downloads/", File)
	http.HandleFunc("/", MainPage)
	http.HandleFunc("/add", AddTorrent)
	TorrentsServe()
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
		files = strings.Replace(files, "{{name}}", "<a href=\"/downloads/"+v.Name()+"\">"+strings.Title(GetFileName(v.Name()))+"</a>", -1)
		files = strings.Replace(files, "{{size}}", Size, -1)
		files = strings.Replace(files, "{{type}}", FileType, -1)
		files = strings.Replace(files, "{{date}}", strings.ReplaceAll(v.ModTime().String(), "+0000 GMT", ""), -1)
	}
	html = strings.Replace(html, "{{files}}", files, -1)
	html = strings.Replace(html, "{{ip}}", IP, -1)
	return html
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	IP := GetIP(r)
	Disk := DiskUsage(root)
	torr := torrents
	torr = strings.Replace(torr, "{{ip}}", IP, -1)
	torrs := GetActiveTorrents()
	tbl := `<tr><th class="id">{{id}}</th><th class="name">{{name}}</th><th class="size">{{size}}</th><th class="date">{{date}}</th><th class="magnet">{{magnet}}</th><th class="action"><a href="#" class="download">Download</a><a href="#" class="delete">Delete</a></th></tr>`
	data := ""
	for i, v := range torrs {
		data += tbl
		data = strings.Replace(data, "{{id}}", strconv.Itoa(i+1), -1)
		data = strings.Replace(data, "{{name}}", v.Name, -1)
		data = strings.Replace(data, "{{size}}", v.Size, -1)
		data = strings.Replace(data, "{{date}}", v.Date, -1)
		data = strings.Replace(data, "{{magnet}}", v.Magnet, -1)
	}
	torr = strings.Replace(torr, "{{#each torrents}}", data, -1)
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	torr = strings.Replace(torr, "{{cpu}}", strconv.Itoa(runtime.NumCPU()), -1)
	torr = strings.Replace(torr, "{{memory}}", ByteCountSI(int64(mem.Alloc)), -1)
	torr = strings.Replace(torr, "{{goroutines}}", strconv.Itoa(runtime.NumGoroutine()), -1)
	torr = strings.Replace(torr, "{{torrents_len}}", strconv.Itoa(len(torrs)), -1)
	torr = strings.Replace(torr, "{{disk}}", fmt.Sprintf("%s/%s", Disk.Used, Disk.All), -1)
	torr = strings.Replace(torr, "{{ip}}", IP, -1)
	w.Write([]byte(torr))
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
	if err := AddMagnet(magnet); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	MainPage(w, r)
}
