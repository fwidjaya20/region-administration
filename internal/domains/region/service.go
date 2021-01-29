package region

import (
	"context"
	"github.com/fwidjaya20/regional-administration/internal/domains/region/repositories"
	"github.com/fwidjaya20/regional-administration/internal/globals"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type service struct {
	actor      string
	logger     log.Logger
	repository repositories.Interface
}

func (s *service) GetRegion(ctx context.Context, payload globals.RegionalRequest) (interface{}, error) {
	logger := log.With(s.logger, "METHOD", "GetRegion()", "SCOPE", payload.Scope)
	_ = level.Info(logger).Log("")

	var result interface{}
	var err error

	m, _ := s.repository.GetRegional(ctx)

	var strategy IScopeStrategy = NullScopeStrategy{}

	if payload.Scope == "province" {
		strategy = ProvinceScopeStrategy{}
	}

	result = strategy.Parse(ctx, m, payload.WithHierarchy)

	return result, err
}

func New(
	logger log.Logger,
	repository repositories.Interface,
) UseCase {
	service := service{
		actor: "REGIONAL",
		repository: repository,
	}

	service.logger = log.With(logger, "ACTOR", service.actor)

	return &service
}