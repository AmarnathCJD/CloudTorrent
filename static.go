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
            align-self: right;
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
            font-size: 18px;
            font-weight: normal;
            line-height: 1.5;
        }
        input {
            width: 100%;
            height: 40px;
            border: 1px solid #ccc;
            border-radius: 4px;
            padding: 0 10px;
            font-size: 14px;
            line-height: 1.5;
            margin-top: 10px;
        }
        input:focus {
            outline: none;
        }
        table {
            border-collapse: collapse;
            width: 100%;
        }
        submit {
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
        submit:hover {
            opacity: 0.8;
        }
        </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>Torrents Cloud</h1>
            <h2>{{ip}}</h2>
            <form action="/torrents/add" method="post">
                <input type="text" name="magnet" placeholder="Magnet Link" class="magnet-input">
                <input type="submit" value="Add">
            </form>
        </div>
        <div class="content">
            <h2>Torrents</h2>
            <table>
                <tr>
                    <th>Name</th>
                    <th>Size</th>
                    <th>Seeders</th>
                    <th>Leechers</th>
                    <th>Status</th>
                </tr>
                {{torrents}}
            </table>
        </div>
    </div>
</body>
</html>
`
)
