import idgen from './idgen.js'
import api from './api.js'
import local from './local.js'

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

// emulate 
// setTimeout(() => {
//     window.location = '/game#912j0921ke09jk'
// }, 3000)