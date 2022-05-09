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
            a = "<a class='list-group-item list-group-item-action flex-column align-items-start bg-secondary text-white'>"
        }
        a += "<div class='d-flex w-100 justify-content-between'>"
        a += "<h5 class='mb-1'>" + torrent.name + "</h5>"
        a += "<small>" + torrent.size + "</small>"
        a += "</div>"
        a += "<p class='mb-1'>" + torrent.status
        if (torrent.status == "Downloading") {
            a += " (" + torrent.speed + ")" + " ETA: " + torrent.eta
        }
        a += "</p>"
        if (torrent.status == "Downloading") {
            a += "<div class='progress'>"
            a += "<div class='progress-bar progress-bar-striped progress-bar-animated' role='progressbar' aria-valuenow='" + torrent.progress + "' aria-valuemin='0' aria-valuemax='100' style='width: " + torrent.progress + "%'>" + torrent.progress + "%</div>"
            a += "</div>"
        }
        a += `<div class="mt-2 pt-2 border-top">`
        a += `<button type="button" class="btn btn-primary btn-sm">Browse</button>`
        a += `<button type="button" class="btn btn-danger btn-sm">Delete</button>`
        if (torrent.status == "Downloading" || torrent.status == "Fetching Metadata") {
            a += `<button type="button" class="btn btn-warning btn-sm" onclick="pauseTorrent('${torrent.uid}')">Pause</button>`
        } else if (torrent.status == "Completed") {
            a += `<button type="button" class="btn btn-success btn-sm"><i class="bi bi-file-earmark-zip"></i> Zip</button>`
        } else if (torrent.status == "Stopped") {
            a += `<button type="button" class="btn btn-secondary btn-sm" onclick="resumeTorrent('${torrent.uid}')">Resume</button>`
        }
        a += `</div>`
        a += "</a>"
        list.append(a);
    }
}

function updateTorrents(data) {
    var torrents = JSON.parse(data);
    var table = $("#files-table");
    table.empty();
    table.append(
        "<caption>Active Torrents</caption><tr><th style='width: 3.66%'>ID</th><th>Name</th><th>Size</th><th>Status</th><th>ETA</th><th>Download Speed</th><th>Actions</th></tr>"
    );
    if (torrents == null || torrents.length == 0) {
        console.log("No torrents");
        table.append("<tr><td colspan='7'>No active torrents.</td></tr>");
        return;
    }
    for (var i = 0; i < torrents.length; i++) {
        var torrent = torrents[i];
        var row = $("<tr></tr>");
        row.append("<td>" + torrent.id + "</td>");
        var progress_bb =
            `<div class="progress" style="height: 10px; border-top: 4px;"><div class="progress-bar progress-bar-striped progress-bar-animated bg-danger" role="progressbar" aria-valuenow="75" aria-valuemin="0" aria-valuemax="100" style="width: ` +
            torrent.progress +
            `%"></div></div>`;
        if (torrent.progress >= 85) {
            progress_bb =
                `<div class="progress" style="height: 10px; border-top: 4px;"><div class="progress-bar progress-bar-striped progress-bar-animated bg-success" role="progressbar" aria-valuenow="75" aria-valuemin="0" aria-valuemax="100" style="width: ` +
                torrent.progress +
                `%"></div></div>`;
        } else if (torrent.progress <= 35) {
            progress_bb =
                `<div class="progress" style="height: 10px; border-top: 4px;"><div class="progress-bar progress-bar-striped progress-bar-animated bg-warning" role="progressbar" aria-valuenow="75" aria-valuemin="0" aria-valuemax="100" style="width: ` +
                torrent.progress +
                `%"></div></div>`;
        }
        row.append("<td>" + torrent.name + progress_bb + `</td>`);
        row.append("<td>" + torrent.size + "</td>");
        row.append("<td>" + torrent.status + "</td>");
        row.append("<td>" + torrent.eta + "</td>");
        row.append("<td>" + torrent.speed + "</td>");
        var actionbutton = "";
        if (torrent.status == "Downloading") {
            actionbutton =
                `<button class="btn btn-danger" onclick="pauseTorrent('` +
                torrent.uid +
                `')"><i class="bi bi-pause-fill"></i></button>`;
        } else {
            actionbutton =
                `<button class="btn btn-success" onclick="resumeTorrent('` +
                torrent.uid +
                `')"><i class="bi bi-play-fill"></i></button>`;
        }
        row.append(
            "<td><div class='btn-group'> <button class='btn btn-danger' onclick='removeTorrent(\"" +
            torrent.uid +
            "\")'><i class='bi bi-x-circle'></i></button><a href='" +
            torrent.path.replace("/downloads", "") +
            "'><button class='btn btn-warning'><i class='bi bi-folder-plus'></i></button></a>" +
            actionbutton +
            "</div></td>"
        );
        table.append(row);
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
