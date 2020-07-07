package endpoints

import (
	"context"
	"github.com/fwidjaya20/regional-administration/internal/domains/region"
	"github.com/fwidjaya20/regional-administration/internal/globals"
	"github.com/fwidjaya20/regional-administration/lib/database"
	libHttp "github.com/fwidjaya20/regional-administration/lib/transport/http"
	"github.com/go-kit/kit/endpoint"
)

func GetRegion(service region.UseCase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		payload := request.(*globals.RegionalRequest)

		err = database.RunInTransaction(ctx, globals.DB(), func(ctx context.Context) error {
			response, err = service.GetRegion(ctx, *payload)
			return err
		})

		return libHttp.Response(ctx, response, nil), err
	}
}