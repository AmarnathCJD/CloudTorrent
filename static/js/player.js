var src = window.location.pathname.replace("/stream", "");
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
        "settings",
        "fullscreen",
        "download",
    ],
    title: src.split("/").pop().split(".")[0],
    ratio: "16:9",
    clickToPlay: true,
    settings: ["captions", "quality", "speed"],
});

window.player = player;
