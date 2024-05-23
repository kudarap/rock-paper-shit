package server

import (
	"context"

	"github.com/kudarap/rockpapershit"
)

type game interface {
	ListGames(ctx context.Context, id []string) *[]rockpapershit.Game // list
	CreateGame(ctx context.Context) *rockpapershit.Game               //post
	JoinGame(ctx context.Context, id string) *rockpapershit.Game      //get
	Cast(ctx context.Context, cast string)                            //Patch
}
