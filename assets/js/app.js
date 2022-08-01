const currentTime = new Date();
const completeSVG = `<svg class="flex-shrink-0 w-5 h-5 text-blue-600 dark:text-blue-500" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"></path></svg>`;
const foxSVG = `<svg class="h-4" viewBox="0 0 40 38" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M39.0728 0L21.9092 12.6999L25.1009 5.21543L39.0728 0Z" fill="#E17726"/><path d="M0.966797 0.0151367L14.9013 5.21656L17.932 12.7992L0.966797 0.0151367Z" fill="#E27625"/><path d="M32.1656 27.0093L39.7516 27.1537L37.1004 36.1603L27.8438 33.6116L32.1656 27.0093Z" fill="#E27625"/><path d="M7.83409 27.0093L12.1399 33.6116L2.89876 36.1604L0.263672 27.1537L7.83409 27.0093Z" fill="#E27625"/><path d="M17.5203 10.8677L17.8304 20.8807L8.55371 20.4587L11.1924 16.4778L11.2258 16.4394L17.5203 10.8677Z" fill="#E27625"/><path d="M22.3831 10.7559L28.7737 16.4397L28.8067 16.4778L31.4455 20.4586L22.1709 20.8806L22.3831 10.7559Z" fill="#E27625"/><path d="M12.4115 27.0381L17.4768 30.9848L11.5928 33.8257L12.4115 27.0381Z" fill="#E27625"/><path d="M27.5893 27.0376L28.391 33.8258L22.5234 30.9847L27.5893 27.0376Z" fill="#E27625"/><path d="M22.6523 30.6128L28.6066 33.4959L23.0679 36.1282L23.1255 34.3884L22.6523 30.6128Z" fill="#D5BFB2"/><path d="M17.3458 30.6143L16.8913 34.3601L16.9286 36.1263L11.377 33.4961L17.3458 30.6143Z" fill="#D5BFB2"/><path d="M15.6263 22.1875L17.1822 25.4575L11.8848 23.9057L15.6263 22.1875Z" fill="#233447"/><path d="M24.3739 22.1875L28.133 23.9053L22.8184 25.4567L24.3739 22.1875Z" fill="#233447"/><path d="M12.8169 27.0049L11.9606 34.0423L7.37109 27.1587L12.8169 27.0049Z" fill="#CC6228"/><path d="M27.1836 27.0049L32.6296 27.1587L28.0228 34.0425L27.1836 27.0049Z" fill="#CC6228"/><path d="M31.5799 20.0605L27.6165 24.0998L24.5608 22.7034L23.0978 25.779L22.1387 20.4901L31.5799 20.0605Z" fill="#CC6228"/><path d="M8.41797 20.0605L17.8608 20.4902L16.9017 25.779L15.4384 22.7038L12.3988 24.0999L8.41797 20.0605Z" fill="#CC6228"/><path d="M8.15039 19.2314L12.6345 23.7816L12.7899 28.2736L8.15039 19.2314Z" fill="#E27525"/><path d="M31.8538 19.2236L27.2061 28.2819L27.381 23.7819L31.8538 19.2236Z" fill="#E27525"/><path d="M17.6412 19.5088L17.8217 20.6447L18.2676 23.4745L17.9809 32.166L16.6254 25.1841L16.625 25.1119L17.6412 19.5088Z" fill="#E27525"/><path d="M22.3562 19.4932L23.3751 25.1119L23.3747 25.1841L22.0158 32.1835L21.962 30.4328L21.75 23.4231L22.3562 19.4932Z" fill="#E27525"/><path d="M27.7797 23.6011L27.628 27.5039L22.8977 31.1894L21.9414 30.5138L23.0133 24.9926L27.7797 23.6011Z" fill="#F5841F"/><path d="M12.2373 23.6011L16.9873 24.9926L18.0591 30.5137L17.1029 31.1893L12.3723 27.5035L12.2373 23.6011Z" fill="#F5841F"/><path d="M10.4717 32.6338L16.5236 35.5013L16.4979 34.2768L17.0043 33.8323H22.994L23.5187 34.2753L23.48 35.4989L29.4935 32.641L26.5673 35.0591L23.0289 37.4894H16.9558L13.4197 35.0492L10.4717 32.6338Z" fill="#C0AC9D"/><path d="M22.2191 30.231L23.0748 30.8354L23.5763 34.8361L22.8506 34.2234H17.1513L16.4395 34.8485L16.9244 30.8357L17.7804 30.231H22.2191Z" fill="#161616"/><path d="M37.9395 0.351562L39.9998 6.53242L38.7131 12.7819L39.6293 13.4887L38.3895 14.4346L39.3213 15.1542L38.0875 16.2779L38.8449 16.8264L36.8347 19.1742L28.5894 16.7735L28.5179 16.7352L22.5762 11.723L37.9395 0.351562Z" fill="#763E1A"/><path d="M2.06031 0.351562L17.4237 11.723L11.4819 16.7352L11.4105 16.7735L3.16512 19.1742L1.15488 16.8264L1.91176 16.2783L0.678517 15.1542L1.60852 14.4354L0.350209 13.4868L1.30098 12.7795L0 6.53265L2.06031 0.351562Z" fill="#763E1A"/><path d="M28.1861 16.2485L36.9226 18.7921L39.7609 27.5398L32.2728 27.5398L27.1133 27.6049L30.8655 20.2912L28.1861 16.2485Z" fill="#F5841F"/><path d="M11.8139 16.2485L9.13399 20.2912L12.8867 27.6049L7.72971 27.5398H0.254883L3.07728 18.7922L11.8139 16.2485Z" fill="#F5841F"/><path d="M25.5283 5.17383L23.0847 11.7736L22.5661 20.6894L22.3677 23.4839L22.352 30.6225H17.6471L17.6318 23.4973L17.4327 20.6869L16.9139 11.7736L14.4707 5.17383H25.5283Z" fill="#F5841F"/></svg>`
const resumeSVG = `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="h-4 bi bi-collection-play" viewBox="0 0 16 16"><path d="M2 3a.5.5 0 0 0 .5.5h11a.5.5 0 0 0 0-1h-11A.5.5 0 0 0 2 3zm2-2a.5.5 0 0 0 .5.5h7a.5.5 0 0 0 0-1h-7A.5.5 0 0 0 4 1zm2.765 5.576A.5.5 0 0 0 6 7v5a.5.5 0 0 0 .765.424l4-2.5a.5.5 0 0 0 0-.848l-4-2.5z"/><path d="M1.5 14.5A1.5 1.5 0 0 1 0 13V6a1.5 1.5 0 0 1 1.5-1.5h13A1.5 1.5 0 0 1 16 6v7a1.5 1.5 0 0 1-1.5 1.5h-13zm13-1a.5.5 0 0 0 .5-.5V6a.5.5 0 0 0-.5-.5h-13A.5.5 0 0 0 1 6v7a.5.5 0 0 0 .5.5h13z"/></svg>`
const pauseSVG = `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="h-4 bi bi-pause-btn" viewBox="0 0 16 16"><path d="M6.25 5C5.56 5 5 5.56 5 6.25v3.5a1.25 1.25 0 1 0 2.5 0v-3.5C7.5 5.56 6.94 5 6.25 5zm3.5 0c-.69 0-1.25.56-1.25 1.25v3.5a1.25 1.25 0 1 0 2.5 0v-3.5C11 5.56 10.44 5 9.75 5z"/><path d="M0 4a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V4zm15 0a1 1 0 0 0-1-1H2a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V4z"/></svg>`
const sleepSVG = `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="red" class="h-4" viewBox="0 0 16 16"><path d="m7.949 7.998.006-.003.003.009-.01-.006Zm.025-.028v-.03l.018.01-.018.02Zm0 .015.04-.022.01.006v.04l-.029.016-.021-.012v-.028Zm.049.057v-.014l-.008.01.008.004Zm-.05-.008h.006l-.006.004v-.004Z"/><path fill-rule="evenodd" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0ZM4.965 1.69a6.972 6.972 0 0 1 3.861-.642c.722.767 1.177 1.887 1.177 3.135 0 1.656-.802 3.088-1.965 3.766 1.263.24 2.655-.815 3.406-2.742.38-.975.537-2.023.492-2.996a7.027 7.027 0 0 1 2.488 3.003c-.303 1.01-1.046 1.966-2.128 2.59-1.44.832-3.09.85-4.26.173l.008.021.012-.006-.01.01c.42 1.218 2.032 1.9 4.08 1.586a7.415 7.415 0 0 0 2.856-1.081 6.963 6.963 0 0 1-1.358 3.662c-1.03.248-2.235.084-3.322-.544-1.433-.827-2.272-2.236-2.279-3.58l-.012-.003c-.845.972-.63 2.71.666 4.327a7.415 7.415 0 0 0 2.37 1.935 6.972 6.972 0 0 1-3.86.65c-.727-.767-1.186-1.892-1.186-3.146 0-1.658.804-3.091 1.969-3.768l-.002-.007c-1.266-.25-2.666.805-3.42 2.74a7.415 7.415 0 0 0-.49 3.012 7.026 7.026 0 0 1-2.49-3.018C1.87 9.757 2.613 8.8 3.696 8.174c1.438-.83 3.084-.85 4.253-.176l.005-.006C7.538 6.77 5.924 6.085 3.872 6.4c-1.04.16-2.03.55-2.853 1.08a6.962 6.962 0 0 1 1.372-3.667l-.002.003c1.025-.243 2.224-.078 3.306.547 1.43.826 2.269 2.23 2.28 3.573L8 7.941c.837-.974.62-2.706-.673-4.319a7.415 7.415 0 0 0-2.362-1.931Z"/></svg>`


