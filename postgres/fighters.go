package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	rockpapershit "github.com/kudarap/rockpapershit"
)

func (c *Client) Fighter(ctx context.Context, id uuid.UUID) (*rockpapershit.Fighter, error) {
	var fighter rockpapershit.Fighter
	fighter.ID = id
	err := c.db.
		QueryRow(ctx, `SELECT id, first_name, last_name FROM fighters WHERE id=$1`, id.String()).
		Scan(&fighter.ID, &fighter.FirstName, &fighter.LastName)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, rockpapershit.ErrFighterNotFound
		}
		return nil, err
	}

	return &fighter, nil
}
