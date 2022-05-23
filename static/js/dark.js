var darkSwitch = document.getElementById("darkSwitch");
var DarkIcon = document.getElementById("dark-icon");

window.addEventListener("load", function () {
    if (darkSwitch) {
        initTheme();
        darkSwitch.addEventListener("change", function () {
            resetTheme();
        });
    }
});

function initTheme() {
    var darkThemeSelected =
        localStorage.getItem("darkSwitch") !== null &&
        localStorage.getItem("darkSwitch") === "dark";
    darkSwitch.checked = darkThemeSelected;
    darkThemeSelected
        ? document.body.setAttribute("data-theme", "dark")
        : document.body.removeAttribute("data-theme");
    if (darkThemeSelected) {
        DarkIcon.innerHTML = '<i class="bi bi-moon-stars-fill"></i>';
        toggleDarkElements();
    } else {
        DarkIcon.innerHTML = "&#x2600;";
    }
}

function resetTheme() {
    if (darkSwitch.checked) {
        document.body.setAttribute("data-theme", "dark");
        localStorage.setItem("darkSwitch", "dark");
        DarkIcon.innerHTML = '<i class="bi bi-moon-stars-fill"></i>';
        toggleDarkElements();
    } else {
        document.body.removeAttribute("data-theme");
        localStorage.removeItem("darkSwitch");
        DarkIcon.innerHTML = "&#x2600;";
        toggleDarkElements();
    }
}

function toggleDarkElements() {
    var DropDown = document.getElementById("drop-down");
    if (DropDown !== null) {
        if (!DropDown.classList.contains("dropdown-menu-dark")) {
            DropDown.classList.add("dropdown-menu-dark");
        } else {
            DropDown.classList.remove("dropdown-menu-dark");
        }
    }
    var navbar = document.getElementById("main-nav");
    if (navbar !== null) {
        if (navbar.classList.contains("navbar-dark")) {
            navbar.classList.remove("navbar-dark");
            navbar.classList.add("navbar-light");
        } else {
            navbar.classList.remove("navbar-light");
            navbar.classList.add("navbar-dark");
        }
    }
    var list = document.getElementById("torrent-list");
    if (list === null) {
        list = document.getElementById("dir-list");
    }
    if (list === null) {
        list = document.getElementById("results");
    }
    if (list !== null) {
        for (var i = 0; i < list.children.length; i++) {
            if (list.children[i].classList.contains("text-white")) {
                list.children[i].classList.remove("text-white");
                list.children[i].style.backgroundColor = "white";
            } else {
                list.children[i].classList.add("text-white");
                list.children[i].style.backgroundColor = "#212529";
            }
        }
    }
    var s = document.getElementById("system-info");
    if (s !== null) {
        if (s.style.backgroundColor === "white") {
            s.style.backgroundColor = "#212529";
        } else {
            s.style.backgroundColor = "white";
        }
    }
}

function IsDark() {
    return document.body.getAttribute("data-theme") === "dark";
}