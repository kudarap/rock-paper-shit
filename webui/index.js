import idgen from './idgen.js'
import api from './api.js'
import local from './local.js'
import { connectSocket } from './ws.js'

function main() {
    let playerID = local.getPlayer()
    console.log(playerID)

    if (!playerID) {
        playerID = idgen()
        api.createPlayer(playerID).then((player) => {
            console.log(player)
            document.getElementById('player_text_id').innerText = player.id
            document.getElementById('player_text_ranking').innerText = player.ranking
            local.setPlayer(player.id)
        }).catch(e => {
            alert('Ooops! curse of demos 1:)')
            console.error(e.message)
        })
        return
    }

    api.getPlayer(playerID).then((player) => {
        console.log(player)
        document.getElementById('player_text_id').innerText = player.id
        document.getElementById('player_text_ranking').innerText = player.ranking

        // connectSocket(playerID)
    }).catch(e => {
        alert('Ooops! curse of demos 2:)')
        console.error(e.message)
    })
}

main()

const findMatch = document.getElementById('find_match')
findMatch.addEventListener('click', () => {
    const playerID = local.getPlayer()
    findMatch.innerText = 'Finding...'
    connectSocket(playerID)
})