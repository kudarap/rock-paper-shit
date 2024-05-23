package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/kudarap/rockpapershit"
)

func (c *Client) CreateGame(ctx context.Context, game *rockpapershit.Game) error {
	sqlStatement := `INSERT INTO games (id, player_id_1, player_id_2, created_at) VALUES ($1, $2, $3) returning id, player_id_1, player_id_2, created_at`
	err := c.db.QueryRow(ctx, sqlStatement, game.ID, game.PlayerID1, game.PlayerID2, game.CreatedAt).Scan(game)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Games(ctx context.Context) ([]rockpapershit.Game, error) {
	rows, err := c.db.Query(ctx, `SELECT id, player_id_1, player_id_2, player_cast_1, player_cast_2, created_at FROM games`)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, rockpapershit.ErrNotFound
		}
		return nil, err
	}
	defer rows.Close()

	var games []rockpapershit.Game

	for rows.Next() {
		var game rockpapershit.Game
		if err := rows.Scan(&game.ID, &game.PlayerID1, &game.PlayerID2, &game.PlayerCast1, &game.PlayerCast2, &game.CreatedAt); err != nil {
			return nil, err
		}
		games = append(games, game)
	}

	return games, nil
}

func (c *Client) Game(ctx context.Context, id string) (*rockpapershit.Game, error) {
	var game rockpapershit.Game
	game.ID = id
	err := c.db.
		QueryRow(ctx, `SELECT id, player_id_1, player_id_2, player_cast_1, player_cast_2, created_at FROM games WHERE id=$1`, id).
		Scan(&game.ID, &game.PlayerID1, &game.PlayerID2, &game.PlayerCast1, &game.PlayerCast2, &game.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, rockpapershit.ErrNotFound
		}
		return nil, err
	}

	return &game, nil
}

func (c *Client) Cast(ctx context.Context, throw string, player int) (*rockpapershit.Game, error) {
	//playerCast := fmt.Sprintf(`player_cast_%d`, player)
	//sqlStatement := fmt.Sprintf(`Update games SET player_cast_1, = $1 where player_id_1 = $2`, throw, throw)
	//var createdGame rockpapershit.Game
	//err := c.db.QueryRow(ctx, sqlStatement, game.ID, game.PlayerID1, game.PlayerID2, game.CreatedAt).Scan(&createdGame)
	//if err != nil {
	//
	//}
	return nil, nil
}
