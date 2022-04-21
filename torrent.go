package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
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

type TpbTorrent struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	InfoHash string `json:"info_hash"`
	Leechers string `json:"leechers"`
	Seeders  string `json:"seeders"`
	NumFiles string `json:"num_files"`
	Size     string `json:"size"`
	Username string `json:"username"`
	Added    string `json:"added"`
	Status   string `json:"status"`
	Magnet   string `json:"magnet"`
}

func InitClient() *torrent.Session {
	config := torrent.DefaultConfig
	client, err := torrent.NewSession(config)
	if err != nil {
		log.Print(err)
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
			if fmt.Sprint(torr.Stats().Status) == "Stopped" {
				return "Completed"
			} else {
				return fmt.Sprint(torr.Stats().Status)
			}
		}
	}
	return "Error"
}

func SearchTorrentReq(query string) []TpbTorrent {
	var baseUrl = "https://tpb23.ukpass.co/apibay/q.php" + "?q=" + url.QueryEscape(query) + "&cat=0"
	var resp *http.Response
	var err error
	if resp, err = http.Get(baseUrl); err != nil {
		return []TpbTorrent{}
	}
	defer resp.Body.Close()
	var tpb []TpbTorrent
	if err = json.NewDecoder(resp.Body).Decode(&tpb); err != nil {
		return []TpbTorrent{}
	}
	return tpb
}

func Top100Torrents() []TpbTorrent {
	baseUrl := "https://tpb23.ukpass.co/apibay/precompiled/data_top100_all.json"
	var resp *http.Response
	var err error
	if resp, err = http.Get(baseUrl); err != nil {
		return []TpbTorrent{}
	}
	defer resp.Body.Close()
	var tpb []TpbTorrent
	var data []interface{}
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return []TpbTorrent{}
	}
	for _, v := range data {
		tpb = append(tpb, TpbTorrent{
			Name:     v.(map[string]interface{})["name"].(string),
			Size:     fmt.Sprint(int64(v.(map[string]interface{})["size"].(float64))),
			Seeders:  fmt.Sprint(int64(v.(map[string]interface{})["seeders"].(float64))),
			Leechers: fmt.Sprint(int64(v.(map[string]interface{})["leechers"].(float64))),
			ID:       fmt.Sprint(int64(v.(map[string]interface{})["id"].(float64))),
			Added:    fmt.Sprint(int64(v.(map[string]interface{})["added"].(float64))),
			InfoHash: fmt.Sprint(v.(map[string]interface{})["info_hash"].(string)),
		})
	}
	return tpb
}

func GenMagnetFromResult(result []TpbTorrent) []TpbTorrent {
	var Torr = result
	for i, t := range Torr {
		Torr[i].Magnet = "magnet:?xt=urn:btih:" + t.InfoHash + "&dn=" + url.QueryEscape(t.Name)
	}
	return Torr
}
