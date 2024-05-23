package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kudarap/rockpapershit"
)

func JoinGame(s service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		game, err := s.GetGame(r.Context(), id)
		if err != nil {
			encodeJSONError(w, err, http.StatusBadRequest)
		}
		encodeJSONResp(w, game, http.StatusOK)
	}
}

func ListGames(s service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		games, err := s.ListGames(r.Context(), "")
		if err != nil {
			encodeJSONError(w, err, http.StatusBadRequest)
			return
		}
		encodeJSONResp(w, games, http.StatusOK)
	}
}

func Cast(s service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		games, err := s.Cast(r.Context(), "rock", id)
		if err != nil {
			encodeJSONError(w, err, http.StatusBadRequest)
			return
		}
		encodeJSONResp(w, games, http.StatusOK)
		return
	}
}

func CurrentGame(s service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		playerID := r.Header.Get("player_id")
		games, err := s.ListGames(r.Context(), playerID)
		if err != nil {
			encodeJSONError(w, err, http.StatusBadRequest)
			return
		}
		encodeJSONResp(w, games, http.StatusOK)
		return
	}
}

func CreatePlayer(s service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var player rockpapershit.Player
		if err := s.CreatePlayer(r.Context(), &player); err != nil {
			encodeJSONError(w, err, http.StatusBadRequest)
			return
		}

		encodeJSONResp(w, player, http.StatusOK)
		return
	}
}
