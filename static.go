package main

const (
	downloads = `<!DOCTYPE html>
    <html>
    
    <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Downloads Index</title>
        <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
        <style>
            body {
                font-family: Arial, Helvetica, sans-serif;
                font-size: 14px;
                line-height: 1.5;
                color: #333;
                background-color: #f5f5f5;
                margin: 0;
                padding: 0;
            }
    
            .header {
                overflow: hidden;
                background-color: #f1f1f1;
                padding: 20px 10px;
            }
    
            .header a {
                float: left;
                color: black;
                text-align: center;
                padding: 12px;
                text-decoration: none;
                font-size: 18px;
                line-height: 25px;
                border-radius: 4px;
            }
    
    
            .header a.logo {
                font-size: 25px;
                font-weight: bold;
            }
    
            .header a:hover {
                background-color: #ddd;
                color: black;
            }
    
            .header a.active {
                background-color: dodgerblue;
                color: white;
            }
    
            .header-right {
                float: right;
            }
    
            a {
                color: #0088cc;
                text-decoration: none;
            }
    
            a:hover {
                color: #005580;
            }
    
            .container {
                width: 960px;
                margin: 0 auto;
                box-shadow: 0 5px 10px -5px rgba(0, 0, 0, 0.5);
            }
    
            .content {
                padding: 10px 0;
            }
    
            .content h2 {
                margin: auto;
                font-size: 25px;
                font-weight: normal;
                line-height: 1.5;
            }
    
            table {
                width: 100%;
                border-collapse: collapse;
                border-spacing: 0;
                border: 1px solid #ddd;
            }
    
            table tr {
                border-bottom: 1px solid #ddd;
            }
    
            table th,
            table td {
                text-align: left;
                padding: 12px;
            }
    
            table th {
                background-color: #f5f5f5;
            }
    
            table tr:nth-child(even) {
                background-color: #f2f2f2;
            }
    
            table th:first-child {
                border-top: 0;
                background-color: solid #f5f5f5;
            }
    
            table tr:hover {
                background-color: #ddd;
            }
    
            table p {
                margin: 0;
                font-size: 14px;
                line-height: 1.5;
            }
    
            footer {
                position: absolute;
                bottom: 0;
                left: 0;
                right: 0;
                background: #111;
                height: auto;
                width: 100vw;
                padding-top: 40px;
                color: #fff;
            }
    
            .footer-content {
                display: flex;
                align-items: center;
                justify-content: center;
                flex-direction: column;
                text-align: center;
            }
    
            button {
                background-color: DodgerBlue;
                border: none;
                color: white;
                text-align: center;
                text-decoration: none;
                display: inline-block;
                padding: 12px 16px;
                font-size: 16px;
                margin: 4px 2px;
                cursor: pointer;
                border-radius: 6px;
            }
    
            button span {
                cursor: pointer;
                display: inline-block;
                position: relative;
                transition: 0.5s;
            }
    
            button span:after {
                content: '\00bb';
                position: absolute;
                opacity: 0;
                top: 0;
                right: -20px;
                transition: 0.5s;
            }
    
            button:hover span {
                padding-right: 25px;
                opacity: 0.5;
            }
    
            button:hover {
                background-color: RoyalBlue;
            }
    
            button:hover span:after {
                opacity: 1;
                right: 0;
            }
        </style>
    </head>
    
    <body>
        <div class="container">
            <div class="header"
                style="overflow: hidden;background-color: #f1f1f1;padding: 20px 10px;border-bottom: 3px solid #ffcc00;">
                <a href="#default" class="logo">Cloud Torrent</a>
                <div class="header-right">
                    <a href="/">Torrents</a>
                    <a href="/torrents/search?query=top100">Search</a>
                    <a href="/downloads" class="active">Files</a>
                    <a href="/">Home</a>
                </div>
            </div>
            <div class="content">
                <table>
                    <tr>
                        <th>File</th>
                        <th>Size</th>
                        <th>Type</th>
                        <th>Date</th>
                    </tr>
                    {{files}}
                </table>
            </div>
            <button onclick="history.back()"><span>Go Back</span></button>
            <button onclick="window.location.href='/'"><span>Add Magnet</span></button>
            <button onclick="window.location.href='/downloads'"><i class="fa fa-home"></i></button>
        </div>
    </body>
    
    </html>
`

	torrents = `<!DOCTYPE html>
    <html>
    
    <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Torrents Cloud</title>
        <style>
            body {
                font-family: Arial, Helvetica, sans-serif;
                font-size: 14px;
                line-height: 1.5;
                color: #333;
                background-color: #f5f5f5;
                margin: 0;
                padding: 0;
            }
    
            .header {
                overflow: hidden;
                background-color: #f1f1f1;
                padding: 20px 10px;
                border-bottom: 3px solid #ffcc00;
            }
    
            .header a {
                float: left;
                color: black;
                text-align: center;
                padding: 12px;
                text-decoration: none;
                font-size: 18px;
                line-height: 25px;
                border-radius: 4px;
    
            }
    
            .header a.logo {
                font-size: 25px;
                font-weight: bold;
            }
    
            .header a:hover {
                background-color: darkgrey;
                color: black;
                transition: 0.5s;
            }
    
            .header a.active {
                background-color: dodgerblue;
                color: white;
            }
    
            .header-right {
                float: right;
            }
    
    
            add-magnet {
                background-color: MediumSeaGreen;
                border: none;
                color: white;
                padding: 15px 32px;
                text-align: center;
                text-decoration: none;
                display: inline-block;
                font-size: 16px;
                margin: 4px 2px;
                cursor: pointer;
                border-radius: 6px;
            }
    
            add-magnet:hover {
                opacity: 0.8;
            }
    
            magnet.input {
                width: 100%;
                height: 40px;
                border: 1px solid #ccc;
                border-radius: 4px;
                padding: 0 10px;
                font-size: 14px;
                line-height: 1.5;
                margin-top: 10px;
            }
    
            magnet.input:focus {
                outline: none;
            }
    
            .container {
                width: 960px;
                margin: 0 auto;
                box-shadow: 0 5px 10px -5px rgba(0, 0, 0, 0.5);
            }
    
            .content {
                padding: 10px 0;
            }
    
            .content h2 {
                margin: auto;
                font-size: 25px;
                font-weight: normal;
                line-height: 1.5;
            }
    
            input[type=text],
            select {
                width: 95%;
                padding: 12px 20px;
                margin: 19px 26px;
                margin-bottom: 8px;
                display: inline-block;
                border: 1px solid #ccc;
                border-radius: 4px;
                box-sizing: border-box;
                background: url(https://img.icons8.com/ios-glyphs/30/000000/magnet-therapy.png) no-repeat scroll 3px 3px;
                padding-left: 37px;
            }
    
            input[type=submit] {
                width: 12%;
                height: 35px;
                background-color: #3c5abd;
                color: white;
                padding: 7px 2px;
                border: none;
                border-radius: 4px;
                cursor: pointer;
                margin-bottom: 10px;
            }
    
            input[type=submit]:hover {
                background-color: #1630a3;
            }
    
            .wrapper {
                text-align: center;
            }
    
            table {
                border-collapse: collapse;
                border-spacing: 0;
                border: 1px solid #ddd;
                width: 100%;
            }
            table.bodywrapcenter>tr>td {
                width: 100%;
                float: left;
            }
            
    
            tr {
                border-bottom: 1px solid #ddd;
            }
    
            th,
            td {
                text-align: left;
                padding: 12px;
            }
    
            th {
                background-color: #f5f5f5;
            }
    
            tr:nth-child(even) {
                background-color: #f2f2f2;
            }
    
            tr:hover {
                background-color: #ddd;
            }
    
            th.id {
                width: 5%;
            }
    
            th.name {
                width: 20%;
            }
    
            th.name a {
                color: #333;
            }
    
            th.name a:hover {
                color: #0088cc;
            }
    
            th.size {
                width: 10%;
            }
    
            th.date {
                width: 15%;
            }
    
            th.magnet {
                width: 20%;
            }
    
            th.action {
                width: 10%;
            }
            .action {
                width: 20%;
            }
    
            .action a {
                color: #fff;
                background-color: #0088cc;
                padding: 5px 10px;
                border-radius: 4px;
                text-decoration: none;
            }
            .action btn {
                color: #fff;
                background-color: #0088cc;
                padding: 5px 10px;
                border-radius: 4px;
                text-decoration: none;
            }
            .action btn:hover {
                background-color: #0077b3;
            }
    
            .action a:hover {
                background-color: #0077b3;
            }
    
            .action a.delete {
                background-color: #ff0000;
            }
    
            .action a.delete:hover {
                background-color: #cc0000;
            }
            .action btn.delete {
                background-color: #ff0000;
            }
    
            .action a.download {
                background-color: #0088cc;
            }
    
            .action a.download:hover {
                background-color: #0077b3;
            }
    
            system-info {
                background-color: #0088cc;
                color: rgba(0, 0, 0, 0.5);
                padding: 10px;
                margin-top: 10px;
            }
    
            p {
                text-align: right;
                color: rgba(200, 200, 200, 1);
                margin-right: 10px;
            }
    
            p:hover {
                text-align: right;
                color: #3A3B3C;
            }
            .button-1 {
              border-radius: 8px;
              border-style: none;
              box-sizing: border-box;
              color: #FFFFFF;
              cursor: pointer;
              display: inline-block;
              font-family: "Haas Grot Text R Web", "Helvetica Neue", Helvetica, Arial, sans-serif;
              font-size: 14px;
              font-weight: 500;
              height: 40px;
              line-height: 20px;
              list-style: none;
              margin: 0;
              outline: none;
              padding: 10px 16px;
              position: relative;
              text-align: center;
              text-decoration: none;
              transition: color 100ms;
              vertical-align: baseline;
              user-select: none;
              -webkit-user-select: none;
              touch-action: manipulation;
            }
            .button-1 download {
                background-color: #0088cc;
            }
            .button-1 delete {
                background-color: #ff0000;
            }
            
            .button-1:hover,
            .button-1:focus {
              background-color: #F082AC;
            }
        </style>
    </head>
    
    <body>
        <div class="container">
            <div class="header"
                style="overflow: hidden;background-color: #f1f1f1;padding: 20px 10px;border-bottom: 3px solid #ffcc00;">
                <a href="#default" class="logo">Cloud Torrent</a>
                <div class="header-right">
                    <a href="/" class="active">Torrents</a>
                    <a href="/torrents/search?query=top100">Search</a>
                    <a href="/downloads">Files</a>
                    <a href="/">Home</a>
                </div>
            </div>
            <form action="add" method="post">
                <input type="text" name="magnet" placeholder="Magnet/HTTPS link" class="input">
                <div class="wrapper">
                    <input type="submit" value="Start Torrent">
    
                </div>
            </form>
            <div class="container">
            <div class="content">
            <h2>Torrents</h2>
            </div>
            </div>
            <div style="overflow-x:auto;">
            <table>
                <tr>
                    <th class="id">ID</th>
                    <th class="name">Name</th>
                    <th class="size">Size</th>
                    <th class="status">Status</th>
                    <th class="status">Progress</th>
                    <th class="status">ETA</th>
                    <th class="status">Download Speed</th>
                    <th class="action">Action</th>
                </tr>
                <tr>
                    {{#each torrents}}
                </tr>
            </table>
            </div>
            <div class="system-info">
                <p>
                    <strong>IP:</strong> {{ip}}
                    <strong>CPU:</strong> {{cpu}}
                    <strong>Memory:</strong> {{memory}}
                    <strong>Disk:</strong> {{disk}}
                    <strong>Torrents</strong>: {{torrents_len}}
                    <strong>Goroutines</strong>: {{goroutines}}
                </p>
            </div>
        </div>
        <div class="browse-torrent" style="display: none;" id="torr-b">
            <div class="browse-torrent files">
                <div class="browse-torrent file">
                </div>
            </div>
        </div>
        <script src="https://code.jquery.com/jquery-3.3.1.min.js"></script>
        <script>
            setInterval(function () {
                $.ajax({
                    url: '/torrents',
                    type: 'GET',
                    success: function (data) {
                        var torrents = JSON.parse(data);
                        console.log(torrents);
                        var html = '';
                        html += '<tr>';
                        html += '<th class="id">ID</th>';
                        html += '<th class="name">Name</th>';
                        html += '<th class="size">Size</th>';
                        html += '<th class="status">Status</th>';
                        html += '<th class="status">Progress</th>';
                        html += '<th class="status">ETA</th>';
                        html += '<th class="status">Download Speed</th>';
                        html += '<th class="action">Action</th>';
                        for (var i = 0; i < torrents.length; i++) {
                            var torrent = torrents[i];
                            html += '<tr>';
                            html += '<th class="id">' + torrent.id + '</td>';
                            html += '<th class="name"><a href="/torrents/details?uid=' + torrent.uid + '">' + torrent.name + '</a>' + '</td>';
                            html += '<th class="size">' + torrent.size + '</td>';
                            html += '<th class="status">' + torrent.status + '</td>';
                            html += '<th class="status">' + torrent.perc + '</td>';
                            html += '<th class="status">' + torrent.eta + '</td>';
                            html += '<th class="status">' + torrent.speed + '</td>';
                            html += '<th class="action">';
                            html += '<a href="torrents/details?uid='+torrent.uid+'" class="download">Download</a>';
                            html += '<a href="/" class="delete" onclick="return DeleteBtn(this)" data-uid="'+torrent.uid+'">Delete</a>';
                            html += '</th>';
                            html += '</tr>';
                        }
                        $('table').html(html);
                    } 
                });
            }, 3000);
            function ToggleTorrent(e) {
                var uid = $(e).attr('data-uid');
                $.ajax({
                    url: '/torrents/details?uid=' + uid,
                    type: 'GET',
                    success: function (data) {
                        var torrent = JSON.parse(data);
                        var node = $('torr-b');
                        var visibility = node.style.display;
                        node.style.display = "block";
                    }
                });
            }
            function confirm_action( e ) {
                return !!confirm( "Are you sure to delete torrent?" );
            }
            function DeleteBtn(e) {
                var result = false
                if (confirm_action(e)) {
                    var uid = $(e).attr('data-uid');
                    $.ajax({
                        url: '/torrents/delete?uid=' + uid,
                        type: 'GET',
                        success: function (data) {
                            console.log("delete");
                            result = true;
                        }
                    });
                }
                return result;
            }
        </script>
    </body>
    
    </html>
    `
	torrentsearch = `<!DOCTYPE html>
    <html>
    
    <head>
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
        <title>Search Results for KGF</title>
        <style>
            * {
                box-sizing: border-box;
            }
    
            .menu {
                float: left;
                width: 20%;
                text-align: center;
            }
    
            .menu a {
                background-color: #e5e5e5;
                padding: 8px;
                margin-top: 7px;
                display: block;
                width: 100%;
                color: black;
            }
    
            .main {
                float: left;
                width: 100%;
                padding: 0 20px;
            }
    
            .right {
                background-color: #e5e5e5;
                float: left;
                width: 20%;
                padding: 15px;
                margin-top: 7px;
                text-align: center;
            }
    
            @media only screen and (max-width: 620px) {
    
                /* For mobile phones: */
                .menu,
                .main {
                    width: 100%;
                }
            }
    
            .header {
                overflow: hidden;
                background-color: #f1f1f1;
                padding: 20px 10px;
            }
    
            .header a {
                float: left;
                color: black;
                text-align: center;
                padding: 12px;
                text-decoration: none;
                font-size: 18px;
                line-height: 25px;
                border-radius: 4px;
            }
    
            .header a.logo {
                font-size: 25px;
                font-weight: bold;
            }
    
            .header a:hover {
                background-color: #ddd;
                color: black;
            }
    
            .header a.active {
                background-color: dodgerblue;
                color: white;
            }
    
            .header-right {
                float: right;
            }
    
            .topnav {
                overflow: hidden;
                background-color: #e9e9e9;
                position: sticky;
                position: -webkit-sticky;
            }
    
            .topnav p {
                float: left;
                display: block;
                color: black;
                text-align: center;
                padding: 14px 16px;
                text-decoration: none;
                font-size: 17px;
            }
    
            .topnav p:hover {
                background-color: #ddd;
                color: black;
            }
    
            .topnav p.active {
                background-color: #2196F3;
                color: white;
            }
    
            .topnav .search-container {
                float: right;
            }
    
            .topnav input[type=text] {
                padding: 6px;
                margin-top: 8px;
                font-size: 17px;
                border: none;
            }
    
            .topnav .search-container button {
                float: right;
                padding: 6px 10px;
                margin-top: 8px;
                margin-right: 16px;
                background: #ddd;
                font-size: 17px;
                border: none;
                cursor: pointer;
            }
    
            .topnav .search-container button:hover {
                background: #ccc;
            }
    
            .btn {
                background-color: DodgerBlue;
                border: none;
                color: white;
                padding: 5px 13px;
                cursor: pointer;
                font-size: 16px;
                border-radius: 8px;
            }
    
            .btn:hover {
                background-color: RoyalBlue;
            }
    
            .container {
                max-width: 1000px;
                margin-left: auto;
                margin-right: auto;
                padding-left: 10px;
                padding-right: 10px;
                font-family: sans-serif;
            }
    
            h2 {
                font-size: 26px;
                margin: 20px 0;
                text-align: center;
            }
    
            h2 small {
                font-size: 0.5em;
            }
    
            .responsive-table li {
                border-radius: 3px;
                padding: 25px 30px;
                display: flex;
                justify-content: space-between;
                margin-bottom: 25px;
                color: #050505;
            }
    
            .responsive-table .table-header {
                background-color: #95A5A6;
                font-size: 14px;
                text-transform: uppercase;
                letter-spacing: 0.03em;
                color: #050505;
            }
    
            .responsive-table .table-row {
                background-color: #ffffff;
                box-shadow: 0px 0px 9px 0px rgba(0, 0, 0, 0.1);
            }
    
            .responsive-table .table-row:hover {
                background-color: #f2f2f2;
            }
    
            .responsive-table .col-1 {
                flex-basis: 10%;
            }
    
            .responsive-table .col-2 {
                flex-basis: 40%;
            }
    
            .responsive-table .col-3 {
                flex-basis: 25%;
            }
    
            .responsive-table .col-4 {
                flex-basis: 25%;
            }
    
            .responsive-table .col-5 {
                flex-basis: 25%;
            }
    
            .responsive-table .col-6 {
                flex-basis: 25%;
            }
    
            @media all and (max-width: 767px) {
                .responsive-table .table-header {
                    display: none;
                }
    
                .responsive-table li {
                    display: block;
                }
    
                .responsive-table .col {
                    flex-basis: 100%;
                }
    
                .responsive-table .col {
                    display: flex;
                    padding: 10px 0;
                }
    
                .responsive-table .col:before {
                    color: #084281;
                    padding-right: 10px;
                    content: attr(data-label);
                    flex-basis: 50%;
                    text-align: right;
                }
            }
    
            .dark-mode {
                background: #333;
                color: #fff;
            }
        </style>
    </head>
    
    <body style="font-family:Verdana;color:#aaaaaa;">
    
        <div class="header"
            style="overflow: hidden;background-color: #f1f1f1;padding: 20px 10px;border-bottom: 3px solid #ffcc00;">
            <a href="#default" class="logo">Cloud Torrent</a>
            <div class="header-right">
                <a href="/">Torrents</a>
                <a href="/torrents/search" class="active">Search</a>
                <a href="/downloads">Files</a>
                <a href="/">Home</a>
            </div>
        </div>
        <div style="overflow:auto">
            <div class="topnav">
                <p>Search Results for <b>{{query}}</b></p>
                <div class="search-container">
                    <form action="/torrents/search">
                        <input type="text" placeholder="Search.." name="query">
                        <button type="submit"><i class="fa fa-search"></i></button>
                    </form>
                </div>
            </div>
    
            <div class="container">
                <ul class="responsive-table">
                    <li class="table-header">
                        <div class="col col-1">ID</div>
                        <div class="col col-2">Name</div>
                        <div class="col col-3">Size</div>
                        <div class="col col-4">Seeders</div>
                        <div class="col col-5">Seeders</div>
                        <div class="col col-6">Seeders</div>
                    </li>
                    {{#torrents}}
                </ul>
            </div>
        </div>
        </div>
        <div style="background-color:#e5e5e5;text-align:center;padding:10px;margin-top:7px;">© copyright roseloverx 2022
        </div>
        <button onclick="Dark()">Toggle dark mode</button>
        <script src="https://code.jquery.com/jquery-3.3.1.min.js"></script>
        <script type="text/javascript">
            function AddedTorr(e) {
                var magnet = e.getAttribute('data-magnet');
                $.ajax({
                    url: '/torrents/add',
                    type: 'GET',
                    data: {
                        magnet: magnet
                    },
                    success: function (data) {
                        console.log("Torrent added");
                        alert("Torrent added\n");
                    }
                });
            }
            function Dark() {
                var element = document.body;
                element.classList.toggle("dark-mode");
            }
        </script>
    </body>
    
    </html>
`
) // templates edit
