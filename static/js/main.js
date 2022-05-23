var Audd = new Audio();
var table = document.getElementById("files-table");
var CurrentPlaying = [null, null];

function updateDirList(url) {
    if (url == null) {
        url = "/dir/" + window.location.pathname.replace("/downloads", "")
    }
    $.ajax({
        url: url,
        type: "GET",
        dataType: "json",
        success: function (data) {
            var dirList = document.getElementById("dir-list");
            dirList.innerHTML = "";
            if (data.length === 0) {
                data.push({ "name": "No Files...", "path": "", "type": "dir" });
            }
            for (var i = 0; i < data.length; i++) {
                var dir = data[i];
                var a = `<a class='list-group-item list-group-item-action flex-column align-items-start'>`;
                if (IsDark()) {
                    a = `<a class='list-group-item list-group-item-action flex-column align-items-start text-white' style='background-color: #212529'>`;
                }
                a += "<div class='d-flex w-100 justify-content-between'>";
                if (dir.is_dir == "true") {
                    a += `<h5 class="mb-1"><img src="https://img.icons8.com/color/48/000000/folder-invoices--v1.png" width="40px" style="margin-right: 3px;"/> ${dir.name}</h5>`;
                } else {
                    a += `<h5 class="mb-1"><img src="https://img.icons8.com/${dir.icon}" width="40px" style="margin-right: 3px;"/> ${dir.name}${dir.ext}</h5>`;
                }
                a += "<small>" + dir.size + "</small>";
                a += "</div>";
                a += "<p class='mb-1 small'>" + dir.type + "</p>";
                a += `<div id="player_id_${i}"></div>`;
                a += `<div class="mt-2 pt-2 border-top">`;
                a += `<div class="btn-group" role="group">`;
                if (dir.is_dir == "true") {
                    a += `<button type="button" class="btn btn-primary btn-sm" data-path="${dir.path}" onclick="zipDir(this)">Zip & Download</button>`;
                    a += `<button type="button" class="btn btn-secondary btn-sm" data-path="${dir.path}" onclick="btnHref(this)">Browse</button>`;
                } else {
                    a += `<button type="button" class="btn btn-primary btn-sm" data-path="${dir.path}" onclick="downloadStart(this)" >Download</button>`;
                    if (dir.type == "Video") {
                        a += `<button type="button" class="btn btn-warning btn-sm" data-src="${dir.path}" data-id="${i}" onclick="playVideo(this)">Play</button>`;
                        if (dir.ext == ".mkv") {
                            a += `<button type="button" class="btn btn-danger btn-sm" data-src="${dir.path}" data-id="${i}" onclick="playAudio(this)">-> MP4</button>`;
                        }
                    } else if (dir.type == "Audio") {
                        a += `<button type="button" class="btn btn-danger btn-sm" data-path="${dir.path}" onclick="playAudio(this)">Play</button>`;
                    } else if (dir.type == "Image") {
                        a += `<button type="button" class="btn btn-success btn-sm" data-src="${dir.path}" onclick="showImage(this)">View</button>`;
                    }
                }
                a += `<button type='button' class='btn btn-danger btn-sm' onclick='deleteFile(this)' data-path='${dir.path}'>Delete</button>`;
                a += `<button type='button' class='btn btn-success btn-sm' data-url='${window.location.host}${dir.path}' onclick='copyToClipboard(this)'><i class="bi bi-clipboard-plus"></i></button>`;
                a += "</div></div>";
                a += "</a>";
                dirList.innerHTML += a;
            }
        },
        error: function (err) {
            console.log(err);
        },
    });
}


function downloadStart(e) {
    var path = e.getAttribute("data-path");
    var name = path.split("/").pop();
    var a = document.createElement("a");
    a.href = path;
    a.download = name;
    a.click();
}

