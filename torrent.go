package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/cenkalti/rain/torrent"
)

var (
	client = InitClient()
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
	return nil
}

func GetTorrents() []*torrent.Torrent {
	return client.ListTorrents()
}

func GetTorrentPath(id string) string {
	if Torr := client.GetTorrent(id); Torr != nil {
		return root + "/downloads/torrents/" + Torr.ID() + "/" + Torr.Stats().Name
	}
	return ""
}

func DeleteTorrentByID(id string) (bool, error) {
	if Torr := client.GetTorrent(id); Torr != nil {
		err := client.RemoveTorrent(id)
		return true, err
	}
	return false, nil
}

func GetTorrentStatus(id string) torrent.Stats {
	if Torr := client.GetTorrent(id); Torr != nil {
		return Torr.Stats()
	}
	return torrent.Stats{}
}

func GetActiveTorrents() []TorrentMeta {
	torr := client.ListTorrents()
	Torrents := TorrentsResponse{}
	IDno := 0
	for _, t := range torr {
		if t.Name() != "" {
			Torrents.Torrents = append(Torrents.Torrents, TorrentMeta{
				Name:   t.Name(),
				Size:   ByteCountSI(GetTorrentSize(t.ID())),
				Perc:   GetDownloadPercentage(t.ID()),
				Status: GetStats(t.ID()),
				Magnet: t.Stats().InfoHash.String(),
				Speed:  fmt.Sprint(ByteCountSI(int64(t.Stats().Speed.Download))) + "/s",
				UID:    t.ID(),
				Eta:    fmt.Sprint(t.Stats().ETA),
			})
		}
	}
	Torrents = SortAlpha(Torrents)
	for i := range Torrents.Torrents {
		Torrents.Torrents[i].ID = fmt.Sprint(IDno)
		IDno++
	}
	return Torrents.Torrents
}

func GetDownloadPercentage(id string) string {
	torr := client.GetTorrent(id)
	if torr != nil {
		if torr.Stats().Pieces.Total != 0 {
			return fmt.Sprintf("%.2f", float64(torr.Stats().Pieces.Have)/float64(torr.Stats().Pieces.Total)*100) + "%"
		}
	}
	return "0%"
}

func GetTorrentSize(id string) int64 {
	torr := client.GetTorrent(id)
	if torr != nil {
		if torr.Stats().Bytes.Total != 0 {
			return torr.Stats().Bytes.Total
		}
	}
	return 0
}

func GetPeers(id string) int {
	torr := client.GetTorrent(id)
	if torr != nil {
		return torr.Stats().Peers.Total
	}
	return 0
}

func UpdateOnComplete() {

}

func CheckDuplicateTorrent(magnet string) bool {
	magnet = ParseHashFromMagnet(magnet)
	for _, t := range GetTorrents() {
		if strings.ToLower(t.Stats().InfoHash.String()) == magnet {
			return true
		}
	}
	return false
}

func ParseHashFromMagnet(magnet string) string {
	var args []string
	if strings.Contains(magnet, "&") {
		args = strings.Split(magnet, "&")
	} else {
		args = []string{magnet}
	}
	argv := strings.Split(args[0], "btih:")
	if len(argv) <= 1 {
		return ""
	}
	return strings.ToLower(argv[1])
}

func GetStats(id string) string {
	torr := client.GetTorrent(id)
	if torr != nil {
		if torr.Stats().Bytes.Total == 0 {
			return "Waiting"
		} else if torr.Stats().Bytes.Downloaded == torr.Stats().Bytes.Total {
			return "Complete"
		} else {
			return "Downloading"
		}
	}
	return "Error"
}
