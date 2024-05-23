const LOCAL_PLAYER_ID_KEY = 'rps_player_id'

export default {
    getPlayer: () => window.localStorage.getItem(LOCAL_PLAYER_ID_KEY),
    setPlayer: (playerID) => window.localStorage.setItem(LOCAL_PLAYER_ID_KEY, playerID),
}