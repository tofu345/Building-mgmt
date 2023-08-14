package internal

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func locations(w http.ResponseWriter, r *http.Request) {
	objs := []Location{}
	err := db.Find(&objs).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		jsonError(w, err)
		return
	}

	jsonResponse(w, 200, objs)
}

func createLocation(w http.ResponseWriter, r *http.Request) {
	var location Location
	err := jsonDecode(r, &location)
	if err != nil {
		jsonError(w, err)
		return
	}

	data, valid := location.validate()
	if !valid {
		jsonError(w, data)
		return
	}

	err = db.Create(&location).Error
	if err != nil {
		jsonError(w, err)
		return
	}

	jsonResponse(w, 200, location)
}

func updateLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		jsonError(w, err)
		return
	}

	location := Location{ID: uint(id)}
	err = db.First(&location).Error
	if err != nil {
		jsonError(w, err)
		return
	}

	err = jsonDecode(r, &location)
	if err != nil {
		jsonError(w, err)
		return
	}

	data, valid := location.validate()
	if !valid {
		jsonError(w, data)
		return
	}

	err = db.Save(&location).Error
	if err != nil {
		jsonError(w, err)
		return
	}

	jsonResponse(w, 200, location)
}

func getTokens(w http.ResponseWriter, r *http.Request) {
	postData := map[string]string{}
	err := jsonDecode(r, &postData)
	if err != nil {
		jsonError(w, err)
		return
	}

	errors_map := map[string]string{}
	if _, exists := postData["email"]; !exists {
		errors_map["email"] = "This field is required"
	}
	if _, exists := postData["password"]; !exists {
		errors_map["password"] = "This field is required"
	}
	if len(errors_map) != 0 {
		jsonError(w, errors_map)
		return
	}

	var user User
	err = db.First(&user, "email = ?", postData["email"]).Error
	if err != nil {
		jsonError(w, err)
		return
	}

	errorMap, valid := user.validate()
	if !valid {
		jsonError(w, errorMap)
		return
	}

	if !CheckPasswordHash(postData["password"], user.Password) {
		jsonError(w, "InvalidLogin")
		return
	}

	access, err := newAccessToken(user)
	if err != nil {
		jsonError(w, TokenError)
		return
	}

	refresh, err := newRefreshToken(user)
	if err != nil {
		jsonError(w, TokenError)
		return
	}

	jsonResponse(w, 200, map[string]any{"access": access, "refresh": refresh})
}

func refreshToken(w http.ResponseWriter, r *http.Request) {
	postData := map[string]string{}
	err := jsonDecode(r, &postData)
	if err != nil {
		jsonError(w, err)
		return
	}

	if _, exists := postData["token"]; !exists {
		jsonError(w, map[string]string{"token": RequiredField})
		return
	}

	payload, err := decodeToken(postData["token"])
	if err != nil {
		jsonError(w, err)
		return
	}

	if _, exists := payload["ref"]; !exists {
		jsonError(w, InvalidToken)
		return
	}

	email := payload["email"]
	switch email := email.(type) {
	case string:
		access, err := newAccessToken(User{Email: email})
		if err != nil {
			jsonError(w, err)
			return
		}

		jsonResponse(w, 200, map[string]any{"access": access})
	default:
		jsonError(w, InvalidToken)
	}
}
