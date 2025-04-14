package web

import (
	"net/http"
	"vehicles/web/middlewares"
)

func (server *Server) initRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"GET /hello",
		manager.With(
			http.HandlerFunc(server.handlers.Hello),
		),
	)

	mux.Handle(
		"GET /api/v1/stats",
		manager.With(
			http.HandlerFunc(server.handlers.Stats),
		),
	)

	mux.Handle(
		"POST /add-vehicle",
		manager.With(
			http.HandlerFunc(server.handlers.AddVehicle),
		),
	)

	mux.Handle(
		"GET /get-vehicle",
		manager.With(
			http.HandlerFunc(server.handlers.GetVehicle),
		),
	)

	mux.Handle(
		"GET /get-vehicles",
		manager.With(
			http.HandlerFunc(server.handlers.GetVehicles),
		),
	)

	mux.Handle(
		"POST /add-routes",
		manager.With(
			http.HandlerFunc(server.handlers.AddRoutes),
			//middlewares.Authenticate,
		),
	)

	mux.Handle(
		"GET /get-routes/vehicle",
		manager.With(
			http.HandlerFunc(server.handlers.GetRoutesByVehicle),
			//middlewares.Authenticate,
		),
	)

	mux.Handle(
		"GET /get-routes/location",
		manager.With(
			http.HandlerFunc(server.handlers.GetRoutesByLocation),
			//middlewares.Authenticate,
		),
	)

	mux.Handle(
		"POST /add-review",
		manager.With(
			http.HandlerFunc(server.handlers.AddReview),
		),
	)

	mux.Handle(
		"GET /get-review",
		manager.With(
			http.HandlerFunc(server.handlers.GetReviewByID),
		),
	)

	mux.Handle(
		"POST /add-pics",
		manager.With(
			http.HandlerFunc(server.handlers.AddPic),
		),
	)

	mux.Handle(
		"GET /get-pics",
		manager.With(
			http.HandlerFunc(server.handlers.GetPic),
		),
	)

}
