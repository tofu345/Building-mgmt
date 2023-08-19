package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/tofu345/Building-mgmt-backend/src/services"
)

var routes = [][]Route{
	locationRoutes,
	userRoutes,
	roomRoutes,
}

func RegisterRoutes(r *mux.Router) {
	routesList := []Route{}
	for _, others := range routes {
		routesList = append(routesList, others...)
	}

	for _, route := range routesList {
		if route.middleware != nil && len(route.middleware) > 0 {
			route.function = middlewareWrapper(route.function, route.middleware...)
		}

		r.HandleFunc(route.url, route.function).Methods(route.methods...)
	}
}

type Route struct {
	url        string
	methods    []string
	function   Handler
	middleware []Middleware
}

type Handler func(http.ResponseWriter, *http.Request)

type Middleware func(http.ResponseWriter, *http.Request) error

func middlewareWrapper(handler Handler, mdws ...Middleware) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, v := range mdws {
			err := v(w, r)
			if err != nil {
				services.BadRequest(w, err)
				return
			}
		}

		handler(w, r)
	}
}