function getTorrents() {
    $.ajax({
        url: '/api/torrents',
        type: 'GET',
        dataType: 'json',
        success: function (data) {
            html = '';
            if (data == null) {
                return $('#torrents').html('<div class="w-full flex justify-center items-center">No torrents found</div>');
            }
            data.forEach(t => {
                var toggleButton = `<button id='toggle' type="button" class="inline-block px-4 py-2.5 bg-yellow-300 text-white font-medium text-xs leading-tight uppercase shadow-md hover:bg-yellow-400 hover:shadow-lg focus:bg-yellow-400 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-yellow-500 active:shadow-lg transition duration-150 ease-in-out" onclick='toggleTorrent("${t.uid}")'>${pauseSVG}</button>`
                if (IsPaused(t)) {
                    toggleButton = `<button id='toggle' type="button" class="inline-block px-4 py-2.5 bg-green-600 text-white font-medium text-xs leading-tight uppercase shadow-md hover:bg-green-700 hover:shadow-lg focus:bg-green-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-green-800 active:shadow-lg transition duration-150 ease-in-out" onclick='toggleTorrent("${t.uid}", true)'>${resumeSVG}</button>`
                }
                var svg = foxSVG;
                if (t.progress == '100.00') {
                    svg = completeSVG
                }
                html += `<div class="card w-full sm:w-1/2 md:w-1/4 xl:w-1/5 p-1" id="card|${t.uid}">
                <div class="block rounded-lg shadow-lg bg-white max-w-sm text-center border-gray-300 ">
                  <div class="py-3 px-6 border-b border-gray-300">
                    <a style='margin-top: 10px;' href="#"
                      class="flex items-center p-3 text-base font-bold text-gray-900 bg-gray-50 rounded-lg hover:bg-gray-100 group hover:shadow dark:bg-gray-600 dark:hover:bg-gray-500 dark:text-white">
                      <span id='svg'>${svg}</span>
                      <span class="flex-1 ml-3 text-xs" id='speed'>${t.speed}</span>
                      <span
                        class="inline-flex items-center justify-center px-2.5 py-0.5 ml-3 text-xs font-medium text-gray-500 bg-gray-200 rounded dark:bg-gray-700 dark:text-gray-400"
                        id='eta'>${t.eta}</span>
                    </a>
                  </div>
                  <div class="p-6 h-28">
                    <h5 class="text-gray-900 text-sm font-medium mb-2 break-words"><span id='title'>${TrimString(t.name, 40)}</span>
                      <span
                        class="text-xs font-semibold inline-block py-1 px-2 uppercase rounded text-pink-600 bg-pink-200 uppercase last:mr-0 mr-1"
                        id='size'>
                        ${t.size}
                      </span></h5>
                  </div>
                  <div class="py-3 px-2 border-t border-gray-300 text-gray-600">
                    <div class='flex flex-wrap justify-center'>
                      <p class="text-gray-700 text-base mb-4">
                      <div class="w-full bg-gray-200 rounded-full h-2.5 dark:bg-gray-700" id='progress'>
                        <div class="${getBarColor(t)} h-2.5 rounded-full" style="width: ${t.perc}"></div>
                      </div>
                      </p>
                    </div>
                  </div>
                  <div class="py-3 px-2 border-t border-gray-300 text-gray-600">
                    <div class='flex flex-wrap justify-center'>
                      <a type="button" href='${t.path}'
                        class=" inline-block px-4 py-2.5 bg-blue-600 text-white font-medium text-xs leading-tight uppercase rounded-l shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out">Browse</a>
                      ${toggleButton}
                      <button type="button"
                        class=" inline-block px-4 py-2.5 bg-red-600 text-white font-medium text-xs leading-tight uppercase rounded-r shadow-md hover:bg-red-700 hover:shadow-lg focus:bg-red-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-red-800 active:shadow-lg transition duration-150 ease-in-out"
                        onclick='ConfirmDelete("${t.uid}")'>Delete</button>
                    </div>
                  </div>
                </div>
              </div>`
            });
            $('#torrents').html(html);
        }
    });
}

