package main

const (
	downloads = `
    <!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8">
    <title>Downloads Index</title>
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

        a {
            color: #0088cc;
            text-decoration: none;
        }

        a:hover {
            color: #005580;
            text-decoration: underline;
        }

        .container {
            width: 960px;
            margin: 0 auto;
            box-shadow: 0 5px 10px -5px rgba(0, 0, 0, 0.5);
        }

        .header {
            background-color: #0088cc;
            color: #fff;
            padding: 10px 0;
        }

        .header h1 {
            margin: 0;
            font-size: 24px;
            font-weight: normal;
            line-height: 1.5;
        }

        .content {
            padding: 10px 0;
        }

        .content h2 {
            margin: auto;
            font-size: 18px;
            font-weight: normal;
            line-height: 1.5;
        }

        table {
            border-collapse: collapse;
            width: 100%;
        }

        table td,
        table th {
            border: 1px solid #ddd;
            padding: 8px;
        }

        table tr:nth-child(even) {
            background-color: #f2f2f2;
        }

        table tr:hover {
            background-color: #ddd;
        }

        table th {
            padding-top: 12px;
            padding-bottom: 12px;
            text-align: left;
            background-color: #fd5e53;
            color: white;
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

        button:hover {
            opacity: 0.8;
        }
    </style>
</head>

<body>
    <div class="container">
        <div class="header">
            <h1>Downloads</h1>
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
        <button onclick="window.location.href='/downloads/'">Go Back</button>
        <button onclick="window.location.href='/'">Add Magnet</button>
    </div>
    <footer>
        <div class="footer-content">
            <p>&copy; 2021-22</p>
            <p><a href="t.me/roseloverx">RoseLoverX</a><br>IP: {{ip}}</p>
        </div>
    </footer>
</body>

</html>
`
	torrents = `
    <!DOCTYPE html>
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
    
            header {
                background-color: #0088cc;
                color: #fff;
                padding: 10px 0;
            }
    
            header h1 {
                margin: 0;
                font-size: 24px;
                font-weight: normal;
                line-height: 1.5;
            }
    
            header h2 {
                margin: 0;
                font-size: 18px;
                font-weight: normal;
                line-height: 1.5;
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
    
            table tr:nth-child(even){background-color: #f2f2f2;}
    
            table tr:hover {background-color: #ddd;}
    
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
            table th.name a:visited {
                color: #333;
            }
            table th.name a:active {
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
            table td.container {
                width: 100%;
                text-align: center;
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
    
            .popup {
                position: fixed;
                top: 0;
                left: 0;
                width: 100%;
                height: 100%;
                background-color: rgba(0, 0, 0, 0.5);
                z-index: 99999;
                display: none;
            }
    
            .popup .popup-content {
                position: absolute;
                top: 50%;
                left: 50%;
                transform: translate(-50%, -50%);
                background-color: #fff;
    
                width: 500px;
                height: 300px;
                border-radius: 4px;
                padding: 20px;
            }
    
            .popup .popup-text {
                font-size: 20px;
                font-weight: bold;
                margin-bottom: 20px;
            }
    
            .popup .popup-text a {
                color: #0088cc;
                text-decoration: none;
            }
    
            .popup .popup-text a:hover {
                text-decoration: underline;
            }
    
            .popup .popup-button {
                margin-top: 20px;
            }
    
            .popup .popup-button button {
                width: 100%;
                height: 35px;
                background-color: #0088cc;
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
            <header>
                <h1>Torrents Cloud
                </h1>
            </header>
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
                    <th class="magnet">Magnet</th>
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
        <script src="https://code.jquery.com/jquery-3.3.1.min.js"></script>
        <script>
            setInterval(function() {
            $.ajax({
                url: '/torrents',
                type: 'GET',
                success: function(data) {
                    var torrents = JSON.parse(data);
                    console.log(torrents);
                    var html = '';
                    html += '<tr>';
                    html += '<th class="id">ID</th>';
                    html += '<th class="name">Name</th>';
                    html += '<th class="size">Size</th>';
                    html += '<th class="status">Status</th>';
                    html += '<th class="magnet">Magnet</th>';
                    html += '<th class="action">Action</th>';
                    for (var i = 0; i < torrents.length; i++) {
                        var torrent = torrents[i];
                        html += '<tr>';
                        html += '<th class="id">' + torrent.id + '</td>';
                        html += '<th class="name"><a href="/torrents/details?uid=' + torrent.uid + '>torrent.name</a>' + '</td>';
                        html += '<th class="status">' + torrent.status + '</td>';
                        html += '<th class="date">' + torrent.date + '</td>';
                        html += '<th class="magnet">' + torrent.magnet + '</td>';
                        html += '<th class="action">';
                        html += '<a href="/torrents' + torrent.id + '/download" class="download">Download</a>';
                        html += '<a href="/torrents/delete?uid=' + torrent.uid + '" class="delete">Delete</a>';
                        html += '</th>';
                        html += '</tr>';
                    }
                    $('table').html(html);
                }
            });
        }, 50000);
    
        </script>
    </body>
    </html>
    `
)
