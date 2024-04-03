package middlewares

import (
	"api/src/authentication"
	"api/src/templates"
	"log"
	"net/http"
)

// Logger writes requests informations inside the router
func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunction(w, r)
	}
}

// Authenticates if a user is authenticated
func Authenticates(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.ValidateToken(r); err != nil {
			templates.Error(w, http.StatusUnauthorized, err)
			return
		}
		nextFunction(w, r)
	}
}
