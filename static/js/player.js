var src = window.location.pathname.replace("/stream/dir", "/dir/downloads/");
var player_div = document.getElementById("player-div");

player_div.innerHTML =
    '<video id="player" playsinline controls><source src="' +
    src +
    '" type="video/mp4" /></video>';


const player = new Plyr("video", {
    controls: [
        "play-large",
        "play",
        "progress",
        "current-time",
        "mute",
        "volume",
        "captions",
        "settings",
        "airplay",
        "fullscreen",
        "download",
    ],
    title: src.split("/").pop().split(".")[0],
    ratio: "16:9",
    clickToPlay: true,
    settings: ["captions", "quality", "speed"],
});


window.player = player;