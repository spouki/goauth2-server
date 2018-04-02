package router

import (
	"net/http"

	"github.com/gorilla/mux"
//	"github.com/spouki/goauth2/handlers"
//	"github.com/spouki/goauth2/routes"
	"github.com/spouki/goauth2/logger"
)
import r "github.com/spouki/goauth2/routes"

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range r.GetRoutes() {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}

// EOF
