package routes

import (
	"net/http"

	"github.com/k2glyph/blog/src/api/controllers"
)

var userRoutes = []Route{
	Route{
		URI:     "/users",
		Method:  http.MethodGet,
		Handler: controllers.GetUsers,
	},
	Route{
		URI:     "/user/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetUser,
	},
	Route{
		URI:     "/users",
		Method:  http.MethodPost,
		Handler: controllers.CreateUser,
	},
	Route{
		URI:     "/users/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateUser,
	},
	Route{
		URI:     "/users/{id}",
		Method:  http.MethodPut,
		Handler: controllers.DeleteUser,
	},
}
