var table = document.getElementById("files-table");
var searchBox = document.getElementById("search-bar");
var Input = document.getElementById("search-input");

const DEFAULTS = {
    treshold: 2,
    maximumItems: 5,
    highlightTyped: true,
    highlightClass: 'text-primary',
};

class Autocomplete {
    constructor(field, options) {
        this.field = field;
        this.options = Object.assign({}, DEFAULTS, options);
        this.dropdown = null;

        field.parentNode.classList.add('dropdown');
        field.setAttribute('data-toggle', 'dropdown');
        field.classList.add('dropdown-toggle');

        const dropdown = ce(`<div id="drop-down" class="dropdown-menu" ></div>`);
        if (this.options.dropdownClass)
            dropdown.classList.add(this.options.dropdownClass);

        insertAfter(dropdown, field);

        this.dropdown = new bootstrap.Dropdown(field, this.options.dropdownOptions)

        field.addEventListener('click', (e) => {
            if (this.createItems() === 0) {
                e.stopPropagation();
                this.dropdown.hide();
            }
        });

        field.addEventListener('input', () => {
            if (this.options.onInput)
                this.options.onInput(this.field.value);
            this.renderIfNeeded();
        });

        field.addEventListener('keydown', (e) => {
            if (e.keyCode === 27) {
                this.dropdown.hide();
                return;
            }
        });
    }

    setData(data) {
        this.options.data = data;
    }

    renderIfNeeded() {
        if (this.createItems() > 0)
            this.dropdown.show();
        else
            this.field.click();
    }

    createItem(lookup, item) {
        let label;
        if (this.options.highlightTyped) {
            const idx = item.label.toLowerCase().indexOf(lookup.toLowerCase());
            const className = Array.isArray(this.options.highlightClass) ? this.options.highlightClass.join(' ')
                : (typeof this.options.highlightClass == 'string' ? this.options.highlightClass : '')
            label = item.label.substring(0, idx)
                + `<span class="${className}">${item.label.substring(idx, idx + lookup.length)}</span>`
                + item.label.substring(idx + lookup.length, item.label.length);
        } else
            label = item.label;
        return ce(`<button type="button" class="dropdown-item" data-value="${item.value}">${label}</button>`);
    }

    createItems() {
        const lookup = this.field.value;
        if (lookup.length < this.options.treshold) {
            this.dropdown.hide();
            return 0;
        }

        const items = this.field.nextSibling;
        items.innerHTML = '';

        let count = 0;
        for (let i = 0; i < this.options.data.length; i++) {
            const { label, value } = this.options.data[i];
            const item = { label, value };
            if (item.label.toLowerCase().indexOf(lookup.toLowerCase()) >= 0) {
                items.appendChild(this.createItem(lookup, item));
                if (this.options.maximumItems > 0 && ++count >= this.options.maximumItems)
                    break;
            }
        }

        this.field.nextSibling.querySelectorAll('.dropdown-item').forEach((item) => {
            item.addEventListener('click', (e) => {
                let dataValue = e.target.getAttribute('data-value');
                this.field.value = e.target.innerText;
                if (this.options.onSelectItem)
                    this.options.onSelectItem({
                        value: e.target.value,
                        label: e.target.innerText,
                    });
                this.dropdown.hide();
            })
        });

        return items.childNodes.length;
    }
}

function ce(html) {
    let div = document.createElement('div');
    div.innerHTML = html;
    return div.firstChild;
}

function insertAfter(elem, refElem) {
    return refElem.parentNode.insertBefore(elem, refElem.nextSibling)
}

Input.addEventListener("keyup", function (event) {
    var val = Input.value;
    if (val.split("").length > 2) {
        $.ajax({
            url: "/api/autocomplete?q=" + val,
            type: "GET",
            success: function (data) {
                var match = JSON.parse(data);
                var data = []
                for (var i = 0; i < match.length; i++) {
                    data.push({
                        label: match[i].substring(0, 22),
                        value: i
                    });
                }
                ac.setData(data);
            }
        });
    }
});

const ac = new Autocomplete(document.getElementById("search-input"));
ac.setData([]);

function FetchTorrents() {
    var Query = Input.value;
    var url = `/api/search?q=` + Query;
    var querytype = "Search results for: " + Query;
    if (Query == "") {
        url = `/api/search?q=top100`;
        querytype = "Top Trending"
    }
    $.ajax({
        url: url,
        type: "GET",
        dataType: "json",
        success: function (data) {
            var table = $("#files-table");
            table.empty();
            table.append(
                "<caption>" + querytype + "</caption><tr><th style='width: 3.66%'>ID</th><th>Name</th><th>Size</th><th>Seeds</th><th>Peers</th><th>Actions</th></tr>"
            );
            for (var i = 0; i < data.length; i++) {
                var file = data[i];
                var row = $("<tr></tr>");
                var num = i + 1;
                row.append("<td>" + num + "</td>");
                row.append("<td>" + file.name.substring(0, 52) + "..." + "</td>");
                row.append("<td>" + file.size + "</td>");
                row.append("<td>" + file.seeders + "</td>");
                row.append("<td>" + file.leechers + "</td>");
                row.append("<td><button class='btn btn-primary' onclick='addTorrent(" + file.magnet + ")'>Add</button></td>");
                table.append(row);
                if (i == 24) {
                    break;
                }
            }
        }
    })
}

function addTorrent(magnet) {
    $.ajax({
        url: "/api/add",
        type: "POST",
        data: {
            magnet: magnet,
        },
        success: function (data) {
            ToastMessage("Torrent added successfully");
        }
    })
}


FetchTorrents();