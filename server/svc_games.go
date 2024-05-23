package server

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/kudarap/rockpapershit"
)

type game interface {
	ListGames(ctx context.Context, id []string) *[]rockpapershit.Game // list
	CreateGame(ctx context.Context) *rockpapershit.Game               //post
	JoinGame(ctx context.Context, id string) *rockpapershit.Game      //get
	Cast(ctx context.Context, cast string)                            //Patch
}

func JoinGame(s service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		encodeJSONResp(w, rockpapershit.Game{
			ID:          id,
			PlayerID1:   "KinagatNgDragon",
			PlayerID2:   "LumpiangKidlat",
			PlayerCast1: "",
			PlayerCast2: "",
			CreatedAt:   time.Now(),
		}, http.StatusCreated)
	}
}

func ListGames(s service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
