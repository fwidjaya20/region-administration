package repositories

import (
	"context"
	"github.com/fwidjaya20/regional-administration/internal/databases/models/province"
)

type Interface interface {
	GetProvincesByCode(ctx context.Context, code []string) ([]*province.Model, error)
	GetProvincesByName(ctx context.Context, name string) ([]*province.Model, error)
}