var Audd = new Audio();
var table = document.getElementById("files-table");
var CurrentPlaying = [null, null];

function UpdateDir() {
    $.ajax({
        url: "/dir/" + window.location.pathname.replace("/downloads", ""),
        type: "GET",
        dataType: "json",
        success: function (data) {
            var table = $("#files-table");
            table.empty();
            table.append(
                "<caption>List of Files</caption><tr><th style='width: 3.66%'>ID</th><th>File</th><th>Size</th><th>Type</th><th >Actions</th></tr>"
            );
            for (var i = 0; i < data.length; i++) {
                var file = data[i];
                var row = $("<tr></tr>");
                var num = i + 1;
                row.append("<td>" + num + "</td>");
                var FileName = file.name;
                if (file.name.length > 45) {
                    FileName = file.name.substring(0, 43) + "...";
                }
                var AbsPath = file.path.replace("/downloads", "/dir");
                if (file.is_dir == "true") {
                    var faclass =
                        "<i style='font-size: 1rem; color: #f3da35;' class='bi bi-folder'></i>";
                    row.append(
                        "<td><a href='" +
                        file.path +
                        "/" +
                        "'><b>" +
                        FileName +
                        "</b></a> " +
                        faclass +
                        "</td>"
                    );
                } else {
                    var faclass =
                        "<i style='font-size: 1rem; color: " +
                        file.color +
                        ";' class='" +
                        file.class +
                        "'></i>    ";
                    row.append(
                        "<td class='text-wrap'>" +
                        "<a href='" +
                        file.path.replace("downloads", "dir") +
                        "'><b>" +
                        FileName +
                        "</b></a> " +
                        faclass +
                        "</td>"
                    );
                }
                row.append("<td>" + file.size + "</td>");
                if (file.is_dir == "true") {
                    AbsPath = file.path;
                    row.append("<td>Folder</td>");
                } else {
                    row.append("<td>" + file.type + "</td>");
                }
                var GroupButtons = `<td style='width: 10.66%'><div class="btn-group" role="group"><button type="button" class="btn btn-danger `;
                if (file.is_dir == "true" && file.size !== "0 B") {
                    GroupButtons += `disabled`;
                }
                GroupButtons +=
                    `" onclick="deleteFile('` +
                    file.name +
                    file.ext +
                    `')"><i class="bi bi-x-circle"></i></button><a href="/dir` +
                    AbsPath.replace("dir", "downloads") +
                    `" download="` +
                    file.name +
                    file.ext +
                    `"><button type="button" class="btn btn-primary"><i class="bi bi-download"></i></button></a>`;
                if (file.is_dir == "true") {
                    GroupButtons +=
                        `<a href='` +
                        file.path +
                        `/` +
                        `'><button type="button" class="btn btn-success"><i class="bi bi-folder-plus"></i></button></a></div>`;
                } else if (
                    file.ext == ".pdf" ||
                    file.ext == ".txt" ||
                    file.ext == ".docx" ||
                    file.ext == ".doc"
                ) {
                    GroupButtons += `<button type="button" class="btn btn-success"><i class="bi bi-file-pdf"></i></button></div>`;
                } else if (
                    file.ext == ".mp4" ||
                    file.ext == ".mkv" ||
                    file.ext == ".avi" ||
                    file.exit == ".webm"
                ) {
                    GroupButtons +=
                        `<a href="/stream/` +
                        AbsPath +
                        `"><button type="button" class="btn btn-warning""><i class="bi bi-play-circle"></i></button></a></div>`;
                } else if (
                    file.ext == ".mp3" ||
                    file.ext == ".wav" ||
                    file.ext == ".flac"
                ) {
                    GroupButtons +=
                        `<button type="button" class="btn btn-danger" onclick='playAudio("` +
                        AbsPath +
                        `", this.id)' id="` +
                        i +
                        `"><i class="bi bi-filetype-mp3"></i></button></div>`;
                } else if (
                    file.ext == ".jpg" ||
                    file.ext == ".png" ||
                    file.ext == ".jpeg"
                ) {
                    GroupButtons +=
                        `<button type="button" class="btn btn-warning" onclick='showImage("` +
                        AbsPath +
                        `", this.id)' id="` +
                        i +
                        `"><i class="bi bi-image"></i></button></div>`;
                } else {
                    GroupButtons += `</div>`;
                }
                row.append(GroupButtons + "</td>");
                table.append(row);
            }
        },
        error: function (data) {
            ToastMessage("Error: " + data.responseText, "danger");
        },
    });
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
    }

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
    ToastMessage("Paused Audio");
}

function showImage(url, uid) {
    var modal = document.getElementById("main-modal");
    var img = document.getElementById("main-modal-img");
    var captionText = document.getElementById("main-modal-caption");
    var closeSpan = document.getElementsByClassName("close")[0];
    modal.style.display = "block";
    img.src = url;
    captionText.innerHTML = url.split("downloads/")[1];
    closeSpan.onclick = function () {
        modal.style.display = "none";
    };
}

function deleteFile(name) {
    $.ajax({
        url: "/delete/" + window.location.pathname + name,
        type: "GET",
        success: function (data) {
            ToastMessage("Deleted " + name);
            UpdateDir();
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

UpdateDir();
if (window.location.pathname == "/downloads/") {
    ToastMessage("Welcome to File Manager", "primary");
}

const handleFileUpload = (e) => {
    const files = e.target.files;
    const uploadData = new FormData();
    uploadData.append("file", files[0]);
    uploadData.append("path", window.location.pathname);
    fetch(
        "/api/upload",
        {
            method: "POST",
            body: uploadData,
        }
    )
        .then((res) => res.json())
        .then((data) => {
            UpdateDir();
        })
        .catch((err) => {
            ToastMessage("Upload Failed", "danger");
        });
}

document.querySelector('#file').addEventListener('change', event => {
    handleFileUpload(event)
})

function CreateFolder() {
    var name = prompt("Enter Folder Name");
    if (name != null) {
        $.ajax({
            url: "/api/create/" + window.location.pathname + name,
            type: "GET",
            success: function (data) {
                ToastMessage("Created " + name, "primary");
                UpdateDir();
            },
            error: function (data) {
                ToastMessage("Error: " + data.responseText, "danger");
            },
        });
    } else {
        ToastMessage("Name cannot be null", "danger");
    }
}