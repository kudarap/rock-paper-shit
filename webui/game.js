import api from './api.js'
import local from './local.js'
import { connectSocket } from './ws.js'

const TIME_LIMIT_SEC = 5;
const SECONDS = 1000;

(async ()=>{
    let game = null
    let player = null
    let you = null
    let enemy = null

    const timerElement = document.getElementById('timer')
    timerElement.innerHTML = TIME_LIMIT_SEC + 's';

    const resultContainer = document.getElementsByClassName('result-container')[0]

    const menuContainer = document.getElementsByClassName('menu-container')[0]

    const playerID = local.getPlayer()
    try {
        player = await api.getPlayer(playerID)
    } catch (e) {
        alert('Ooops! curse of demos :)')
        console.error(e.message)
        return
    }

    const gameID = window.location.hash.replace('#', '')
    document.getElementById('game_id').innerHTML = '#'+gameID;
    try {
        game = await api.getGame(gameID)
        if (game.player_id_1 == playerID) {
            you = player
            enemy = { id: game.player_id_2 }
        } else {
            you = player
            enemy = { id: game.player_id_1 }
        }
       
        const ticker = setInterval(() => {
            const d = (Date.parse(game.created_at) + TIME_LIMIT_SEC * SECONDS) - (new Date()).getTime();
            const t = Math.round(d / SECONDS)
            if (t < 1) {
                clearInterval(ticker)

                disableCastBtn()

                // get game result
                api.getGame(gameID).then(g => {
                    let youThrow = 'X'
                    let enemyThrow = 'X'
                    if (game.player_id_1 == playerID) {
                        youThrow = game.player_cast_1
                        enemyThrow = game.player_cast_2
                    } else {
                        youThrow = game.player_cast_2
                        enemyThrow = game.player_cast_1
                    }

                    youThrow = youThrow ? youThrow : 'X'
                    enemyThrow = enemyThrow ? enemyThrow : 'X'

                    document.getElementsByClassName('cast-container')[0].style.display = 'none'
                    resultContainer.style.display = 'flex'
                    resultContainer.innerHTML = '<div><button>'+youThrow.toUpperCase()+'</button></div>' 
                        + '<div><button>'+enemyThrow.toUpperCase()+'</button></div>' 
                    timerElement.innerHTML = evalResult(you, g)

                    menuContainer.style.display = 'flex'
                })
                return
            }
            timerElement.innerHTML = t + 's';
        }, SECONDS)
    } catch (e) {
        alert('Ooops! curse of demos :)')
        console.error(e.message)
        return
    }

    try {
        enemy = await api.getPlayer(enemy.id)
    } catch (e) {
        alert('Ooops! curse of demos :)')
        console.error(e.message)
        return
    }

    // render
    console.log({
        game,
        player,
        you,
        enemy,
    })
    document.getElementById('player1_text_id').innerText = you.id
    document.getElementById('player1_text_ranking').innerHTML = you.ranking+'<small>mmr</small>'
    document.getElementById('player2_text_id').innerText = enemy.id
    document.getElementById('player2_text_ranking').innerHTML = enemy.ranking+'<small>mmr</small>'
})()


function evalResult(you, game) {
    if (game.is_draw || game.winner == '') {
        return '<span style="color: grey;">DRAW!</span>'
    }
    if (game.winner == you.id) {
        return '<span style="color: green;">YOU WON!</span>'
    }
    return '<span style="color: crimson;">YOU LOSE!</span>'
}

const castRockBtn = document.getElementById('cast_rock')
const castPaperBtn = document.getElementById('cast_paper')
const castShitBtn = document.getElementById('cast_shit')

castRockBtn.addEventListener('click', () => {
    highlightBtn(castRockBtn)
    castPaperBtn.parentElement.remove()
    castShitBtn.parentElement.remove()
})
castPaperBtn.addEventListener('click', () => {
    highlightBtn(castPaperBtn)
    castRockBtn.parentElement.remove()
    castShitBtn.parentElement.remove()
})
castShitBtn.addEventListener('click', () => {
    highlightBtn(castShitBtn)
    castPaperBtn.parentElement.remove()
    castRockBtn.parentElement.remove()
})

function highlightBtn(el) {
    el.style.backgroundColor = 'lightYellow';
    el.style.border = '1px solid orange';
    el.style.borderRadius = '10px';
}

function disableCastBtn() {
    castRockBtn.setAttribute('disabled', 'disabled');
    castPaperBtn.setAttribute('disabled', 'disabled');
    castShitBtn.setAttribute('disabled', 'disabled');
}

export const findMatch = document.getElementById('find_match')
findMatch.addEventListener('click', () => {
    const playerID = local.getPlayer()
    findMatch.innerText = 'Finding...'
    connectSocket(playerID)
})

