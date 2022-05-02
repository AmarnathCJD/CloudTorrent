<img src="https://camo.githubusercontent.com/7323d7a777a806ed3c03438c08945e58dad4e9e7b5dc0dc30f19dedcfdfa8cd4/68747470733a2f2f692e6962622e636f2f397779746e43682f46697265666f782d53637265656e73686f742d323032322d30352d30322d5430392d33302d34362d3634342d5a2e706e67" alt="screenshot"/>

**Cloud torrent** is a a self-hosted remote torrent client, written in Go (golang). You start torrents remotely, which are downloaded as sets of files on the local disk of the server, which are then retrievable or streamable via HTTP.

### Features

* Single binary
* Cross platform
* Embedded torrent search
* Real-time updates
* Mobile-friendly
* Fast [content server](http://golang.org/pkg/net/http/#ServeContent)

See [Future Features here](#future-features)

### Install

**Binaries**

[![Releases](https://img.shields.io/github/release/jpillora/cloud-torrent.svg)](https://github.com/jpillora/cloud-torrent/releases) [![Releases](https://img.shields.io/github/downloads/jpillora/cloud-torrent/total.svg)](https://github.com/jpillora/cloud-torrent/releases)

See [the latest release](https://github.com/jpillora/cloud-torrent/releases/latest) or download and install it now with

```
curl https://github.com/AmarnathCJD/CloudTorrent && cd CloudTorrent && go build . && ./main
```

**Source**

*[Go](https://golang.org/dl/) is required to install from source*

``` sh
$ go get -v github.com/jpillora/cloud-torrent
```

### Usage

```
clone the repository
go build 
sudo ./main```

### Future features

In summary, the core features will be:
TODO

* **File Transforms**

  During a file tranfer, one could apply different transforms against the byte stream for various effect. For example, supported transforms might include: video transcoding (using ffmpeg), encryption and decryption, [media sorting](https://github.com/jpillora/cloud-torrent/issues/4) (file renaming), and writing multiple files as a single zip file.
  
* **Automatic updates** Binary will upgrade itself, adding new features as they get released.
  
* **RSS** Automatically add torrents, with smart episode filter.

Once completed, cloud-torrent will no longer be a simple torrent client and most likely project be renamed.

### Notes
nil

Credits to {Rain Torrent]

Copyright (c) 2022 RoseLoverX
