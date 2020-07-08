package region

import (
	"context"
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
	case "regional":
	case "regency":
	case "district":
	case "village":
		result, err = s.VillageScope(ctx, payload)
	}

	return result, err
}

func (s *service) VillageScope(ctx context.Context, payload globals.RegionalRequest) ([]globals.VillageResponse, error) {
	logger := log.With(s.logger, "METHOD", "VillageScope()", "SCOPE", payload.Scope)

	var result = make([]globals.VillageResponse, 0)
	var err error

	m, err := s.repository.GetVillages(ctx, payload.Code, payload.Name)

	if nil != err {
		_ = level.Error(logger).Log("Error", err)
		return nil, libError.NewError(err, http.StatusInternalServerError, "get_village_error")
	}

	for _, root := range m {
		result = append(result, globals.VillageResponse{
			Code:       root.VillageModel.Code.String,
			Name:       root.VillageModel.Name.String,
			PostalCode: root.VillageModel.PostalCode.String,
			District:   &globals.DistrictResponse{
				Code:     root.DistrictModel.Code.String,
				Name:     root.DistrictModel.Name.String,
				Regency:  &globals.RegencyResponse{
					Code:      root.RegencyModel.Code.String,
					Name:      root.RegencyModel.Name.String,
					Province:  &globals.ProvinceResponse{
						Code:      root.ProvinceModel.Code.String,
						Name:      root.ProvinceModel.Name.String,
					},
				},
			},
		})
	}

	return result, nil
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