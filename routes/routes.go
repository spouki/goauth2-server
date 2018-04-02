package routes

import "net/http"
import "github.com/spouki/goauth2/handlers"


type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func GetRoutes() Routes{
	var rroutes = Routes{
		Route{
			"Index", // Name
			"GET", // Method
			"/", // Pattern
			handlers.Index, // HandlerFunc
		},
		Route{
			"Registration",
			"POST",
			"/auth/registration",
			handlers.Registration,
		},
		Route{
			"Login",
			"POST",
			"/auth/login",
			handlers.Login,
		},
		Route{
			"Logout",
			"GET",
			"/auth/logout",
			handlers.Logout,
		},
		Route{
			"PasswordRequest",
			"POST",
			"/auth/password",
			handlers.PasswordRequest,
		},
	}
	return rroutes
}
