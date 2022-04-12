package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	Status string `json:"status,omitempty"`
	Magnet string `json:"magnet,omitempty"`
	ID     string `json:"id,omitempty"`
	UID    string `json:"uid,omitempty"`
}

func InitClient() *torrent.Session {
	config := torrent.DefaultConfig
	config.DataDir = root + "/torrents"
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

func GetTorrentPath(id string) string {
	torr := client.ListTorrents()
	for _, t := range torr {
		if t.ID() == id {
			return root + "/torrents/" + t.ID()
		}
	}
	return ""
}

func DeleteTorrentByID(id string) (bool, error) {
	for _, t := range client.ListTorrents() {
		if t.ID() == id {
			err := client.RemoveTorrent(id)
			delete(Torrents, t.Stats().InfoHash.String())
			return true, err
		}
	}
	return false, nil
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

func GetActiveTorrents() []TorrentMeta {
	torr := client.ListTorrents()
	Torrents := TorrentsResponse{}
	Magnets := []string{}
	IDno := 0
	for _, t := range torr {
		if t.Name() != "" {
			if !StringInSlice(t.Stats().InfoHash.String(), Magnets) {
				Magnets = append(Magnets, t.Stats().InfoHash.String())
				IDno++
				Torrents.Torrents = append(Torrents.Torrents, TorrentMeta{
					Name:   t.Name(),
					Size:   fmt.Sprint(t.Stats().Pieces.Total),
					Status: t.Stats().Status.String(),
					Magnet: t.Stats().InfoHash.String(),
					ID:     fmt.Sprintf("%d", IDno),
					UID:    t.ID(),
				})
			}
		}
	}
	Torrents = SortAlpha(Torrents)
	return Torrents.Torrents
}

// TODO make torrent name clickable, add popup confirmation, link torrent file dir, ddos protection 
