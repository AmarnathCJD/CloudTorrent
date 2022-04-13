package main

import (
	"fmt"
	"net/http"
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
	http.HandleFunc("/torrents/add", AddTorrent)
	http.HandleFunc("/torrents/delete", DeleteTorrent)
	http.HandleFunc("/torrents", TorrentsServe)
	http.HandleFunc("/torrents/details", GetTorrDir)
	http.ListenAndServe(":80", nil)
}
