package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

var (
	WD, _ = os.Getwd()
	Root  = filepath.Join(WD, "downloads")
)

func main() {
	fmt.Print("Starting server...")
	http.Handle("/", http.FileServer(http.Dir("./static/")))
	http.HandleFunc("/api/status", SystemStats)
	http.HandleFunc("/api/torrents", TorrentsStats)
	http.HandleFunc("/api/autocomplete/", AutoComplete)
	http.HandleFunc("/api/search/", SearchTorrents)
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
	http.Handle("/torrents/update", SSEFeed)
	go streamTorrentUpdate()
	http.HandleFunc("/dir/", GetDirContents)
	http.HandleFunc("/add", AddTorrent)
	http.HandleFunc("/delete/", DeleteFile)
	http.HandleFunc("/torrents/add", AddTorrent)
	http.HandleFunc("/torrents/delete", DeleteTorrent)
	http.HandleFunc("/torrents/details", GetTorrDir)
	fmt.Println(http.ListenAndServe(":"+PORT(), nil))
}
