package rockpapershit

type Player struct {
	ID         string `json:"id"`
	Ranking    int    `json:"ranking"`
	Wins       int    `json:"wins"`
	Loses      int    `json:"loses"`
	Draws      int    `json:"draws"`
	PlaysCount int    `json:"plays_count"`
}
