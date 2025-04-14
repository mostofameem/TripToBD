package web

import (
	"net/http"
	"restaurant-service/web/middlewares"
)

func (server *Server) initRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"GET /hello",
		manager.With(
			http.HandlerFunc(server.handlers.Hello),
		),
	)

	mux.Handle(
		"POST /add-restaurant",
		manager.With(
			http.HandlerFunc(server.handlers.AddRestaurent),
		),
	)

	mux.Handle(
		"GET /get-restaurant/{id}", // id is restaurant id
		manager.With(
			http.HandlerFunc(server.handlers.GetRestaurant),
		),
	)

	mux.Handle(
		"GET /get-restaurants",
		manager.With(
			http.HandlerFunc(server.handlers.GetRestaurants),
		),
	)

	mux.Handle(
		"GET /search",
		manager.With(
			http.HandlerFunc(server.handlers.Search),
		),
	)

	mux.Handle(
		"GET /get-restaurants/location/{id}",
		manager.With(
			http.HandlerFunc(server.handlers.GetRestaurantsByLocation),
		),
	)
	mux.Handle(
		"POST /review-restaurant",
		manager.With(
			http.HandlerFunc(server.handlers.AddReview),
		),
	)
	mux.Handle(
		"GET /Update-restaurant/{id}",
		manager.With(
			http.HandlerFunc(server.handlers.UpdateRestaurant),
		),
	)
	mux.Handle(
		"POST /restaurant/add-branch",
		manager.With(
			http.HandlerFunc(server.handlers.AddBranch),
		),
	)
}
