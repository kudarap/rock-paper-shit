package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

type matchmakingService interface {
	FindMatch(ctx context.Context, playerID uuid.UUID) error
}

func FindMatch(s matchmakingService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		playerID := r.URL.Query().Get("player_id")
		if playerID == "" {
			encodeJSONError(w, errors.New("player_id is required"), http.StatusBadRequest)
			return
		}

		playerUUID, err := uuid.Parse(playerID)
		if err != nil {
			encodeJSONError(w, errors.New("invalid player_id"), http.StatusBadRequest)
			return
		}

		err = s.FindMatch(r.Context(), playerUUID)
		if err != nil {
			encodeJSONError(w, errors.New("error adding player to queue"), http.StatusInternalServerError)
			return
		}

		response := map[string]string{"status": "queued"}
		encodeJSONResp(w, response, http.StatusOK)
	}
}
