package schemas

var GetTokenPair = map[string]any{
	"email":    "required",
	"password": "required",
}

var RefreshToken = map[string]any{
	"token": "required",
}
