package worker

import (
	"context"
	"errors"
	"log/slog"
	"strconv"
	"time"

	"github.com/kudarap/rockpapershit"
	"github.com/kudarap/rockpapershit/redis"
	r "github.com/redis/go-redis/v9"
)

type Matchmaker struct {
	Redis   *redis.Client
	Service *rockpapershit.Service
	Logger  *slog.Logger
}

func (m *Matchmaker) Run() {
	ctx := context.Background()

	for {
		p1, err := m.Redis.BRPop(ctx, 0, "matchmaking_queue").Result()
		if err != nil {
			m.Logger.DebugContext(ctx, "error retrieving player from queue", err)
			continue
		}

		p1ID, err := strconv.Atoi(p1[1])
		if err != nil {
			m.Logger.DebugContext(ctx, "error parsing player ID", err)
			continue
		}

		time.Sleep(5 * time.Second)

		p2, err := m.Redis.RPop(ctx, "matchmaking_queue").Result()
		if err != nil {
			if errors.Is(err, r.Nil) {
				m.Redis.LPush(ctx, "matchmaking_queue", p1ID)
				continue
			}

			m.Logger.DebugContext(ctx, "error retrieving opponent from queue", err)
			continue
		}

		_, err = strconv.Atoi(p2)
		if err != nil {
			m.Logger.DebugContext(ctx, "error parsing opponent ID", err)
			continue
		}

		// TODO: create game.
		/*err = m.service.CreateGame((uint(p1ID), uint(p2ID))
		if err != nil {
			m.Logger.DebugContext(ctx, "error creating match", err)
		} else {
			m.Logger.InfoContext(ctx, "match created between player", "player1", p1ID, "player2", p2)
		}*/
	}
}
