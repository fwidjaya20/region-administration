package http

import (
	"github.com/fwidjaya20/regional-administration/cmd/container"
	libServer "github.com/fwidjaya20/regional-administration/lib/server"
	"github.com/go-chi/chi"
	kitHttp "github.com/go-kit/kit/transport/http"
	"net/http"
)

func MakeHandler(
	router *chi.Mux,
	container *container.Container,
) http.Handler {
	opts := []kitHttp.ServerOption{
		kitHttp.ServerErrorEncoder(libServer.ErrorEncoder),
	}

	generateRegionalAdmRoute(router, container, opts)

	return router
}

func generateRegionalAdmRoute(router chi.Router, container *container.Container, opts []kitHttp.ServerOption) {
	router.Group(func(r chi.Router) {
	})
}