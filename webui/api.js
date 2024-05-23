const API_URL = 'http://localhost:8000'

const apiOpts = {
    headers: {
        'Content-Type': 'application/json',
    },
}

async function createPlayer(playerID) {
    const res = await fetch(API_URL + '/players', {
        method: 'POST',
        ...apiOpts,
        body: JSON.stringify({
            id: playerID
        }),
    });
    const player = await res.json();
    return player
}

async function getPlayer(playerID) {
    const res = await fetch(API_URL + '/players/' + playerID, {
        method: 'GET',
        ...apiOpts,
    });
    const player = await res.json();
    return player
}

export default {
    createPlayer,
    getPlayer,
}