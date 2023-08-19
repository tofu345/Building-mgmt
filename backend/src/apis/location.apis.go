package apis

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	m "github.com/tofu345/Building-mgmt-backend/src/models"
	s "github.com/tofu345/Building-mgmt-backend/src/services"
)

func GetLocations(w http.ResponseWriter, r *http.Request) {
	locations, err := s.GetLocations()
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	s.Success(w, locations)
}

func GetLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	loc, err := s.GetLocation(id)
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	s.Success(w, loc)
}

func GetLocationRooms(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	loc, err := s.GetLocation(id)
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	s.Success(w, loc.Rooms)
}

func CreateLocation(w http.ResponseWriter, r *http.Request) {
	var loc m.Location
	err := s.JsonDecode(r, &loc)
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	errSlice := s.ValidateModel(loc)
	if errSlice != nil {
		s.BadRequest(w, errSlice)
		return
	}

	loc.ID = 0
	err = db.Create(&loc).Error
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	s.Success(w, loc)
}

func UpdateLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	loc, err := s.GetLocation(id)
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	err = s.JsonDecode(r, &loc)
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	loc.ID = uint(id)
	errSlice := s.ValidateModel(loc)
	if errSlice != nil {
		s.BadRequest(w, errSlice)
		return
	}

	err = db.Save(&loc).Error
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	s.Success(w, loc)
}
