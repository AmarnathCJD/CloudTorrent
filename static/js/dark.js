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
        DarkDropDown(true)
    } else {
        document.body.removeAttribute("data-theme");
        localStorage.removeItem("darkSwitch");
        DarkIcon.innerHTML = "&#x2600;";
        ToggleTableDark()
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