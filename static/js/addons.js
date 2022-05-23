var NumToasts = 0;

function ToastMessage(message, bg) {
    NumToasts++;
    document.querySelector(".toast-container").innerHTML +=
        `<div class="toast align-items-center text-white bg-` +
        bg +
        ` border-0" role="alert" aria-live="assertive" aria-atomic="true" id="toast-main-${NumToasts}"><div class="toast-header">
        <img src="https://img.icons8.com/material-rounded/24/fa314a/appointment-reminders.png" class="rounded me-2" alt="notif"><strong class="me-auto">Cloud Torrent</strong><small>Just Now</small><button type="button" class="btn-close" data-bs-dismiss="toast" aria-label="Close"></button></div><div class="toast-body">` +
        message +
        `</div></div>`;
    var Toast = bootstrap.Toast.getOrCreateInstance(
        document.getElementById("toast-main-" + NumToasts)
    );
    Toast.options = {
        delay: 5000,
        autohide: true,
    };
    Toast.show();
    var PrevToast = document.getElementById("toast-main-" + (NumToasts - 1));
    if (PrevToast) {
        PrevToast.remove();
    }
}

function copyToClipboard(e) {
    var copyText = $(e).data("url");
    navigator.clipboard.writeText(copyText).then(
        function () {
            ToastMessage("Copied to clipboard", "success");
        },
        function () {
            ToastMessage("Failed to copy to clipboard", "error");
        }
    );
}

function zipDir(e) {
    var path = e.getAttribute("data-path");
    $.ajax({
        url: "/api/zip/" + path,
        type: "GET",
        dataType: "json",
        success: function (data) {
            ToastMessage("Zipped Directory", "success");
            e.outerHTML =
                `<button type="button" class="btn btn-primary btn-sm" onclick='downloadStart(this)' data-path="` +
                data.file +
                `">Download</button>`;
        },
    });
}

function btnHref(e) {
    var path = $(e).data("path");
    window.location.href = path;
}

function ToClipboard(id) {
    elem = document.getElementById("btn-" + id);
    data = elem.getAttribute("data-clipboard-text");
    navigator.clipboard.writeText(data).then(
        function () {
            ToastMessage("Copied to clipboard", "success");
        },
        function () {
            ToastMessage("Failed to copy to clipboard", "error");
        }
    );
}
