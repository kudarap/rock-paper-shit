import idgen from './idgen.js'
import api from './api.js'

const LOCAL_PLAEYR_ID_KEY = 'rps_player_id'

function main() {
    let playerID = window.localStorage.getItem(LOCAL_PLAEYR_ID_KEY)
    if (!playerID) {
        alert('Hi Welcome! we will be generating name for you!')

        playerID = idgen()
        api.createPlayer(playerID).then((player) => {
            console.log(player)
            document.getElementById('player_text_id').innerText = player.id
            document.getElementById('player_text_ranking').innerText = player.ranking
            window.localStorage.setItem(LOCAL_PLAEYR_ID_KEY, playerID)
        }).catch(e => {
            alert('Ooops! curse of demos :)')
            console.error(e.message)
        })

        return
    }

    api.getPlayer(playerID).then((player) => {
        console.log(player)
        document.getElementById('player_text_id').innerText = player.id
        document.getElementById('player_text_ranking').innerText = player.ranking
    }).catch(e => {
        alert('Ooops! curse of demos :)')
        console.error(e.message)
    })
}

main()