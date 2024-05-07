package web

import (
	"net/http"
	"project/web/middlewares"
)

func StartServer() *http.ServeMux {
	manager := middlewares.NewManager()
	mux := http.NewServeMux()
	InitRoutes(mux, manager)
	return mux
}
