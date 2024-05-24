// const API_URL = 'http://192.168.9.239:8000'
const API_URL = 'http://localhost:8000'

export function connectSocket(playerID) {
    const url = API_URL.replaceAll('http://', 'ws://') + '/ws/findmatch?player_id=' + playerID
    if (window['WebSocket']) {
        let gameID;
        const socket = new WebSocket(url);
        socket.addEventListener('open', (event) => {
            socket.send('Hello Server!');
        });
        socket.addEventListener('close', (event) => {
            window.location = '/game#' + gameID
        });
        socket.addEventListener('message', (event) => {
            console.log('WS: ' +event.data)
            socket.close()
            gameID = event.data
        });
        
    } else {
        alert('Your browser does not support WebSockets')
    }
}

