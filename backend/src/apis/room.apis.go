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

	room, err := s.GetRoom(id)
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

	room, err := s.GetRoom(id)
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

	err = db.Save(&room).Error
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

	loc, err := s.GetLocation(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.BadRequest(w, "Location not found")
		} else {
			s.BadRequest(w, err)
		}
		return
	}

	err = db.Create(&room).Error
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	loc.Rooms = append(loc.Rooms, room)
	err = db.Save(&loc).Error
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	s.Success(w, room)
}
