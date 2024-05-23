
const API_URL = 'http://localhost:3000'

async function createPlayer(playerID) {
    const res = await fetch(API_URL + '/mockapi/postplayer.json', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            id: playerID
        }),
    });
    const player = await res.json();
    return player
}


export default {
    createPlayer
}