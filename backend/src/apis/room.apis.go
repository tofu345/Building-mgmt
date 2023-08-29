package apis

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	m "github.com/tofu345/Building-mgmt-backend/src/models"
	s "github.com/tofu345/Building-mgmt-backend/src/services"
)

func GetRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	room_id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	room, err := m.GetRoom(room_id)
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	s.Success(w, room)
}

func UpdateRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	room_id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.BadRequest(w, err)
		return
	}

	room, err := m.GetRoom(room_id)
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
