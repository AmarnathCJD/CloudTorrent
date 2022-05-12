var table = document.getElementById("files-table");

function Log() {
    console.log("Magnet Link: ");
}

function addTorrent() {
    var Input = document.getElementById("input");
    if (Input.value == "") {
        ToastMessage("Please enter a magnet link.", "danger");
        return;
    }
    var magnet = Input.value;
    $.ajax({
        url: "/api/add",
        type: "POST",
        data: {
            magnet: magnet,
        },
        success: function (data) {
            ToastMessage("Torrent added successfully.", "success");
            Input.value = "";
        },
        error: function (data) {
            ToastMessage("Error adding torrent, " + data.responseText, "danger");
        },
    });
}

function getTorrents() {
    $.ajax({
        url: "/api/torrents",
        type: "GET",
        success: function (data) {
            updateTorrentList(data);
        },
    });
}

function updateTorrentList(data) {
    var torrents = JSON.parse(data);
    var list = $("#torrent-list");
    list.empty();
    if (torrents == null || torrents.length == 0) {
        list.append("<li class='list-group-item'>No torrents found.</li>");
        return;
    }
    for (var i = 0; i < torrents.length; i++) {
        var torrent = torrents[i];
        var a = "<a class='list-group-item list-group-item-action flex-column align-items-start '>"
        if (IsDark()) {
            a = "<a class='list-group-item list-group-item-action flex-column align-items-start text-white' style='background-color: #212529'>"
        }
        a += "<div class='d-flex w-100 justify-content-between'>"
        a += "<h5 class='mb-1'><div style='word-wrap: break-word;'><img src='https://img.icons8.com/fluency/48/000000/utorrent.png' width='38px' style='margin-right: 3px;' /> " + torrent.name + "</div></h5>"
        a += "<small>" + torrent.size + "</small>"
        a += "</div>"
        a += "<p class='mb-1 small'>" + torrent.status
        if (torrent.status == "Downloading") {
            a += " (" + torrent.speed + ")" + " ETA: " + torrent.eta
        }
        a += "</p>"
        if (torrent.status == "Downloading") {
            a += "<div class='progress'>"
            a += "<div class='progress-bar progress-bar-striped progress-bar-animated bg-" + getBarColor(torrent.progress) + "' role='progressbar' aria-valuenow='" + torrent.progress + "' aria-valuemin='0' aria-valuemax='100' style='width: " + torrent.progress + "%'>" + torrent.progress + "%</div>"
            a += "</div>"
        }
        a += `<div class="mt-2 pt-2 border-top">`
        a += `<div class="btn-group" role="group">`
        a += `<button type="button" class="btn btn-primary btn-sm" data-path="${torrent.path}" onclick="btnHref(this)">Browse</button>`
        a += `<button type="button" class="btn btn-danger btn-sm" onclick="removeTorrent('${torrent.uid}')">Delete</button>`
        if (torrent.status == "Downloading" || torrent.status == "Fetching Metadata") {
            a += `<button type="button" class="btn btn-warning btn-sm" onclick="pauseTorrent('${torrent.uid}')">Pause</button>`
        } else if (torrent.status == "Completed") {
            a += `<button type="button" class="btn btn-success btn-sm" onclick='zipDir(this)' data-path="${torrent.path}"><i class="bi bi-file-earmark-zip"></i> Zip</button>`
        } else if (torrent.status == "Stopped") {
            a += `<button type="button" class="btn btn-info btn-sm" onclick="resumeTorrent('${torrent.uid}')">Resume</button>`
        }
        a += `</div></div>`
        a += "</a>"
        list.append(a);
    }
}

function getBarColor(perc) {
    if (perc < 20) {
        return "danger";
    } else if (perc < 60) {
        return "warning";
    } else if (perc < 80) {
        return "success";
    } else {
        return "primary";
    }
}


function removeTorrent(id) {
    $.ajax({
        url: "/api/remove",
        type: "POST",
        data: {
            uid: id,
        },
        success: function (data) {
            ToastMessage("Torrent removed successfully.", "success");
            getTorrents();
        },
    });
}

function pauseTorrent(id) {
    $.ajax({
        url: "/api/pause",
        type: "POST",
        data: {
            uid: id,
        },
        success: function (data) {
            ToastMessage("Torrent paused successfully.", "primary");
            getTorrents();
        },
    });
}

function resumeTorrent(id) {
    $.ajax({
        url: "/api/resume",
        type: "POST",
        data: {
            uid: id,
        },
        success: function (data) {
            ToastMessage("Torrent resumed successfully.", "success");
            getTorrents();
        },
    });
}

function stopAll() {
    $.ajax({
        url: "/api/stopall",
        type: "POST",
        success: function (data) {
            ToastMessage("All torrents stopped.", "danger");
            getTorrents();
        },
    });
}

function startAll() {
    $.ajax({
        url: "/api/startall",
        type: "POST",
        success: function (data) {
            ToastMessage("All torrents started.", "primary");
            getTorrents();
        },
    });
}

const torr = new EventSource("/torrents/update");

torr.addEventListener(
    "torrents",
    (e) => {
        updateTorrentList(e.data);
    },
    false
);

function GetSystemInfo() {
    $.ajax({
        url: "/api/status",
        type: "GET",
        success: function (data) {
            $(".system-info").html(data);
        },
    });
}
setInterval(GetSystemInfo, 10000);

function removeAll() {
    $.ajax({
        url: "/api/removeall",
        type: "POST",
        success: function (data) {
            ToastMessage("All torrents removed successfully.", "success");
            getTorrents();
        },
    });
}

GetSystemInfo();
getTorrents();
ToastMessage("Welcome to the Torrent Manager.", "success");
