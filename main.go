package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const (
	root = "C:/Users/Rose/Downloads"
	// dir to serve
)

func main() {
	fmt.Println("Server started on port " + PORT())
	http.Handle("/", http.FileServer(http.Dir("./static/")))
	http.HandleFunc("/api/v1/status", SystemStats)
	http.HandleFunc("/api/v1/torrents", TorrentsStats)
	http.HandleFunc("/downloads/", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("./static/downloads.html"))
		template.Execute(w, nil)
	})
	http.HandleFunc("/stream/", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("./static/player.html"))
		template.Execute(w, nil)
	})
	http.Handle("/torrents/update", SSEFeed)
	go streamTorrentUpdate()
	http.HandleFunc("/dir/", GetDirContents)
	fmt.Println(http.ListenAndServe(":"+PORT(), nil))
}
