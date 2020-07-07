package container

import (
	"github.com/fwidjaya20/regional-administration/internal/domains/region"
	provinceRepo "github.com/fwidjaya20/regional-administration/internal/domains/region/repositories"
	"github.com/go-kit/kit/log"
)

type Container struct {
	ProvinceService region.UseCase
}

func New(
	logger log.Logger,
) *Container {
	return &Container{
		ProvinceService: region.New(logger, provinceRepo.NewProvinceRepository()),
	}
}
