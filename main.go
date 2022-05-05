package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	Wd, _ = os.Getwd()
	Root  = filepath.Join(Wd, "downloads")
	Port  = GetOutboundPort()
)

func main() {
	fmt.Print("Starting server...")
	ServeApiEndpoints()
	HTMLServe()
	go streamTorrentUpdate()
	if err := http.ListenAndServe(Port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func ServeApiEndpoints() {
	var API = []Handle{
		{"/api/add", AddTorrent},
		{"/api/torrents", ActiveTorrents},
		{"/api/status", SystemStats},
		{"/api/remove", DeleteTorrent},
		{"/api/pause", PauseTorrent},
		{"/api/resume", ResumeTorrent},
		{"/api/search", SearchTorrents},
		{"/api/autocomplete", AutoComplete},
		{"/api/removeall", DropAll},
		{"/api/stopall", StopAllHandler},
		{"/api/startall", StartAllHandler},
		{"/api/upload", UploadFileHandler},
		{"/api/create/", CreateFolderHandler},
	}
	for _, api := range API {
		http.HandleFunc(api.Path, api.Func)
	}
	// update Server Events
	http.Handle("/torrents/update", SSEFeed)
	http.HandleFunc("/dir/", GetDirContents)
	http.HandleFunc("/delete/", DeleteFile)
}

func HTMLServe() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})
	http.HandleFunc("/downloads/", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("./static/downloads.html"))
		template.Execute(w, nil)
	})
	http.HandleFunc("/stream/", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("./static/player.html"))
		template.Execute(w, nil)
	})
	http.HandleFunc("/search/", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("./static/search.html"))
		template.Execute(w, nil)
	})
	// static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
}
