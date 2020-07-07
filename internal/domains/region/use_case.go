package region

import (
	"context"
	"github.com/fwidjaya20/regional-administration/internal/globals"
)

type UseCase interface {
	GetRegion(ctx context.Context, payload globals.RegionalRequest) (interface{}, error)
}