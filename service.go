package rockpapershit

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/kudarap/rockpapershit/redis"
)

const GameTimeOut = 5 * time.Second

// Service represents foo service.
type Service struct {
	repo   repository
	cache  cache
	redis  *redis.Client
	logger *slog.Logger
}

// NewService returns new foo service.
func NewService(r repository, c cache, redis *redis.Client, l *slog.Logger) *Service {
	return &Service{repo: r, cache: c, redis: redis, logger: l}
}

// ListGames returns a list of games
func (s *Service) ListGames(ctx context.Context, playerID string) ([]Game, error) {
	s.logger.InfoContext(ctx, "listing all games")
	g, err := s.repo.Games(ctx, playerID)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, ErrNotFound.X(err)
		}
		return nil, fmt.Errorf("could not list games on repository: %s", err)
	}
	current := time.Now()
	for k, v := range g {
		if playerID != "" && v.CreatedAt.Add(GameTimeOut).After(current) {
			continue
		}
		g[k] = v.setResult()
	}
	return g, nil
}

// CreateGame creates a game and returns the game details
func (s *Service) CreateGame(ctx context.Context, game *Game) error {
	s.logger.InfoContext(ctx, "create game")
	game.CreatedAt = time.Now()
	err := s.repo.CreateGame(ctx, game)
	if err != nil {
		return fmt.Errorf("could not create game: %s", err)
	}

	return nil
}

// GetGame joins/reconnects to a game and returns the game details
func (s *Service) GetGame(ctx context.Context, id string) (*Game, error) {
	s.logger.InfoContext(ctx, "getting game by id", "id", id)

	g, err := s.repo.Game(ctx, id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, ErrNotFound.X(err)
		}
		return nil, fmt.Errorf("could not find game on repository: %s", err)
	}
	g1 := g.setResult()
	return &g1, nil
}

// Cast updates player_cast and returns game details
func (s *Service) Cast(ctx context.Context, throw, playerID string) (*Game, error) {
	s.logger.InfoContext(ctx, "cast vote", "id", playerID)

	//get game by id
	game, err := s.GetGame(ctx, playerID)
	if err != nil {
		return nil, err
	}
	var playerN string
	if game.Winner == game.PlayerID1 {
		playerN = "player_cast_1"
	} else {
		playerN = "player_cast_2"
	}

	game, err = s.repo.Cast(ctx, throw, playerN, game.ID)
	if err != nil {
		return nil, fmt.Errorf("player could not cast: %s", err)
	}

	// check if game completed player then eval ranking
	if game.setResult().Winner != "" {
		winner, err := s.GetPlayerByID(ctx, game.Winner)
		if err != nil {
			return nil, err
		}
		s.repo.CalcRanking(ctx, game.Winner, winner.Ranking+10)

		loser, err := s.GetPlayerByID(ctx, game.loser())
		if err != nil {
			return nil, err
		}
		s.repo.CalcRanking(ctx, game.loser(), loser.Ranking-10)
	}

	return game, nil
}

// CreatePlayer creates a player and returns player details
func (s *Service) CreatePlayer(ctx context.Context) (*Player, error) {
	return nil, nil
}

// ListPlayers returns a list of players
func (s *Service) ListPlayers(ctx context.Context) (*[]Player, error) {
	s.logger.InfoContext(ctx, "listing all players")

	p, err := s.repo.Players(ctx)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, ErrNotFound.X(err)
		}
		return nil, fmt.Errorf("could not list players on repository: %s", err)
	}
	return p, nil
}

// UpdateRanking calculates ranking and returns player details
func (s *Service) UpdateRanking(ctx context.Context, id string) (*Player, error) {

	return nil, nil
}

// GetPlayerByID returns a player by id
func (s *Service) GetPlayerByID(ctx context.Context, id string) (*Player, error) {
	s.logger.InfoContext(ctx, "getting player by id", "id", id)

	p, err := s.repo.Player(ctx, id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, ErrNotFound.X(err)
		}
		return nil, fmt.Errorf("could not find player on repository: %s", err)
	}
	return p, nil
}

// FighterByID returns a fighter by id.
func (s *Service) FighterByID(ctx context.Context, sid string) (*Fighter, error) {
	// NOTE this is a just a demo logging and should use InfoContext enabling telemetry logs.
	s.logger.InfoContext(ctx, "getting foo fighter by id", "id", sid)

	id, err := uuid.Parse(sid)
	if err != nil {
		return nil, err
	}

	f, err := s.repo.Fighter(ctx, id)
	if err != nil {
		if errors.Is(err, ErrFighterNotFound) {
			return nil, ErrFighterNotFound.X(err)
		}
		return nil, fmt.Errorf("could not find fighter on repository: %s", err)
	}
	return f, nil
}

// repository manages storage operation for fighters.
type repository interface {
	Fighter(ctx context.Context, id uuid.UUID) (*Fighter, error)
	CreateGame(ctx context.Context, game *Game) error
	Games(ctx context.Context, playerID string) ([]Game, error)
	Game(ctx context.Context, gameID string) (*Game, error)
	Cast(ctx context.Context, throw, playerN, gameID string) (*Game, error)
	CalcRanking(ctx context.Context, winner string, mmr int)
	Players(ctx context.Context) (*[]Player, error)
	Player(ctx context.Context, playerID string) (*Player, error)
}

type cache interface {
	FindMatch(ctx context.Context, playerID uuid.UUID) error
}
