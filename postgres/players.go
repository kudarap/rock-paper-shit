package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	rockpapershit "github.com/kudarap/rockpapershit"
)

func (c *Client) Players(ctx context.Context) (*[]rockpapershit.Player, error) {
	rows, err := c.db.Query(ctx, `SELECT id, ranking, wins, loses, draws, plays_count FROM players`)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, rockpapershit.ErrFighterNotFound
		}
		return nil, err
	}
	defer rows.Close()

	var players []rockpapershit.Player

	for rows.Next() {
		var player rockpapershit.Player
		if err := rows.Scan(&player.ID, &player.Ranking, &player.Wins,
			&player.Loses, &player.Draws, &player.PlaysCount); err != nil {
			return &players, err
		}
		players = append(players, player)
	}

	return &players, nil
}