getTorrents();

function UpdateEvent(data) {
    var data = JSON.parse(data);
    var card_uids = [];
    document.querySelectorAll('#torrents .card').forEach(function (card) {
        card_uids.push(card.id.split('|')[1]);
    })
    var torrent_uids = [];
    if (data !== null) {
        data.forEach(t => {
            var card = document.getElementById('card|' + t.uid);
            torrent_uids.push(t.uid);
            if (card) {
                card.querySelector('#speed').innerHTML = t.speed;
                card.querySelector('#eta').innerHTML = t.eta;
                card.querySelector('#progress').innerHTML = `<div class="${getBarColor(t)} h-2.5 rounded-full" style="width: ${t.perc}"></div>`;
                card.querySelector('#size').innerHTML = t.size;
                card.querySelector('#title').innerHTML = TrimString(t.name, 40);
                var svg = foxSVG && t.progress === '100.00' ? completeSVG : foxSVG;
                card.querySelector('#svg').innerHTML = svg;
            }
        });
    } else {
        $('#torrents').html('<div class="w-full flex justify-center items-center">No torrents found</div>');
    }
    card_uids.forEach(uid => {
        if (!torrent_uids.includes(uid)) {
            console.log(torrent_uids, card_uids, uid);
            var pcard = document.getElementById('card|' + uid)
            if (pcard) {
                pcard.remove();
            }
        }
    })
}

