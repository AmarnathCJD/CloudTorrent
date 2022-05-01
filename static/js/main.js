var Audd = new Audio();
var table = document.getElementById("files-table");
var CurrentPlaying = [null, null];


function UpdateDir() {
    $.ajax({
        url: "/dir/" + window.location.pathname,
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
                    console.log(file.path);
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
                        "<a href='/dir" +
                        AbsPath.replace("dir", "downloads") +
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
                if (file.is_dir == "true") {
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
                        `<button type="button" class="btn btn-danger" onclick='playAudio("/dir/` +
                        file.path +
                        `", this.id)' id="` +
                        i +
                        `"><i class="bi bi-filetype-mp3"></i></button></div>`;
                } else if (
                    file.ext == ".jpg" ||
                    file.ext == ".png" ||
                    file.ext == ".jpeg"
                ) {
                    GroupButtons +=
                        `<button type="button" class="btn btn-warning" onclick='showImage("/dir/` +
                        file.path +
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
    });
}

function ToastMessage(message) {
    document.querySelector(".toast-container").innerHTML =
        `<div class="toast align-items-center text-white bg-primary border-0" role="alert" aria-live="assertive" aria-atomic="true" id="toast-main"><div class="d-flex"><div class="toast-body">` +
        message +
        `</div><button type="button" class="btn-close btn-close-white me-2 m-auto" data-bs-dismiss="toast" aria-label="Close"></button></div></div>`;
    var Toast = bootstrap.Toast.getOrCreateInstance(
        document.getElementById("toast-main")
    );
    Toast.options = {
        delay: 5000,
        autohide: true,
    };
    Toast.show();
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
    ToastMessage("Playing " + url.split("downloads/")[1]);
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

UpdateDir();
ToastMessage("Welcome to the File Manager!");