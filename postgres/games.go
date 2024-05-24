package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

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
	fmt.Println("GameGameGameGame", id)
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
	fmt.Println("repo Cast", throw, playerN, gameID)

	idn, err := strconv.Atoi(gameID)
	if err != nil {
		return nil, err
	}

	sqlStatement := fmt.Sprintf(`Update games SET %s = $1 where id = $2 RETURNING id, player_id_1, player_id_2, player_cast_1, player_cast_2, created_at`, playerN)
	fmt.Println("CAST", sqlStatement, idn)

	var cast1 sql.NullString
	var cast2 sql.NullString
	var createdGame rockpapershit.Game
	err = c.db.QueryRow(ctx, sqlStatement, throw, idn).Scan(
		&createdGame.ID,
		&createdGame.PlayerID1,
		&createdGame.PlayerID2,
		&cast1,
		&cast2,
		&createdGame.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("cast query row: %s", err)
	}

	createdGame.PlayerCast1 = cast1.String
	createdGame.PlayerCast2 = cast2.String
	return &createdGame, nil
}
