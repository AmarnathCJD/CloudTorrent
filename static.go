package main

const (
	downloads = `<!DOCTYPE html>
    <html>
    
    <head>
        <meta charset="utf-8">
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
    
            table tr:hover {
                background-color: #ddd;
            }
    
            table th.id {
                width: 5%;
            }
    
            table th.name {
                width: 20%;
            }
    
            table th.name a {
                color: #333;
            }
    
            table th.name a:hover {
                color: #0088cc;
            }
    
            table th.size {
                width: 10%;
            }
    
            table th.date {
                width: 15%;
            }
    
            table th.magnet {
                width: 20%;
            }
    
            table th.action {
                width: 10%;
            }
    
            .action {
                width: 10%;
            }
    
            .action a {
                color: #fff;
                background-color: #0088cc;
                padding: 5px 10px;
                border-radius: 4px;
                text-decoration: none;
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
    
            browse-torrent {
                background-color: #0088cc;
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
    
            browse-torrent files {
                background-color: #0088cc;
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
    
            browse-torrent file {
                width: 100%;
                height: 40px;
                border: 1px solid #ccc;
                border-radius: 4px;
                padding: 0 10px;
                font-size: 14px;
                line-height: 1.5;
                margin-top: 10px;
            }
    
            browse-torrent file:focus {
                outline: none;
            }
    
            browse-torrent file download {
                background-color: #3c5abd;
                color: white;
                padding: 7px 2px;
                border: none;
                border-radius: 4px;
                cursor: pointer;
                margin-bottom: 10px;
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
            </div class="content">
            <h2>Torrents</h2>
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
                            html += '<a href="/torrents/details?uid=' + torrent.uid + '" class="download">Download</a>';
                            html += '<a href="/torrents/delete?uid=' + torrent.uid + '" class="delete">Delete</a>';
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
        </script>
    </body>
    
    </html>
    `
	torrentsearch = `<!DOCTYPE html>
    <html lang="en">
    
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta http-equiv="X-UA-Compatible" content="ie=edge">
        <title>Search Results for {{query}}</title>
        <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css"
            integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
        <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
        <style>
            .container {
                width: auto;
                margin: 0 auto;
                box-shadow: 0 5px 10px -5px rgba(0, 0, 0, 0.5);
            }
    
            * {
                box-sizing: border-box;
            }
    
            body {
                margin: 0;
                font-family: Arial, Helvetica, sans-serif;
                font-size: 14px;
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
    
            result {
                margin-top: 50px;
            }
    
            table {
                margin-top: 50px;
                border-collapse: collapse;
                width: 100%;
                background-color: #f2f2f2;
                overflow-x: auto;
            }
    
            th,
            td {
                text-align: left;
                padding: 8px;
            }
    
            tr:first-child th {
                background-color: #d6d6d6;
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
            }
    
            .btn:hover {
                background-color: RoyalBlue;
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
                    <a href="/torrents/search" class="active">Search</a>
                    <a href="/downloads">Files</a>
                    <a href="/">Home</a>
                </div>
            </div>
            <div class="topnav">
                <p>Search Results for <b>{{query}}</b></p>
                <div class="search-container">
                    <form action="/torrents/search">
                        <input type="text" placeholder="Search.." name="query">
                        <button type="submit"><i class="fa fa-search"></i></button>
                    </form>
                </div>
            </div>
            <div class="result">
                <table>
                    <tr>
                        <th>ID</th>
                        <th>Name</th>
                        <th>Size</th>
                        <th>Seeds</th>
                        <th>Leechers</th>
                        <th>Action</th>
                    </tr>
                    {{results}}
                </table>
            </div>
        </div>
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
        </script>
    </body>
    
    </html>
`
)
