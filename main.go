package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

var (
	Wd, _ = os.Getwd()
	Root  = filepath.Join(Wd, "downloads")
)

func main() {
	fmt.Print("Starting server...")
	ServeApiEndpoints()
	HTMLServe()
	go streamTorrentUpdate()
	http.HandleFunc("/dir/", GetDirContents)
	http.HandleFunc("/delete/", DeleteFile)
	http.HandleFunc("/torrents/details", GetTorrDir)
	if err := http.ListenAndServe(":"+PORT(), nil); err != nil {
		panic(err)
	}
}

func ServeApiEndpoints() {
	http.HandleFunc("/api/status", SystemStats)
	http.HandleFunc("/api/torrents", ActiveTorrents)
	http.HandleFunc("/api/autocomplete/", AutoComplete)
	http.HandleFunc("/api/search/", SearchTorrents)
	http.HandleFunc("/api/add", AddTorrent)
	http.HandleFunc("/api/remove", DeleteTorrent)
	http.HandleFunc("/api/pause", PauseTorrent)
	http.HandleFunc("/api/resume", ResumeTorrent)
	http.HandleFunc("/api/removeall", DropAll)
	// update Server Events
	http.Handle("/torrents/update", SSEFeed)
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
