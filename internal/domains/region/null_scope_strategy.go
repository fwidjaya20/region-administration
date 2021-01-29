package region

import (
	"context"
	"fmt"
	"github.com/fwidjaya20/regional-administration/internal/databases/models/regional"
	"github.com/fwidjaya20/regional-administration/internal/globals"
)

type NullScopeStrategy struct {}

func (NullScopeStrategy) Parse(ctx context.Context, models []*regional.Model, withHierarchy bool) []globals.RegionalResponse {
	fmt.Println("[NullScopeStrategy]")
	return []globals.RegionalResponse{}
}
