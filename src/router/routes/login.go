package routes

import (
	"api/src/controllers"
	"net/http"
)

var LoginRoutes = Route{
	URI:                   "/login",
	Method:                http.MethodPost,
	Function:              controllers.Login,
	RequireAuthentication: false,
}
