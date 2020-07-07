package region

import (
	"context"
	"fmt"
	"github.com/fwidjaya20/regional-administration/internal/databases/models/province"
	"github.com/fwidjaya20/regional-administration/internal/domains/region/repositories"
	"github.com/fwidjaya20/regional-administration/internal/globals"
	libError "github.com/fwidjaya20/regional-administration/lib/error"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"net/http"
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

	switch payload.Scope {
	case "province":
		result, err = s.ProvinceScope(ctx, payload)
	case "regency":
	case "district":
	case "village":
	}

	return result, err
}

type Foo struct {
	Result []globals.ProvinceResponse
}

func (s *service) ProvinceScope(ctx context.Context, payload globals.RegionalRequest) ([]globals.ProvinceResponse, error) {
	logger := log.With(s.logger, "METHOD", "ProvinceScope()", "SCOPE", payload.Scope)

	var result Foo
	var err error

	m, err := s.repository.GetProvincesByCode(ctx, payload.Code)

	if nil != err {
		_ = level.Error(logger).Log("Error", err)
		return nil, libError.NewError(err, http.StatusInternalServerError, "get_province_by_code_error")
	}

	return result.Result, nil
}
func New(
	logger log.Logger,
	repository repositories.Interface,
) UseCase {
	service := service{
		actor: "PROVINCE",
		repository: repository,
	}

	service.logger = log.With(logger, "ACTOR", service.actor)

	return &service
}