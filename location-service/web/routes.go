package web

import (
	"net/http"
	"post-service/web/middlewares"
)

func (server *Server) initRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"GET /hello",
		manager.With(
			http.HandlerFunc(server.handlers.Hello),
		),
	)

	mux.Handle(
		"POST /add-location",
		manager.With(
			http.HandlerFunc(server.handlers.AddLocation),
		),
	)

	mux.Handle(
		"GET /get-location",
		manager.With(
			http.HandlerFunc(server.handlers.GetLocation),
		),
	)

	mux.Handle(
		"GET /get-locations",
		manager.With(
			http.HandlerFunc(server.handlers.GetLocations),
		),
	)
	mux.Handle(
		"POST /add-review",
		manager.With(
			http.HandlerFunc(server.handlers.AddReview),
		),
	)
	mux.Handle(
		"POST /add-route",
		manager.With(
			http.HandlerFunc(server.handlers.AddRoute),
		),
	)
	mux.Handle(
		"GET /get-route",
		manager.With(
			http.HandlerFunc(server.handlers.GetRoute),
		),
	)
}
