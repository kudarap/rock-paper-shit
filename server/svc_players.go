package server

import (
	"context"

	rockpapershit "github.com/kudarap/foo"
)

type players interface {
	CreatePlayer(ctx context.Context) (*rockpapershit.Player, error)
	ListPlayers(ctx context.Context) ([]*rockpapershit.Player, error)
	GetPlayer(ctx context.Context, id string) (*rockpapershit.Player, error)
	UpdateRanking(ctx context.Context, id string) (*rockpapershit.Player, error)
}
