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
	if !CheckDuplicateTorrent(magnet) {
		if err := AddMagnet(magnet); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Torrent already exists", http.StatusBadRequest)
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

func SystemStats(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	Disk := DiskUsage(root)
	SystemStat := "<p><b>IP:</b> " + GetIP(r) + " " + "<b>OS:</b> " + runtime.GOOS + " " + "<b>Arch:</b> " + runtime.GOARCH + " " + "<b>CPU:</b> " + fmt.Sprint(runtime.NumCPU()) + " " + "<b>RAM:</b> " + MemUsage() + " " + "<b>Disk:</b> " + fmt.Sprintf("%s/%s", Disk.Used, Disk.All) + " " + "<b>Downloads:</b> " + strconv.Itoa(0) + "</p>"
	w.Write([]byte(SystemStat))
}

func UpdateTorrents(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	w.Header().Set("Content-Type", "text/event-stream")
	Torrents := GetTorrents()
	var data = ""
	for _, torrent := range Torrents {
		data += torrent.InfoHash().String() + "," + torrent.Name()
	}
	w.Write([]byte("data: " + data + "\n\n"))
	w.(http.Flusher).Flush()
}

func streamTorrentUpdate() {
	fmt.Println("Streaming torr  started")
	for range time.Tick(time.Second * (1 / 3)) {
		SSEFeed.SendString("", "torrents", TorrHtml())
	}
}

func TorrentsStats(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	fmt.Fprint(w, TorrHtml())
}

func TorrentsToHtml(t []TorrentMeta) string {
	var html = `<tr><th class="id">ID</th><th class="name">Name</th><th class="size">Size</th><th class="status">Status</th><th class="status">Progress</th><th class="status">ETA</th><th class="status">Download Speed</th><th class="action">Action</th></tr>`
	for _, torrent := range t {
		html += "<tr><td class='id'>" + torrent.ID + "</td>" + "<td class='name'><a href='/torrents/details?uid=" + torrent.UID + "'>" + torrent.Name + "</a>" + "</td>" + "<td class='size'>" + torrent.Size + "</td>" + "<td class='status'>" + torrent.Status + "</td>" + "<td class='status'>" + torrent.Perc + "</td>" + "<td class='status'>" + torrent.Eta + "</td>" + "<td class='status'>" + torrent.Speed + "</td>" + "<td class='action'>" + "<a href='torrents/details?uid=" + torrent.UID + "' class='download'>Download</a>" + "<a href='/' class='delete' onclick='return DeleteBtn(this)' data-uid='" + torrent.UID + "'>Delete</a>" + "</td></tr>"
	}
	return html
}

func TorrHtml() string {
	var html = `<tr><th class="id">ID</th><th class="name">Name</th><th class="size">Size</th><th class="status">Status</th><th class="status">Progress</th><th class="status">ETA</th><th class="status">Download Speed</th><th class="action">Action</th></tr>`
	for _, torrent := range GetActiveTorrents() {
		html += "<tr><th class='id'>" + torrent.ID + "</th>" + "<th class='name'><a href='/torrents/details?uid=" + torrent.UID + "'>" + torrent.Name + "</a>" + "</th>" + "<th class='size'>" + torrent.Size + "</th>" + "<th class='status'>" + torrent.Status + "</th>" + "<th class='status'>" + torrent.Perc + "</th>" + "<th class='status'>" + torrent.Eta + "</th>" + "<th class='status'>" + torrent.Speed + "</th>" + "<th class='action'>" + "<a href='torrents/details?uid=" + torrent.UID + "' class='download'>Download</a>" + "<a href='/' class='delete' onclick='return DeleteBtn(this)' data-uid='" + torrent.UID + "'>Delete</a>" + "</th></tr>"
	}
	return html
}

func TorrentsServe(w http.ResponseWriter, r *http.Request) {
	Torrents := GetActiveTorrents()
	d, _ := json.Marshal(Torrents)
	w.Write(d)
}

func GetDirContents(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	path := filepath.Join(root, r.URL.Path)
	path = strings.Replace(path, "\\dir", "", -1)
	if IsDir, err := isDirectory(path); err == nil && IsDir {
		var files []map[string]string
		if _, err := os.Stat(path); os.IsNotExist(err) {
			http.Error(w, "Directory not found", http.StatusNotFound)
			return
		}
		files = make([]map[string]string, 0)
		f, err := ioutil.ReadDir(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for _, info := range f {
			AbsPath := strings.Replace(path, "\\", "/", -1)
			FType, FaClass, FaColor := GetFileType(info.Name())
			var Name string
			if info.IsDir() {
				Name = GetDirName(info.Name())
			} else {
				Name = GetFileName(info.Name())
			}
			files = append(files, map[string]string{
				"name":  Name,
				"size":  ByteCountSI(int64(info.Size())),
				"type":  FType,
				"isdir": strconv.FormatBool(info.IsDir()),
				"path":  "/downloads" + strings.Replace(AbsPath, root, "", 1) + "/" + info.Name(),
				"class": FaClass,
				"color": FaColor,
				"ext":   filepath.Ext(info.Name()),
			})
		}
		d, _ := json.Marshal(files)
		w.Write(d)
	} else {
		http.ServeFile(w, r, path)
	}
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
		FileType, _, _ := GetFileType(v.Name())
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
	tbl := `<tr><th class="id">{{id}}</th><th class="name"><a href="/torrents/details?uid={{uid}}">{{name}}</a></th><th class="size">{{size}}</th><th class="status">{{status}}</th><th class="status">{{percent}}</th><th class="status">{{eta}}</th><th class="status">{{speed}}</th><th class="action"><a href="%s" class="download">Download</a><a href="%s" class="delete">Delete</a></th></tr>`
	data := ""
	for i, v := range torrs {
		data += fmt.Sprintf(tbl, "/torrents/details?uid="+v.UID, "/torrents/delete?uid="+v.UID)
		data = strings.Replace(data, "{{id}}", strconv.Itoa(i+1), -1)
		data = strings.Replace(data, "{{name}}", v.Name, -1)
		data = strings.Replace(data, "{{size}}", v.Size, -1)
		data = strings.Replace(data, "{{speed}}", v.Speed, -1)
		data = strings.Replace(data, "{{status}}", v.Status, -1)
		data = strings.Replace(data, "{{percent}}", v.Perc, -1)
		data = strings.Replace(data, "{{eta}}", v.Eta, -1)
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
	if p, err := os.Stat(path); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else {
		if p.IsDir() {
			serveDir(w, r, path)
		} else {
			http.ServeFile(w, r, path)
		}
	}

}

func TorrentSearchPage(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	r.ParseForm()
	query := r.Form.Get("query")
	if query == "" {
		http.Error(w, "No query", http.StatusBadRequest)
		return
	}
	var t []TpbTorrent
	if query == "top100" {
		t = GenMagnetFromResult(Top100Torrents())
	} else {
		t = GenMagnetFromResult(SearchTorrentReq(query))
	}
	data := ""
	table := `<li class="table-row">
    <div class="col col-1" data-label="ID">{{id}}</div>
    <div class="col col-2" data-label="Name">{{name}}</div>
    <div class="col col-3" data-label="Size">{{size}}</div>
    <div class="col col-4" data-label="Seeders">{{seeders}}</div>
    <div class="col col-5" data-label="Leechers">{{leechers}}</div>
    <div class="col col-6" data-label="Action">
        <button class="btn" onclick="AddedTorr(this)" data-magnet="{{magnet}}"><i class="fa fa-download"></i>Download</button>
    </div>
</li>`
	total := 0
	page := torrentsearch
	for i, v := range t {
		total++
		if total > 25 {
			break
		}
		data += table
		data = strings.Replace(data, "{{id}}", strconv.Itoa(i+1), -1)
		data = strings.Replace(data, "{{name}}", GetFileName(v.Name), -1)
		data = strings.Replace(data, "{{size}}", ByteCountSI(StringToInt64(v.Size)), -1)
		data = strings.Replace(data, "{{seeders}}", v.Seeders, -1)
		data = strings.Replace(data, "{{leechers}}", v.Leechers, -1)
		data = strings.Replace(data, "{{added}}", v.Added, -1)
		data = strings.Replace(data, "{{magnet}}", v.Magnet, -1)
	}
	page = strings.Replace(page, "{{#torrents}}", data, -1)
	page = strings.Replace(page, "{{query}}", query, -1)
	page = strings.Replace(page, "{{count}}", strconv.Itoa(len(t)), -1)
	w.Write([]byte(page))
}
