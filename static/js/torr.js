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
    var a =
      "<a class='list-group-item list-group-item-action flex-column align-items-start '>";
    if (IsDark()) {
      a =
        "<a class='list-group-item list-group-item-action flex-column align-items-start text-white' style='background-color: #212529'>";
    }
    a += "<div class='d-flex w-100 justify-content-between'>";
    a += "<h5 class='mb-1'>No torrents</h5>";
    a += "</div>";
    a +=
      "<p class='mb-1'>There are no torrents currently being downloaded.</p>";
    a += "</a>";
    list.append(a);
    return;
  }
  for (var i = 0; i < torrents.length; i++) {
    var torrent = torrents[i];
    var a =
      "<a class='list-group-item list-group-item-action flex-column align-items-start shadow bg-body rounded'>";
    if (IsDark()) {
      a =
        "<a class='list-group-item list-group-item-action flex-column align-items-start text-white' style='background-color: #212529' shadow bg-body rounded>";
    }
    a += "<div class='d-flex w-100 justify-content-between'>";
    speed = "";
    if (torrent.status == "Downloading") {
      speed = ` <small class="text-success">[` + torrent.speed + `]</small>`;
    }
    a +=
      "<h6 class='mb-1'><div style='word-wrap: break-word;'><b>" +
      torrent.name +
      speed +
      "</b></div></h6>";
    a += "<small><b class='text-secondary'>" + torrent.size + "</b></small>";
    a += "</div>";
    a += "<p class='mb-1 small fw-semibold'>" + torrent.status;
    if (torrent.status == "Downloading" && torrent.eta != "") {
      a += ` <small class="text-danger fw-bold">[` + torrent.eta + `]</small>`;
    }
    a += "</p>";
    a += "<div class='progress' style='height: 4px;'>";
    a +=
      "<div class='progress-bar progress-bar-striped progress-bar-animated bg-" +
      getBarColor(torrent.status) +
      "' role='progressbar' aria-valuenow='" +
      torrent.progress +
      "' aria-valuemin='0' aria-valuemax='100' style='width: " +
      torrent.progress +
      "%'></div>";
    a += "</div>";
    a += `<div class="mt-2 pt-2 border-top">`;
    a += `<div class="btn-group"><div class="btn-group me-2" role="group">`;
    a += `<button type="button" class="btn btn-primary btn-sm" data-path="${torrent.path}" onclick="btnHref(this)">Browse</button></div>`;
    a += `<div class="btn-group me-2" role="group"><button type="button" class="btn btn-danger btn-sm" onclick="removeTorrent('${torrent.uid}')">Delete</button></div>`;
    if (
      torrent.status == "Downloading" ||
      torrent.status == "Fetching Metadata"
    ) {
      a += `<div class="btn-group me-2" role="group"><button type="button" class="btn btn-warning btn-sm" onclick="pauseTorrent('${torrent.uid}')">Pause</button></div>`;
    } else if (torrent.status == "Completed") {
      a += `<div class="btn-group me-2" role="group"><button type="button" class="btn btn-success btn-sm" onclick='zipDir(this)' data-path="${torrent.path}"><i class="bi bi-file-earmark-zip"></i> Zip</button></div>`;
    } else if (torrent.status == "Stopped") {
      a += `<div class="btn-group me-2" role="group"><button type="button" class="btn btn-secondary btn-sm" onclick="resumeTorrent('${torrent.uid}')">Resume</button></div>`;
    }
    a += `</div>`;
    a += "</div></div></a>";
    list.append(a);
  }
}

function getBarColor(status) {
  if (status == "Downloading") {
    return "primary";
  } else if (status == "Completed") {
    return "success";
  } else if (status == "Paused") {
    return "warning";
  } else if (status == "Stopped") {
    return "danger";
  } else {
    return "secondary";
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
      WriteSysInfo(data);
    },
  });
}
setInterval(GetSystemInfo, 10000);

function WriteSysInfo(data) {
  sys = document.getElementById("system-info");
  sys.innerHTML = "";
  html_ =
    `<div class="card"><div class="card-body bg-light fw-bold round shadow-lg p-1"><p class="card-text text-success small">;CPU: ` +
    data.cpu +
    `, Memory: ` +
    data.mem +
    `, Disk: ` +
    data.disk +
    `, OS: ` +
    data.os +
    `, Arch: ` +
    data.arch +
    `, Downloads: ` +
    data.downloads +
    `, Folder: /downloads/torrents/</p></div></div>`;
  sys.innerHTML = html_;
}

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