function playAudio(url, uid) {
    Audd.src = url;
    Audd.play();
    var btn = document.getElementById(uid);
    btn.outerHTML =
        `<button type="button" class="btn btn-danger" onclick='pauseAudio("` +
        url +
        `", this.id)' id="` +
        uid +
        `"><i class="bi bi-pause-circle"></i></button>`;
    ToastMessage("Playing " + url.split("/").pop(), "primary");
    if (CurrentPlaying[0] != null && CurrentPlaying[0] != uid) {
        var btn = document.getElementById(CurrentPlaying[0]);
        btn.outerHTML =
            `<button type="button" class="btn btn-danger" onclick='playAudio("` +
            CurrentPlaying[1] +
            `", this.id)' id="` +
            CurrentPlaying[0] +
            `"><i class="bi bi-play-circle"></i></button>`;
    }
    CurrentPlaying = [uid, url];
    Audd.onended = function () {
        btn.outerHTML =
            `<button type="button" class="btn btn-danger" onclick='playAudio("` +
            url +
            `", this.id)' id="` +
            uid +
            `"><i class="bi bi-play-circle"></i></button>`;
        CurrentPlaying = [null, null];
    };
}

function pauseAudio(url, uid) {
    Audd.pause();
    var btn = document.getElementById(uid);
    btn.outerHTML =
        `<button type="button" class="btn btn-danger" onclick='playAudio("` +
        url +
        `", this.id)' id="` +
        uid +
        `"><i class="bi bi-play-circle"></i></button>`;
    ToastMessage("Paused Audio", "primary");
}

function showImage(e) {
    url = $(e).data("src");
    var modal = document.getElementById("main-modal");
    var img = document.getElementById("main-modal-img");
    var captionText = document.getElementById("main-modal-caption");
    var closeSpan = document.getElementsByClassName("close")[0];
    modal.style.display = "block";
    img.src = url;
    captionText.innerHTML = url;
    closeSpan.onclick = function () {
        modal.style.display = "none";
    };
}


function playVideo(e) {
    var url = $(e).data("src");
    window.location.href = "/stream/" + url;
}

function deleteFile(e) {
    var current_url = window.location.href;
    var path = $(e).data("path");
    var name = path.split("/").pop();
    var r = confirm("Are you sure you want to delete " + name + "?");
    if (r !== true) {
        return;
    }
    path = path.replace("/dir/", "");
    console.log(path);
    $.ajax({
        url: "/api/deletefile/" + path,
        type: "GET",
        success: function (data) {
            ToastMessage("Deleted " + name, "danger");
            updateDirList(current_url.replace("downloads", "dir"));
        },
        error: function (err) {
            ToastMessage("Failed to delete " + err.responseText, "danger");
        },
    });
}

function backButton() {
    var path = window.location.pathname;
    if (path.length > 1) {
        var newPath = path.substring(0, path.lastIndexOf("/"));
        window.location.href = newPath;
    }
}

updateDirList();
if (window.location.pathname == "/downloads/") {
    ToastMessage("Welcome to File Manager", "primary");
}

const handleFileUpload = (e) => {
    const files = e.target.files;
    const uploadData = new FormData();
    uploadData.append("file", files[0]);
    uploadData.append("path", window.location.pathname);
    fetch("/api/upload", {
        method: "POST",
        body: uploadData,
    }).then((response) => {
        ToastMessage("Uploaded " + files[0].name, "success");
        UpdateDir();
    });
};

document.querySelector("#file").addEventListener("change", (event) => {
    handleFileUpload(event);
});

function CreateFolder() {
    var name = prompt("Enter Folder Name");
    if (name != null) {
        $.ajax({
            url: "/api/create/" + window.location.pathname + name,
            type: "GET",
            success: function (data) {
                ToastMessage("Created " + name, "primary");
                updateDirList();
            },
            error: function (data) {
                ToastMessage("Error: " + data.responseText, "danger");
            },
        });
    } else {
        ToastMessage("Name cannot be null", "danger");
    }
}
