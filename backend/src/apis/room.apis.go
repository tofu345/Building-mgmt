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

func GetRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	room, err := m.GetRoom(id)
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	s.Success(w, room)
}

func UpdateRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	room, err := m.GetRoom(id)
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	err = s.JsonDecode(r, &room)
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	errs := s.ValidateModel(room)
	if errs != nil {
		s.BadRequest(w, errs)
		return
	}

	err = m.UpdateRoom(&room)
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	s.Success(w, room)
}

func CreateRoomForLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	var room m.Room
	err = s.JsonDecode(r, &room)
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	errs := s.ValidateModel(room)
	if errs != nil {
		s.BadRequest(w, errs)
		return
	}

	loc, err := m.GetLocation(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.BadRequest(w, "Location not found")
		} else {
			s.BadRequest(w, err)
		}
		return
	}

	err = m.CreateRoom(&room)
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	loc.Rooms = append(loc.Rooms, room)
	err = m.UpdateLocation(&loc)
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	s.Success(w, room)
}
