package routes

import "github.com/tofu345/Building-mgmt-backend/src/apis"

var userRoutes = []Route{
	{url: "/users", methods: []string{"GET"}, function: apis.GetUserList},
	{url: "/token", methods: []string{"POST"}, function: apis.GenerateTokenPair},
	{url: "/token/refresh", methods: []string{"POST"}, function: apis.RegenerateAccessToken},
}
