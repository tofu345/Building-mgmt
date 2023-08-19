package routes

import (
	"github.com/tofu345/Building-mgmt-backend/src/apis"
	mdw "github.com/tofu345/Building-mgmt-backend/src/middleware"
)

var locationRoutes = []Route{
	{"/locations", []string{"GET"}, apis.GetLocations, []Middleware{mdw.AuthRequired}},
	{"/locations", []string{"POST"}, apis.CreateLocation, []Middleware{mdw.AuthRequired}},
	{"/locations/{id}", []string{"GET"}, apis.GetLocation, []Middleware{mdw.AuthRequired}},
	{"/locations/{id}", []string{"PUT"}, apis.UpdateLocation, []Middleware{mdw.AuthRequired}},
	{"/locations/{id}/rooms", []string{"GET"}, apis.GetLocationRooms, []Middleware{mdw.AuthRequired}},
	{"/locations/{id}/rooms", []string{"POST"}, apis.CreateRoomForLocation, []Middleware{mdw.AuthRequired}},
}
