package rockpapershit

import (
	"time"
)

type Game struct {
	ID          string    `json:"id"`
	PlayerID1   string    `json:"player_id_1"`
	PlayerID2   string    `json:"player_id_2"`
	PlayerCast1 string    `json:"player_cast_1"`
	PlayerCast2 string    `json:"player_cast_2"`
	CreatedAt   time.Time `json:"created_at"`
}

type GameRequest struct {
	Game
}
