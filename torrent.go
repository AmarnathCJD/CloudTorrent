package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
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
	Name     string `json:"name,omitempty"`
	Size     string `json:"size,omitempty"`
	Status   string `json:"status,omitempty"`
	Magnet   string `json:"magnet,omitempty"`
	ID       string `json:"id,omitempty"`
	UID      string `json:"uid,omitempty"`
	Perc     string `json:"perc,omitempty"`
	Eta      string `json:"eta,omitempty"`
	Speed    string `json:"speed,omitempty"`
	Progress string `json:"progress,omitempty"`
	Icon     string `json:"icon,omitempty"`
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
	if _, err := os.Stat(Root + "/downloads/torrents/"); err != nil {
		err := os.Mkdir(Root+"/downloads/torrents/", 0777)
		if err != nil {
			fmt.Println(err)
		}
	}
	config.DataDir = Root + "/downloads/torrents/"
	config.Database = Root + "/downloads/torrents/torrents.db"
	client, err := torrent.NewSession(config)
	if err != nil {
		log.Print(err)
	}
	return client
}

func AddTorrentByMagnet(magnet string) (bool, error) {
	if CheckDuplicateTorrent(magnet) {
		return false, fmt.Errorf("torrent already exists")
	}
	_, err := client.AddURI(magnet, &torrent.AddTorrentOptions{StopAfterDownload: true})
	if err != nil {
		return false, err
	}
	return true, nil
}

func DeleteTorrentByID(id string) (bool, error) {
	if Torr := client.GetTorrent(id); Torr != nil {
		err := client.RemoveTorrent(id)
		return true, err
	}
	return false, nil
}

func PauseTorrentByID(id string) (bool, error) {
	if t := client.GetTorrent(id); t != nil {
		err := t.Stop()
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func ResumeTorrentByID(id string) (bool, error) {
	if t := client.GetTorrent(id); t != nil {
		err := t.Start()
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func DropAllTorrents() error {
	log.Println("Dropping all torrents")
	var err error
	for _, t := range GetTorrents() {
		err = client.RemoveTorrent(t.ID())
	}
	return err
}

func GetTorrents() []*torrent.Torrent {
	return client.ListTorrents()
}

func GetTorrentPath(id string) string {
	if Torr := client.GetTorrent(id); Torr != nil {
		return "/downloads/torrents/" + Torr.ID() + "/" + Torr.Stats().Name
	}
	return "404 Not Found"
}

func GetAllTorrents() []TorrentMeta {
	var Torrents []TorrentMeta
	for _, t := range GetTorrents() {
		Perc := GetDownloadPercentage(t.ID())
		Name := t.Stats().Name
		Icon := "bi bi-pause-circle"
		if Name == "" {
			Name = "unknown"
		}
		Stats, Icon := GetStats(t.ID())
		Torrents = append(Torrents, TorrentMeta{
			Name:     Name,
			Size:     ByteCountSI(t.Stats().Bytes.Total),
			Status:   Stats,
			Magnet:   t.Stats().InfoHash.String(),
			UID:      t.ID(),
			Perc:     Perc,
			Eta:      fmt.Sprint(t.Stats().ETA),
			Speed:    GetDownloadSpeed(t),
			Progress: GetProgress(Perc),
			Icon:     Icon,
		})
	}
	Torrents = SortAlpha(Torrents)
	for i := range Torrents {
		Torrents[i].ID = strconv.Itoa(i + 1)
	}
	return Torrents
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

func GetProgress(perc string) string {
	return strings.Replace(perc, "%", "", 1)
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

func GetDownloadSpeed(t *torrent.Torrent) string {
	if t.Stats().Speed.Download != 0 {
		return ByteCountSI(int64(t.Stats().Speed.Download)) + "/s"
	} else {
		return "-/-"
	}
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
	log.Println(argv)
	if len(argv) <= 1 {
		return ""
	}
	return strings.ToLower(argv[1])
}

func GetStats(id string) (string, string) {
	torr := client.GetTorrent(id)
	if torr != nil {
		if torr.Stats().Bytes.Total == 0 || torr.Stats().Status == torrent.DownloadingMetadata {
			return "Fetching Metadata", "bi bi-meta"
		} else if torr.Stats().Bytes.Downloaded >= torr.Stats().Bytes.Total {
			return "Completed", "bi bi-cloud-upload"
		} else if torr.Stats().Status == torrent.Downloading {
			return "Downloading", "bi bi-pause-circle"
		} else {
			if fmt.Sprint(torr.Stats().Status) == "Stopped" {
				return "Stopped", "bi bi-skip-start"
			} else {
				return fmt.Sprint(torr.Stats().Status), "bi bi-play-circle"
			}
		}
	}
	return "Error", "bi bi-bug"
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

func PretifyResult(result []TpbTorrent) []TpbTorrent {
	var Torr = result
	for i, t := range Torr {
		Torr[i].Magnet = "magnet:?xt=urn:btih:" + t.InfoHash + "&dn=" + url.QueryEscape(t.Name)
		Torr[i].Size = ByteCountSI(StringToInt64(t.Size))
	}
	return Torr
}

func GetLenTorrents() int {
	return len(GetTorrents())
}
