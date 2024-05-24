package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/kudarap/rockpapershit"
)

func (c *Client) CreateGame(ctx context.Context, game *rockpapershit.Game) error {
	sqlStatement := `INSERT INTO games (player_id_1, player_id_2, created_at) VALUES ($1, $2, $3) returning id`
	var id int
	err := c.db.QueryRow(ctx, sqlStatement, game.PlayerID1, game.PlayerID2, game.CreatedAt).Scan(&id)
	if err != nil {
		return err
	}

	game.ID = fmt.Sprintf("%d", id)
	fmt.Println("game", game)
	return nil
}

const (
	withFilter    = `SELECT id, player_id_1, player_id_2, player_cast_1, player_cast_2, created_at where player_id_1 = $1 OR player_id_2 =$1 FROM games`
	withoutFilter = `SELECT id, player_id_1, player_id_2, player_cast_1, player_cast_2, created_at  FROM games`
)

func (c *Client) Games(ctx context.Context, playerID string) ([]rockpapershit.Game, error) {
	q := withoutFilter
	var args []interface{}
	if playerID != "" {
		q = withFilter
		args = append(args, playerID)
	}

	rows, err := c.db.Query(ctx, q, args...)
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
	var cast1 sql.NullString
	var cast2 sql.NullString

	var game rockpapershit.Game
	game.ID = id
	err := c.db.
		QueryRow(ctx, `SELECT id, player_id_1, player_id_2, player_cast_1, player_cast_2, created_at FROM games WHERE id=$1`, id).
		Scan(&game.ID, &game.PlayerID1, &game.PlayerID2, &cast1, &cast2, &game.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, rockpapershit.ErrNotFound
		}
		return nil, err
	}

	game.PlayerCast1 = cast1.String
	game.PlayerCast2 = cast2.String
	return &game, nil
}

func (c *Client) Cast(ctx context.Context, throw, playerN, gameID string) (*rockpapershit.Game, error) {
	sqlStatement := fmt.Sprintf(`Update games SET %s = $1 where id = $2`, playerN)

	var createdGame rockpapershit.Game
	err := c.db.QueryRow(ctx, sqlStatement, throw, gameID).Scan(&createdGame)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
