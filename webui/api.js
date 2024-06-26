// const API_URL = 'http://192.168.9.239:8000'
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
    return await res.json();
}

async function getPlayer(playerID) {
    const res = await fetch(API_URL + '/players/' + playerID, {
        method: 'GET',
        ...apiOpts,
    });
    return await res.json();
}

async function getGame(gameID) {
    const res = await fetch(API_URL + '/games/' + gameID, {
        method: 'GET',
        ...apiOpts,
    });
    return await res.json();
}

async function castGame(gameID, playerID, cast) {
    const res = await fetch(API_URL + '/castgames/' + gameID, {
        method: 'POST',
        ...apiOpts,
        body: JSON.stringify({
            player_id: playerID,
            throw: cast,
        }),
    });
    return await res.json();
}

export default {
    API_URL,
    createPlayer,
    getPlayer,
    getGame,
    castGame,
}