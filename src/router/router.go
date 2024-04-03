package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Gerar vai retornar um router com as rotas configuradas
func Gerar() *mux.Router {
	return routes.ConfigureRoutes(mux.NewRouter())
}
