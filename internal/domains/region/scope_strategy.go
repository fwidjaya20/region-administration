package region

import (
	"context"
	"github.com/fwidjaya20/regional-administration/internal/databases/models/regional"
	"github.com/fwidjaya20/regional-administration/internal/globals"
)

type IScopeStrategy interface {
	Parse(ctx context.Context, models []*regional.Model, withHierarchy bool) []globals.RegionalResponse
}
