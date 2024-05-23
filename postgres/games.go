package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	rockpapershit "github.com/kudarap/foo"
)

func (c *Client) Games(ctx context.Context) (*[]rockpapershit.Game, error) {
	rows, err := c.db.Query(ctx, `SELECT id, player_id_1, player_id_2, player_1_cast, player_2_cast, created_at FROM games`)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, rockpapershit.ErrFighterNotFound
		}
		return nil, err
	}
	defer rows.Close()

	var games []rockpapershit.Game

	for rows.Next() {
		var game rockpapershit.Game
		if err := rows.Scan(&game.ID, &game.PlayerID1, &game.PlayerID2, &game.Player1Cast, &game.Player2Cast, &game.CreatedAt); err != nil {
			return &games, err
		}
		games = append(games, game)
	}

	return &games, nil
}

func (c *Client) Game(ctx context.Context, id string) (*rockpapershit.Game, error) {
	var game rockpapershit.Game
	game.ID = id
	err := c.db.
		QueryRow(ctx, `SELECT id, player_id_1, player_id_2, player_1_cast, player_2_cast, created_at FROM games WHERE id=$1`, id).
		Scan(&game.ID, &game.PlayerID1, &game.PlayerID2, &game.Player1Cast, &game.Player2Cast, &game.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, rockpapershit.ErrFighterNotFound
		}
		return nil, err
	}

	return &game, nil
}
