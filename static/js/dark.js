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
        ToggleTableDark()
        NavBarDark()
    } else {
        DarkIcon.innerHTML = "&#x2600;";
    }
}

function resetTheme() {
    if (darkSwitch.checked) {
        document.body.setAttribute("data-theme", "dark");
        localStorage.setItem("darkSwitch", "dark");
        DarkIcon.innerHTML = '<i class="bi bi-moon-stars-fill"></i>';
        ToggleTableDark()
        NavBarDark()
        DarkDropDown(true)
    } else {
        document.body.removeAttribute("data-theme");
        localStorage.removeItem("darkSwitch");
        DarkIcon.innerHTML = "&#x2600;";
        ToggleTableDark()
        NavBarDark()
        DarkDropDown(false)
    }
}

function DarkDropDown(mode) {
    var DropDown = document.getElementById("drop-down");
    if (DropDown !== null) {
        if (mode) {
            DropDown.classList.add("dropdown-menu-dark");
        } else {
            DropDown.classList.remove("dropdown-menu-dark");
        }
    }
}

function ToggleTableDark() {
    var table = document.getElementById("files-table");
    if (table !== null) {
        if (table.classList.contains("table-dark")) {
            table.classList.remove("table-dark");
        } else {
            table.classList.add("table-dark");
        }
    }
}

function NavBarDark() {
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
}