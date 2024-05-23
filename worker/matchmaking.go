package worker

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/kudarap/rockpapershit"
	"github.com/kudarap/rockpapershit/redis"
	r "github.com/redis/go-redis/v9"
)

type Matchmaker struct {
	Redis   *redis.Client
	Service *rockpapershit.Service // game service
	Logger  *slog.Logger
}

func (m *Matchmaker) Run() {
	m.Logger.Info("matchmaking worker is now running...")
	ctx := context.Background()

	for {
		p1, err := m.Redis.BRPop(ctx, 0, "matchmaking_queue").Result()
		if err != nil {
			m.Logger.DebugContext(ctx, "error retrieving player from queue", err)
			continue
		}

		player1ID, err := uuid.Parse(p1[1])
		if err != nil {
			m.Logger.DebugContext(ctx, "error parsing player ID", err)
			continue
		}

		time.Sleep(5 * time.Second) // 5 seconds delay

		p2, err := m.Redis.RPop(ctx, "matchmaking_queue").Result()
		if err != nil {
			if errors.Is(err, r.Nil) {
				m.Redis.LPush(ctx, "matchmaking_queue", player1ID)
				continue
			}

			m.Logger.DebugContext(ctx, "error retrieving opponent from queue", err)
			continue
		}

		player2ID, err := uuid.Parse(p2)
		if err != nil {
			m.Logger.DebugContext(ctx, "error parsing opponent ID", err)
			continue
		}

		gameID := uuid.New()
		game := rockpapershit.Game{
			ID:        gameID.String(),
			PlayerID1: player1ID.String(),
			PlayerID2: player2ID.String(),
			CreatedAt: time.Time{},
		}

		err = m.Service.CreateGame(ctx, &game)
		if err != nil {
			m.Logger.DebugContext(ctx, "error creating match", err)
		} else {
			m.Logger.InfoContext(ctx, "match created between player", "player1", player1ID, "player2", player2ID)
			m.notifyPlayers(ctx, player1ID, player2ID, gameID)
		}
	}
}

func (m *Matchmaker) notifyPlayers(ctx context.Context, p1Id, p2Id, gameId uuid.UUID) {
	m.Redis.Publish(ctx, "matchmaking_notifications", p1Id.String()+":"+gameId.String())
	m.Redis.Publish(ctx, "matchmaking_notifications", p2Id.String()+":"+gameId.String())
}
