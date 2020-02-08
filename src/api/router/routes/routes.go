package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route Struct
type Route struct {
	URI     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

func load() []Route {
	routes := userRoutes
	return routes
}

// Setup Routes
func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range load() {
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}
	return r
}
