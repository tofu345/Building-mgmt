package apis

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	m "github.com/tofu345/Building-mgmt-backend/src/models"
	s "github.com/tofu345/Building-mgmt-backend/src/services"
	"gorm.io/gorm"
)

func GetLocations(w http.ResponseWriter, r *http.Request) {
	locations, err := m.GetLocations()
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	s.Success(w, locations)
}

func GetLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	loc_id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	loc, err := m.GetLocation(loc_id)
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	s.Success(w, loc)
}

func CreateLocation(w http.ResponseWriter, r *http.Request) {
	postData := s.GetContextData(r, "validated_data").(map[string]string)
	loc := m.Location{
		Name:    postData["name"],
		Address: postData["address"],
	}

	loc.ID = 0
	err := m.CreateLocation(&loc)
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	s.Success(w, loc)
}

func UpdateLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	loc_id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	loc, err := m.GetLocation(loc_id)
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	err = s.JsonDecode(r, &loc)
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	loc.ID = uint(loc_id)
	errSlice := s.ValidateModel(loc)
	if errSlice != nil {
		s.BadRequest(w, errSlice)
		return
	}

	err = m.UpdateLocation(&loc)
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	s.Success(w, loc)
}

func GetLocationRooms(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	loc_id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	loc, err := m.GetLocation(loc_id)
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	s.Success(w, loc.Rooms)
}

func CreateRoomForLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	loc_id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	_, err = m.GetLocation(loc_id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.BadRequest(w, "Location not found")
			return
		}

		s.BadRequest(w, err)
		return
	}

	postData := s.GetContextData(r, "validated_data").(map[string]string)
	owner_id, err := strconv.Atoi(postData["owner_id"])
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	_, err = m.GetUserByID(uint(owner_id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.BadRequest(w, "User not found")
			return
		}

		s.BadRequest(w, err)
		return
	}

	room := m.Room{
		Name:       postData["name"],
		OwnerID:    uint(owner_id),
		LocationID: uint(loc_id),
	}
	err = m.CreateRoom(&room)
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	s.Success(w, room)
}
