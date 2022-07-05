package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/cenkalti/rain/torrent"
)

var (
	client   *torrent.Session
	hClient  = &http.Client{Timeout: time.Second * 10}
	Trackers []string
)

func InitClient() *torrent.Session {
	config := torrent.DefaultConfig
	config.DataDir = Root + "/torrents/"
	config.Database = Root + "/torrents.db"
	client, err := torrent.NewSession(config)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

type TorrentData struct {
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
	Path     string `json:"path,omitempty"`
}

func AddTorrentByMagnet(magnet string) (bool, error) {
	if CheckDuplicateTorrent(magnet) {
		return false, fmt.Errorf("torrent already exists")
	}
	m, err := client.AddURI(magnet, &torrent.AddTorrentOptions{StopAfterDownload: true})
	if err != nil {
		return false, err
	}
	for i := range Trackers {
		m.AddTracker(Trackers[i])
	}
	return true, nil
}

func GetTrakers() {
	url := "https://raw.githubusercontent.com/ngosang/trackerslist/master/trackers_all.txt"
	resp, err := hClient.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	t, _ := ioutil.ReadAll(resp.Body)
	for _, line := range strings.Split(string(t), "\n") {
		if strings.HasPrefix(line, "udp://") {
			Trackers = append(Trackers, line)
		}
	}
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
		return true, nil
	}
	return false, nil
}

func ResumeTorrentByID(id string) (bool, error) {
	if t := client.GetTorrent(id); t != nil {
		err := t.Start()
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

func GetTorrentByID(id string) TorrentData {
	t := client.GetTorrent(id)
	Stats, Icon := GetStats(t)
	Perc := GetDownloadPercentage(t)
	var Path = ServerPath(GetTorrentPath(t))
	if f, err := os.Stat(Path); err != nil && os.IsNotExist(err) || !f.IsDir() {
		Path = strings.Replace(Path, filepath.Base(Path), "", 1)
	}
	if t != nil {
		torrent := TorrentData{
			Name:     t.Stats().Name,
			Size:     ByteCountSI(t.Stats().Bytes.Total),
			Status:   Stats,
			Magnet:   t.Stats().InfoHash.String(),
			UID:      t.ID(),
			Perc:     Perc,
			Eta:      fmt.Sprint(t.Stats().ETA),
			Speed:    GetDownloadSpeed(t),
			Progress: GetProgress(Perc),
			Icon:     Icon,
			Path:     Path,
		}
		return torrent
	}
	return TorrentData{}
}

func StopAll() {
	client.StopAll()
}

func StartAll() {
	client.StartAll()
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

func GetTorrentPath(Torr *torrent.Torrent) string {
	return "/downloads/torrents/" + Torr.ID() + "/" + Torr.Stats().Name
}

func GetAllTorrents() []TorrentData {
	var Torrents []TorrentData
	for _, t := range GetTorrents() {
		Perc := GetDownloadPercentage(t)
		Name := t.Stats().Name
		Icon := "bi bi-pause-circle"
		if Name == "" {
			Name = "fetching metadata..."
		}
		var Path = ServerPath(GetTorrentPath(t))
		if f, err := os.Stat(Path); err != nil && os.IsNotExist(err) || !f.IsDir() {
			Path = strings.Replace(Path, filepath.Base(Path), "", 1)
		}
		Stats, Icon := GetStats(t)
		Torrents = append(Torrents, TorrentData{
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
			Path:     Path,
		})
	}
	Torrents = SortAlpha(Torrents)
	for i := range Torrents {
		Torrents[i].ID = strconv.Itoa(i + 1)
	}
	return Torrents
}

func GetDownloadPercentage(torr *torrent.Torrent) string {
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

func GetStats(torr *torrent.Torrent) (string, string) {
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

func GatherSearchResults(query string) []byte {
	var tpb []SearchReq
	if query == "top100" {
		var top100 []TopTorr
		if resp, err := hClient.Get("https://tpb23.ukpass.co/apibay/precompiled/data_top100_all.json"); err != nil {
			return []byte("[]")
		} else {
			defer resp.Body.Close()
			if err = json.NewDecoder(resp.Body).Decode(&top100); err != nil {
				return []byte("[]")
			}
			for _, v := range top100 {
				tpb = append(tpb, SearchReq{
					Name:     v.Name,
					Size:     fmt.Sprint(int64(v.Size)),
					Seeders:  fmt.Sprint(int64(v.Seeders)),
					Leechers: fmt.Sprint(int64(v.Leechers)),
					InfoHash: fmt.Sprint(v.InfoHash),
				})
			}
		}
	} else {
		if resp, err := hClient.Get("https://tpb23.ukpass.co/apibay/q.php" + "?q=" + url.QueryEscape(query) + "&cat=0"); err != nil {
			return []byte("[]")
		} else {
			defer resp.Body.Close()
			if err = json.NewDecoder(resp.Body).Decode(&tpb); err != nil {
				return []byte("[]")
			}
		}
	}
	for i, t := range tpb {
		tpb[i].Magnet = "magnet:?xt=urn:btih:" + t.InfoHash + "&dn=" + url.QueryEscape(t.Name)
		tpb[i].Size = ByteCountSI(StringToInt64(t.Size))
	}
	data, _ := json.Marshal(tpb)
	return data
}

func GetLenTorrents() int {
	return len(GetTorrents())
}

func init() {
	PrepareWD()
	GetTrakers()
	client = InitClient()
}