function getBarColor(torrent) {
    if (torrent.progress == '100.00') {
        return 'bg-green-500';
    } else if (torrent.progress == '0') {
        return 'bg-red-500';
    } else if (torrent.status == 'Stopped' && torrent.progress !== '100.00') {
        return 'bg-yellow-300';
    } else {
        return 'bg-blue-500';
    }
}

function IsPaused(torrent) {
    return torrent.status == 'Stopped' && torrent.progress !== '100.00';
}

const updater = new EventSource("/torrents/update");
updater.addEventListener("torrents", (e) => { UpdateEvent(e.data); }, false);

function AddTorrent() {
    $.ajax({
        url: '/api/add',
        type: 'POST',
        dataType: 'json',
        data: {
            magnet: $('#input-magnet').val(),
        },
        error: function (xhr, status, error) {
            tata.error("Failed Adding torrent", xhr.responseText);
            return;
        }
    });
    tata.success('Success!', 'Torrent added');
    $('#input-magnet').val('');
    getTorrents();
}


function DeleteTorrent(uid) {
    $.ajax({
        url: '/api/remove',
        type: 'POST',
        dataType: 'json',
        data: {
            uid: uid,
        },
        success: function (data) {
            console.log(data);
        },
        error: function (data) {
            console.log(data);
        }
    });
}

