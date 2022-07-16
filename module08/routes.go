package module08

import (
	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

type Routes []Route

func AllRoutes() Routes {
	routes := Routes{
		Route{"healthz", "GET", "/healthz", healthz},
		Route{"ip", "GET", "/ip", ip},
	}
	return routes
}
