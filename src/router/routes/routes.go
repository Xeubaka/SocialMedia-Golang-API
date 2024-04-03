package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents all routes of the API
type Route struct {
	URI                   string
	Method                string
	Function              func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}

// ConfigureRoutes configure all routes for the API
func ConfigureRoutes(r *mux.Router) *mux.Router {
	routes := getAllRoutes()

	for _, route := range routes {
		if route.RequireAuthentication {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticates(route.Function)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}

func getAllRoutes() (routes []Route) {
	routes = append(routes, LoginRoutes)
	routes = append(routes, UserRoutes...)
	routes = append(routes, PostsRoutes...)

	return
}
