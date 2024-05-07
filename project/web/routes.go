package web

import (
	"net/http"
	"project/web/handlers"
	"project/web/middlewares"
)

func InitRoutes(mux *http.ServeMux, manager *middlewares.Manager) {

	mux.Handle(
		"POST /users",
		manager.With(
			http.HandlerFunc(handlers.Create),
		),
	)

	mux.Handle(
		"GET /users",
		manager.With(
			http.HandlerFunc(handlers.List),
		),
	)
	mux.Handle(
		"GET /users/",
		manager.With(
			http.HandlerFunc(handlers.GetUserByName),
		),
	)
	mux.Handle(
		"DELETE /users/{id}",
		manager.With(
			http.HandlerFunc(handlers.Delete),
		),
	)
}
