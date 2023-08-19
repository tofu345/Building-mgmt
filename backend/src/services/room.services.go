package services

import m "github.com/tofu345/Building-mgmt-backend/src/models"

func GetRoom(id int) (m.Room, error) {
	room := m.Room{ID: uint(id)}
	err := db.First(&room).Error
	if err != nil {
		return m.Room{}, err
	}

	return room, nil
}
