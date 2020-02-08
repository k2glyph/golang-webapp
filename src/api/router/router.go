package router

import (
	"github.com/gorilla/mux"
	"github.com/k2glyph/blog/src/api/router/routes"
)

// Create New Router
func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutes(r)
}
