package rockpapershit

import (
	"time"
)

type Game struct {
	ID          string
	PlayerID1   string
	PlayerID2   string
	Player1Cast string
	Player2Cast string
	CreatedAt   time.Time
}