function ConfirmDelete(uid) {
    swal({
        title: "Are you sure?",
        text: "You will not be able to recover this torrent!",
        icon: "warning",
        buttons: ["Cancel", "Delete"],
    }).then((willDelete) => {
        if (willDelete) {
            $.ajax({
                url: '/api/torrent',
                type: 'GET',
                dataType: 'json',
                data: {
                    uid: uid,
                },
                success: function (data) {
                    tata.error('Deleted!', 'Deleted ' + TrimString(data.name, 20));
                    DeleteTorrent(uid);
                }
            });
        }
    });
}

function toggleTorrent(uid, unpause) {
    var url = '/api/pause';
    if (unpause) {
        url = '/api/resume';
    }
    $.ajax({
        url: url,
        type: 'GET',
        data: {
            uid: uid,
        },
        success: function (data) {
            $.ajax({
                url: '/api/torrent',
                type: 'GET',
                dataType: 'json',
                data: {
                    uid: uid,
                },
                success: function (data) {
                    card = document.getElementById('card|' + uid);
                    if (unpause) {
                        card.querySelector('#progress').innerHTML = `<div class="bg-blue-400 h-2.5 rounded-full" style="width: ${data.perc}"></div>`;
                        card.querySelector('#toggle').outerHTML = `<button id='toggle' type="button" class="inline-block px-4 py-2.5 bg-yellow-300 text-white font-medium text-xs leading-tight uppercase shadow-md hover:bg-yellow-400 hover:shadow-lg focus:bg-yellow-400 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-yellow-500 active:shadow-lg transition duration-150 ease-in-out" onclick='toggleTorrent("${uid}")'>${pauseSVG}</button>`;
                        card.querySelector('#svg').innerHTML = foxSVG;
                        tata.success("Done!", 'Resumed ' + TrimString(data.name, 20));
                    } else {
                        card.querySelector('#progress').innerHTML = `<div class="bg-yellow-300 h-2.5 rounded-full" style="width: ${data.perc}"></div>`;
                        card.querySelector('#toggle').outerHTML = `<button id='toggle' type="button" class="inline-block px-4 py-2.5 bg-green-600 text-white font-medium text-xs leading-tight uppercase shadow-md hover:bg-green-700 hover:shadow-lg focus:bg-green-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-green-800 active:shadow-lg transition duration-150 ease-in-out" onclick='toggleTorrent("${uid}", true)'>${resumeSVG}</button>`
                        card.querySelector('#svg').innerHTML = sleepSVG;
                        card.querySelector('#speed').innerHTML = '-/-';
                        card.querySelector('#eta').innerHTML = ``;
                        tata.warn("Done!", 'Paused ' + TrimString(data.name, 20));
                    }
                }
            });
        }
    });
}



document.getElementById('submit-magnet').addEventListener('click', AddTorrent);

function TrimString(str, s) {
    return str.substring(0, s) + '...';
}

// https://itunes.apple.com/search?term=Avengers&entity=movie