package router

import (
	"github.com/carlos1377/devbook/api/router/routes"
	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	r := mux.NewRouter()

	return routes.ConfigRoutes(r)
}
