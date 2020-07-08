package repositories

import (
	"context"
	"github.com/fwidjaya20/regional-administration/internal/databases/models/regional"
)

type Interface interface {
	GetVillages(ctx context.Context, code []string, name string) ([]*regional.Model, error)
}