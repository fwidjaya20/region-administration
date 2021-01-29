package repositories

import (
	"context"
	"github.com/fwidjaya20/regional-administration/internal/databases/models/regional"
)

type Interface interface {
	GetRegional(ctx context.Context) ([]*regional.Model, error)
}