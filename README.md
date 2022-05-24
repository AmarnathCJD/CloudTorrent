<img src="https://raw.githubusercontent.com/AmarnathCJD/CloudTorrent/master/static/img/Screenshot_2022-05-23-12-00-44.png" alt="screenshot"/>

**Cloud torrent** is a a self-hosted remote torrent client, written in Go (golang). You start torrents remotely, which are downloaded as sets of files on the local disk of the server, which are then retrievable or streamable via HTTP.

### Features

- Single binary
- Cross platform
- Embedded torrent search
- Real-time updates
- Mobile-friendly
- Fast [content server](http://golang.org/pkg/net/http/#ServeContent)

See [Future Features here](#future-features)

### Install

**Binaries**

[![Releases](https://img.shields.io/github/release/jpillora/cloud-torrent.svg)](https://github.com/jpillora/cloud-torrent/releases) [![Releases](https://img.shields.io/github/downloads/jpillora/cloud-torrent/total.svg)](https://github.com/jpillora/cloud-torrent/releases)

See [the latest release](https://github.com/jpillora/cloud-torrent/releases/latest) or download and install it now with

```
git clone https://github.com/AmarnathCJD/CloudTorrent && cd CloudTorrent && go build . && ./main
```

**Source**

_[Go](https://golang.org/dl/) is required to install from source_

```sh
$ go get -v github.com/jpillora/cloud-torrent
```

### Usage

```
$ https://github.com/AmarnathCJD/CloudTorrent
$ go build
$ ./main
```

### Future features

In summary, the core features will be:
TODO

- **File Transforms**

  During a file tranfer, one could apply different transforms against the byte stream for various effect. For example, supported transforms might include: video transcoding (using ffmpeg), encryption and decryption, [media sorting](https://github.com/jpillora/cloud-torrent/issues/4) (file renaming), and writing multiple files as a single zip file.

- **Automatic updates** Binary will upgrade itself, adding new features as they get released.
- **RSS** Automatically add torrents, with smart episode filter.

Once completed, cloud-torrent will no longer be a simple torrent client and most likely project be renamed.

Copyright (c) 2022 RoseLoverX
