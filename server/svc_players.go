package server

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	rockpapershit "github.com/kudarap/rockpapershit"
)

type players interface {
	CreatePlayer(ctx context.Context) (*rockpapershit.Player, error)
	ListPlayers(ctx context.Context) ([]*rockpapershit.Player, error)
	GetPlayer(ctx context.Context, id string) (*rockpapershit.Player, error)
	UpdateRanking(ctx context.Context, id string) (*rockpapershit.Player, error)
}

func GetPlayerByID(s service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		encodeJSONResp(w, rockpapershit.Player{
			ID:         id,
			Ranking:    100,
			Wins:       3,
			Loses:      2,
			Draws:      1,
			PlaysCount: 6,
		}, http.StatusCreated)
	}
}

func PostPlayer(s service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input rockpapershit.Player
		if err := decodeJSONReq(r, &input); err != nil {
			encodeJSONResp(w, err, http.StatusBadRequest)
			return
		}

		encodeJSONResp(w, rockpapershit.Player{
			ID:         input.ID,
			Ranking:    100,
			Wins:       3,
			Loses:      2,
			Draws:      1,
			PlaysCount: 6,
		}, http.StatusCreated)
	}
}

func GetGameByID(s service) http.HandlerFunc {
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

func ListPlayers(s service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func ListGames(s service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
