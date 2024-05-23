package rockpapershit

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
)

// Service represents foo service.
type Service struct {
	repo   repository
	logger *slog.Logger
}

// NewService returns new foo service.
func NewService(r repository, l *slog.Logger) *Service {
	return &Service{repo: r, logger: l}
}

// ListGames returns a list of games
func (s *Service) ListGames(ctx context.Context) (*[]Game, error) {
	s.logger.InfoContext(ctx, "listing all games")

	g, err := s.repo.Games(ctx)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, ErrNotFound.X(err)
		}
		return nil, fmt.Errorf("could not list games on repository: %s", err)
	}
	return g, nil
}

// CreateGame creates a game and returns the game details
func (s *Service) CreateGame(ctx context.Context) (*Game, error) {
	s.logger.InfoContext(ctx, "create game")
	g, err := s.repo.CreateGame(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not create game: %s", err)
	}

	return g, nil
}

// JoinGame joins/reconnects to a game and returns the game details
func (s *Service) JoinGame(ctx context.Context, id string) (*Game, error) {
	s.logger.InfoContext(ctx, "getting game by id", "id", id)

	p, err := s.repo.Game(ctx, id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, ErrNotFound.X(err)
		}
		return nil, fmt.Errorf("could not find game on repository: %s", err)
	}
	return p, nil
}

// Cast updates player_cast and returns game details
func (s *Service) Cast(ctx context.Context, id string) (*Game, error) {
	return nil, nil
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
	CreateGame(ctx context.Context) (*Game, error)
	Games(ctx context.Context) (*[]Game, error)
	Game(ctx context.Context, gameID string) (*Game, error)
	Players(ctx context.Context) (*[]Player, error)
	Player(ctx context.Context, playerID string) (*Player, error)
}
