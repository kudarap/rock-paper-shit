package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/kudarap/rockpapershit"
)

func (c *Client) Player(ctx context.Context, playerID string) (*rockpapershit.Player, error) {
	var player rockpapershit.Player
	player.ID = playerID
	err := c.db.
		QueryRow(ctx, `SELECT id, ranking, wins, loses, draws, plays_count FROM players WHERE id=$1`, playerID).
		Scan(&player.ID, &player.Ranking, &player.Wins, &player.Loses, &player.Draws, &player.PlaysCount)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, rockpapershit.ErrNotFound
		}
		return nil, err
	}

	return &player, nil
}

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

func (c *Client) CreatePlayer(ctx context.Context, player *rockpapershit.Player) error {
	sqlStatement := `INSERT INTO players (id, ranking, wins, loses, draws, plays_count) VALUES ($1, $2, $3, $4, $5, $6) `
	_, err := c.db.Exec(ctx, sqlStatement, player.ID, player.Ranking, player.Wins, player.Loses, player.Draws, player.PlaysCount)
	if err != nil {
		return err
	}
	return nil

}

func (c *Client) CalcRanking(ctx context.Context, player string, mmr int) {
	sqlStatement := fmt.Sprintf(`Update players SET ranking = $1 where id = $2`, mmr, player)
	_, err := c.db.Query(ctx, sqlStatement)
	if err != nil {
		fmt.Errorf(`error calculating rank: %v`, err)
	}
}
