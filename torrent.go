package main

import (
	"fmt"
	"os"
	"strings"

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
	Perc   string `json:"perc,omitempty"`
	Eta    string `json:"eta,omitempty"`
	Speed  string `json:"speed,omitempty"`
}

func InitClient() *torrent.Session {
	config := torrent.DefaultConfig
	if _, err := os.Stat(root + "/downloads/torrents/"); err != nil {
		err := os.Mkdir(root+"/downloads/torrents/", 0777)
		if err != nil {
			fmt.Println(err)
		}
	}
	config.DataDir = root + "/downloads/torrents/"
	config.Database = root + "/downloads/torrents/torrents.db"
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
	fmt.Print("Added torrent: ", t.Name())
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
			return root + "/downloads/torrents/" + t.ID() + "/" + t.Stats().Name
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
					Size:   ByteCountSI(GetTorrentSize(t.ID())),
					Perc:   GetDownloadPercentage(t.ID()),
					Status: fmt.Sprint(t.Stats().Status),
					Magnet: t.Stats().InfoHash.String(),
					Speed:  fmt.Sprint(ByteCountSI(int64(t.Stats().Speed.Download))) + "/s",
					ID:     fmt.Sprintf("%d", IDno),
					UID:    t.ID(),
					Eta:    fmt.Sprint(t.Stats().ETA),
				})
			}
		}
	}
	Torrents = SortAlpha(Torrents)
	return Torrents.Torrents
}

func GetDownloadPercentage(id string) string {
	torr := client.ListTorrents()
	for _, t := range torr {
		if t.ID() == id {
			if t.Stats().Pieces.Total != 0 {
				p := float64(t.Stats().Pieces.Have) / float64(t.Stats().Pieces.Total)
				return fmt.Sprintf("%.2f", p*100) + "%"
			} else {
				return "-"
			}
		}
	}
	return "0" + "%"
}

func GetTorrentSize(id string) int64 {
	torr := client.ListTorrents()
	for _, t := range torr {
		if t.ID() == id {
			if t.Stats().Bytes.Total != 0 {
				return int64(t.Stats().Bytes.Total)
			} else {
				return 0
			}
		}
	}
	return 0
}

func GetPeers(id string) int {
	torr := client.ListTorrents()
	for _, t := range torr {
		if t.ID() == id {
			return t.Stats().Peers.Total
		}
	} // soon
	return 0
}

func UpdateOnComplete() {
	// soon
}

func CheckDuplicateTorrent(magnet string) bool {
	magnet = ParseHashFromMagnet(magnet)
	torr := client.ListTorrents()
	for _, t := range torr {
		fmt.Println(t.Stats().InfoHash.String())
		if t.Stats().InfoHash.String() == magnet {
			return true
		}
	}
	return false
}

func ParseHashFromMagnet(magnet string) string {
	args := strings.Split(magnet, "&")
	argv := strings.Split(args[0], "btih:")
	if len(argv) <= 1 {
		return ""
	}
	return strings.ToLower(argv[1])
}
