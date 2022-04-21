package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

const (
	root = ""
	// dir to serve
)

func main() {
	fmt.Println("Server started on port 8080")
	http.Handle("/", http.FileServer(http.Dir("./static/")))
	http.HandleFunc("/api/v1/disk", func(w http.ResponseWriter, r *http.Request) {
		diskUsage := DiskUsage(root)
		json, _ := json.Marshal(diskUsage)
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	})
	http.HandleFunc("/downloads/", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("./static/downloads.html"))
		template.Execute(w, nil)
	})
	http.HandleFunc("/home/", MainPage)
	http.HandleFunc("/add", AddTorrent)
	http.HandleFunc("/torrents/add", AddTorrent)
	http.HandleFunc("/torrents/delete", DeleteTorrent)
	http.HandleFunc("/torrents", TorrentsServe)
	http.HandleFunc("/torrents/details", GetTorrDir)
	http.HandleFunc("/torrents/search/", TorrentSearchPage)
	http.HandleFunc("/dir/", GetDirContents)
	fmt.Println(http.ListenAndServe(":80", nil))
}
