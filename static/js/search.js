var Results = document.getElementById("results");
var searchBox = document.getElementById("search");
var Button = document.getElementById("searchBtn");

function SearchTorrents() {
  var Query = searchBox.value;
  if (Query == "") {
    Query = "top100";
  }
  var url = `/api/search?q=` + Query;
  $.ajax({
    url: url,
    type: "GET",
    dataType: "json",
    success: function (data) {
      Results.innerHTML = "";
      var searchCount = document.getElementById("search-count");
      searchCount.innerHTML = `(${data.length} results)`;
      for (var i = 0; i < data.length; i++) {
        var file = data[i];
        file.magnet = file.magnet.replace(/\s/g, "+");
        var a = `<a class="list-group-item list-group-item-action flex-column align-items-start ">`;
        if (IsDark()) {
          a = `<a class="list-group-item list-group-item-action flex-column align-items-start text-white" style="background-color: #212529">`;
        }
        a += `<div class="d-flex w-100 justify-content-between">`;
        a +=
          `<h6 class="mb-1">1. ` +
          file.name +
          `<img width='25px' style='margin-right: 5px; margin-left: 4px;' src="https://img.icons8.com/external-yogi-aprelliyanto-glyph-yogi-aprelliyanto/64/fa314a/external-magnet-marketing-and-seo-yogi-aprelliyanto-glyph-yogi-aprelliyanto.png"/></h5>`;
        a += `<small>` + file.size + `</small>`;
        a += `</div>`;
        a += `<p class="mb-1">Seeds: `;
        if (file.seeders > 2000) {
          a += `<b class="text-success">` + file.seeders + `</b>`;
        } else if (file.seeders > 1000) {
          a += `<b class="text-warning">` + file.seeders + `</b>`;
        } else if (file.seeders > 500) {
          a += `<b class="text-danger">` + file.seeders + `</b>`;
        } else if (file.seeders < 5) {
          a += `<b class="text-danger">` + file.seeders + `</b>`;
          a += `<i class="bi bi-exclamation" style="color: red;"></i>`;
        } else {
          a += `<b class="text-secondary">` + file.seeders + `</b>`;
        }
        a += `, Leeches: <b class='text-muted'>` + file.leechers + `</b></p>`;
        a += `<div class="mt-2 pt-2 border-top">`;
        a += `<div class="btn-group" role="group">`;
        a += `<button class="btn btn-primary btn-sm" onclick="addTorrent('${file.magnet}')">Download <i class="bi bi-download"></i></button>`;
        a += `<button class="btn btn-danger btn-sm" onclick="copyToClipboard(this)" data-url="${file.magnet}"><i class="bi bi-clipboard"></i></button>`;
        a += `</div></div>`;
        a += `</a>`;
        Results.innerHTML += a;
      }
    },
  });
}

SearchTorrents();

Button.addEventListener("click", SearchTorrents);

function addTorrent(magnet) {
  $.ajax({
    url: "/api/add",
    type: "POST",
    data: {
      magnet: magnet,
    },
    success: function (data) {
      ToastMessage("Torrent added successfully", "success");
    },
    error: function (data) {
      if (data.status == 500) {
        ToastMessage("Torrent already added", "warning");
      }
    },
  });
}
