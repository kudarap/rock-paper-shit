package rockpapershit

import (
	"time"

	"github.com/google/uuid"
)

type Player struct {
	ID        uuid.UUID `json:"id"`
	Ranking   int       `json:"ranking"`
	Wins      int       `json:"wins"`
	Losses    int       `json:"losses"`
	Draws     int       `json:"draws"`
	PlayCount int       `json:"play_count"`
}

type Match struct {
	ID       uuid.UUID `json:"id"`
	PlayerID uuid.UUID `json:"player_id"`
	Status   string    `json:"status"`
	DateTime time.Time `json:"datetime"`
}

type Game struct {
	ID          uuid.UUID `json:"id"`
	GameID      uuid.UUID `json:"game_id"`
	Winner      uuid.UUID `json:"winner"`
	Player1ID   uuid.UUID `json:"player1_id"`
	Player2ID   uuid.UUID `json:"player2_id"`
	Player1Cast string    `json:"player1_cast"`
	Player2Cast string    `json:"player2_cast"`
	DateTime    time.Time `json:"datetime"`
}
