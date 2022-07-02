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
	HTMLServe()
	go streamTorrentUpdate()
	if err := http.ListenAndServe(Port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
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

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	http.HandleFunc("/test/", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("./assets/test.html"))
		template.Execute(w, nil)
	})
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
}
