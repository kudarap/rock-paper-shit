package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kudarap/rockpapershit"
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
		var player rockpapershit.Player
		if err := decodeJSONReq(r, &player); err != nil {
			encodeJSONResp(w, err, http.StatusBadRequest)
			return
		}

		if err := s.CreatePlayer(r.Context(), &player); err != nil {
			encodeJSONResp(w, err, http.StatusBadRequest)
			return
		}

		encodeJSONResp(w, player, http.StatusCreated)
	}
}

func ListPlayers(s service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
