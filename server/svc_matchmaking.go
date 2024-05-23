package server

import (
	"errors"
	"net/http"
)

func FindMatch(s service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		playerID := r.URL.Query().Get("player_id")
		if playerID == "" {
			encodeJSONError(w, errors.New("player_id is required"), http.StatusBadRequest)
			return
		}

		if err := s.QueuePlayer(r.Context(), playerID); err != nil {
			encodeJSONError(w, errors.New("error adding player to queue"), http.StatusInternalServerError)
			return
		}

		response := map[string]string{"status": "queued"}
		encodeJSONResp(w, response, http.StatusOK)
	}
}
