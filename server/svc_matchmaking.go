package server

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/kudarap/rockpapershit/redis"
)

var wsUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type MatchmakingHandler struct {
	Redis  redis.Client
	Logger *slog.Logger
}

type Player struct {
	ID         uint
	Ranking    int
	Wins       int
	Loses      int
	Draws      int
	PlaysCount int
}

func (h *MatchmakingHandler) FindMatch(w http.ResponseWriter, r *http.Request) {
	ws, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		h.Logger.DebugContext(context.Background(), "error upgrading to WebSocket", err)
		http.Error(w, "could not open websocket conn", http.StatusBadRequest)
		return
	}
	defer ws.Close()

	var player Player
	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		ws.WriteJSON(map[string]string{"error": "invalid request"})
		return
	}

	ctx := r.Context()
	err = h.Redis.LPush(ctx, "matchmaking_queue", player.ID).Err()
	if err != nil {
		ws.WriteJSON(map[string]string{"error": err.Error()})
		return
	}

	ws.WriteJSON(map[string]string{"message": "Player added to queue"})

	pubsub := h.Redis.Subscribe(ctx, "matchmaking_notifications")
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			h.Logger.DebugContext(context.Background(), "error receiving matchmaking notification", err)
			break
		}

		data := strings.Split(msg.Payload, ":")
		playerID, err := strconv.Atoi(data[0])
		if err != nil {
			h.Logger.DebugContext(context.Background(), "error parsing player id from notification", err)
			continue
		}

		if player.ID == uint(playerID) {
			gameID := data[1]
			ws.WriteJSON(map[string]string{"message": "Match found", "game_id": gameID})
			break
		}
	}
}
