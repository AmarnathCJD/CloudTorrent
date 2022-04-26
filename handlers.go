package main

import (
	"encoding/json"
	"fmt"
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

func SystemStats(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	Disk := DiskUsage(Root)
	SystemStat := "<p><b>IP:</b> " + GetIP(r) + " " + "<b>OS:</b> " + runtime.GOOS + " " + "<b>Arch:</b> " + runtime.GOARCH + " " + "<b>CPU:</b> " + fmt.Sprint(runtime.NumCPU()) + " " + "<b>RAM:</b> " + MemUsage() + " " + "<b>Disk:</b> " + fmt.Sprintf("%s/%s", Disk.Used, Disk.All) + " " + "<b>Downloads:</b> " + strconv.Itoa(0) + "</p>"
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

func streamTorrentUpdate() {
	fmt.Println("Streaming Torrents started")
	for range time.Tick(time.Second * 1) {
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

func TorrHtml() string {
	var html = `<tr><th class="id">ID</th><th class="name">Name</th><th class="size">Size</th><th class="status">Status</th><th class="status">Progress</th><th class="status">ETA</th><th class="status">Download Speed</th><th class="action">Action</th></tr>`
	for _, torrent := range GetAllTorrents() {
		html += "<tr><th class='id'>" + torrent.ID + "</th>" + "<th class='name'><a href='" + GetTorrentPath(torrent.UID) + "'>" + torrent.Name + "</a>" + "</th>" + "<th class='size'>" + torrent.Size + "</th>" + "<th class='status'>" + torrent.Status + "</th>" + "<th class='status'>" + torrent.Perc + "</th>" + "<th class='status'>" + torrent.Eta + "</th>" + "<th class='status'>" + torrent.Speed + "</th>" + "<th class='action'>" + "<a href='torrents/details?uid=" + torrent.UID + "' class='download'>Download</a>" + "<a href='/' class='delete' onclick='return DeleteBtn(this)' data-uid='" + torrent.UID + "'>Delete</a>" + "</th></tr>"
	}
	return html
}

func GetDirContents(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	path := strings.Replace(AbsPath(filepath.Join(Root, r.URL.Path)), "/dir", "", 1)
	fmt.Println(path)
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
		d, _ := json.Marshal(files)
		w.Write(d)
	} else {
		http.ServeFile(w, r, path)
	}
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
	fmt.Println(path)
	if p, err := os.Stat(path); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else {
		if p.IsDir() {
			http.Redirect(w, r, "/downloads/downloads/torrents/"+uid, http.StatusFound)
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
		data = strings.Replace(data, "{{name}}", v.Name, -1)
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
