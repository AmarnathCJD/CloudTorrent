package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/cenkalti/rain/torrent"
)

var (
	client   = InitClient()
	Torrents = make(map[string]torrent.Torrent)
)

type TorrentsResponse struct {
	Torrents []TorrentMeta `json:"torrents,omitempty"`
}

type TorrentMeta struct {
	Name   string `json:"name,omitempty"`
	Size   string `json:"size,omitempty"`
	Date   string `json:"date,omitempty"`
	Magnet string `json:"magnet,omitempty"`
	ID     string `json:"id,omitempty"`
}

func InitClient() *torrent.Session {
	config := torrent.DefaultConfig
	config.DataDir = root + "/downloads"
	client, err := torrent.NewSession(config)
	if err != nil {
		panic(err)
	}
	return client
}

func AddMagnet(magnet string) error {
	t, err := client.AddURI(magnet, &torrent.AddTorrentOptions{
		StopAfterDownload: true,
	})
	if err != nil {
		return err
	}
	Torrents[magnet] = *t
	d, _ := json.Marshal(Torrents)
	fmt.Println(string(d))
	return nil
}

func GetTorrents() map[string]torrent.Torrent {
	return Torrents
}

func GetTorrent(magnet string) torrent.Torrent {
	return Torrents[magnet]
}

func GetDownloads(magnet string) {
	x := client.ListTorrents()
	for _, t := range x {
		fmt.Println(t.Stats())
	}
}

func CancelTorrent(magnet string) {
	t := Torrents[magnet]
	t.Stop()
	delete(Torrents, magnet)
}

func GetTorrentStatus(magnet string) torrent.Stats {
	torr := client.ListTorrents()
	for _, t := range torr {
		if t.InfoHash().String() == magnet {
			return t.Stats()
		}
	}
	return torrent.Stats{}
}

func TorrentsServe() {
	http.HandleFunc("/torrents", func(w http.ResponseWriter, r *http.Request) {
		t := GetActiveTorrents()
		d, _ := json.Marshal(t)
		w.Write(d)
	})
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func GetActiveTorrents() []TorrentMeta {
	torr := client.ListTorrents()
	var torrents []torrent.Stats
	for _, t := range torr {
		torrents = append(torrents, t.Stats())
	}
	t := TorrentsResponse{}
	Magnets := []string{}
	j := 0
	for _, v := range torrents {
		if StringInSlice(v.InfoHash.String(), Magnets) {
			continue
		}
		j++
		Magnets = append(Magnets, v.InfoHash.String())
		t.Torrents = append(t.Torrents, TorrentMeta{
			Name:   v.Name,
			Size:   "_",
			Date:   "-",
			Magnet: v.InfoHash.String(),
			ID:     fmt.Sprint(j),
		})
	}
	for range t.Torrents {
		sort.Slice(t.Torrents, func(i, j int) bool {
			return t.Torrents[i].ID <= t.Torrents[j].ID
		}) // sort by ID fix chane to sort alphabetically TODO
	}
	return t.Torrents
}
