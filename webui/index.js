import idgen from './idgen.js'
import api from './api.js'

const LOCAL_PLAEYR_ID_KEY = 'rps_player_id'

function main() {
    let playerID = window.localStorage.getItem(LOCAL_PLAEYR_ID_KEY)
    if (!playerID) {
        alert('Hi Welcome! we will be generating name for you!')

        playerID = idgen()
        
        api.createPlayer(playerID).then((data) => {
            console.log(data)
        }).catch(e => {
            alert('Ooops! curse of demos :)')
            console.error(e.message)
            return
        })

        window.localStorage.setItem(LOCAL_PLAEYR_ID_KEY, playerID)
    }

    document.getElementById('player_text_id').innerText = playerID
}

main()