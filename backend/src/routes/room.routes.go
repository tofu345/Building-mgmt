package routes

import (
	"github.com/tofu345/Building-mgmt-backend/src/apis"
	mdw "github.com/tofu345/Building-mgmt-backend/src/middleware"
)

var roomRoutes = []Route{
	{"/rooms/{id}", []string{"GET"}, apis.GetRoom, []Middleware{mdw.AuthRequired}},
	{"/rooms/{id}", []string{"PUT"}, apis.UpdateRoom, []Middleware{mdw.AuthRequired}},
}
