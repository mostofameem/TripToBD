package web

import (
	"hotel-service/web/middlewares"
	"net/http"
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
		"POST /hotel",
		manager.With(
			http.HandlerFunc(server.handlers.AddHotel),
		),
	)

	// mux.Handle(
	// 	"GET /hotel",
	// 	manager.With(
	// 		http.HandlerFunc(server.handlers.GetHotel),
	// 	),
	// )

	mux.Handle(
		"GET /hotel",
		manager.With(
			http.HandlerFunc(server.handlers.GetHotels),
		),
	)

	mux.Handle(
		"POST /review",
		manager.With(
			http.HandlerFunc(server.handlers.AddReview),
		),
	)

	mux.Handle(
		"GET /review",
		manager.With(
			http.HandlerFunc(server.handlers.GetReview),
		),
	)

	mux.Handle(
		"POST /pics",
		manager.With(
			http.HandlerFunc(server.handlers.AddPic),
		),
	)

	mux.Handle(
		"GET /pics",
		manager.With(
			http.HandlerFunc(server.handlers.GetPic),
		),
	)

}
