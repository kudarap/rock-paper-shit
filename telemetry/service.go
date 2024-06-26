package telemetry

import (
	"context"
	"log/slog"

	"github.com/kudarap/rockpapershit"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

const serviceName = "fooservice"

type FooService struct {
	*rockpapershit.Service
	logger     *slog.Logger
	tracerName string
}

func (s *FooService) FighterByID(ctx context.Context, id string) (*rockpapershit.Fighter, error) {
	ctx, span := otel.Tracer(s.tracerName).Start(ctx, "fooservice.FighterByID")
	defer span.End()

	s.logger.DebugContext(ctx, "params", "fighter_id", id)

	f, err := s.Service.FighterByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	s.logger.DebugContext(ctx, "returns", "fighter", f)
	return f, nil
}

func TraceFooService(s *rockpapershit.Service, l *slog.Logger) *FooService {
	return &FooService{s, l, serviceName}
}
