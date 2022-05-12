function ToastMessage(message, bg) {
    document.querySelector(".toast-container").innerHTML =
        `<div class="toast align-items-center text-white bg-` +
        bg +
        ` border-0" role="alert" aria-live="assertive" aria-atomic="true" id="toast-main" style="position: absolute; top: 0; right: 0;"><div class="d-flex"><div class="toast-body">` +
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