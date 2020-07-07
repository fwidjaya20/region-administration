package http

import (
	"github.com/fwidjaya20/regional-administration/internal/domains/region"
	"github.com/fwidjaya20/regional-administration/internal/domains/region/endpoints"
	"github.com/fwidjaya20/regional-administration/internal/globals"
	"github.com/fwidjaya20/regional-administration/lib/server"
	kitHttp "github.com/go-kit/kit/transport/http"
	"net/http"
)

func GetRegion(service region.UseCase, opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		server.NewHTTPServer(endpoints.GetRegion(service), server.HTTPOption{
			DecodeModel: &globals.RegionalRequest{},
		}, opts).ServeHTTP(w, r)
	}
}
