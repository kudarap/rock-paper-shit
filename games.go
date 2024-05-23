package rockpapershit

import (
	"time"
)

type Game struct {
	ID          string    `json:"id"`
	Winner      string    `json:"winner"`
	IsDraw      bool      `json:"is_draw"`
	PlayerID1   string    `json:"player_id_1"`
	PlayerID2   string    `json:"player_id_2"`
	PlayerCast1 string    `json:"player_cast_1"`
	PlayerCast2 string    `json:"player_cast_2"`
	CreatedAt   time.Time `json:"created_at"`
}

const (
	ThrowRock  = "rock"
	ThrowPaper = "paper"
	ThrowShit  = "shit"
)

func (g Game) setResult() Game {
	if g.PlayerCast1 != "" && g.PlayerCast2 == "" {
		g.Winner = g.PlayerID1
		return g
	} else if g.PlayerCast2 != "" && g.PlayerCast1 == "" {
		g.Winner = g.PlayerID2
		return g
	}

	if g.PlayerCast1 == g.PlayerCast2 {
		g.IsDraw = true
	} else if g.PlayerCast1 == ThrowRock {
		if g.PlayerCast2 == ThrowShit {
			g.Winner = g.PlayerID1
			return g
		}
		g.Winner = g.PlayerCast2
		return g
	} else if g.PlayerCast1 == ThrowPaper {
		if g.PlayerCast2 == ThrowRock {
			g.Winner = g.PlayerID1
			return g
		}
		g.Winner = g.PlayerCast2
		return g
	} else if g.PlayerCast1 == ThrowShit {
		if g.PlayerCast2 == ThrowPaper {
			g.Winner = g.PlayerID1
			return g
		}
		g.Winner = g.PlayerCast2
		return g
	}

	return g
}
