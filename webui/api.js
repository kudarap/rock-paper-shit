const API_URL = 'http://localhost:8000'

async function createPlayer(playerID) {
    const res = await fetch(API_URL + '/version', {
        method: 'GET',
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