package models

import "time"

type Room struct {
	ID   uint   `gorm:"primarykey" json:"id"`
	Name string `json:"name" validate:"required"`
	// Owner          User      `json:"owner"  gorm:"foreignKey:ID"`
	// Tenant         User      `json:"tenant"  gorm:"foreignKey:ID"`
	TenancyEndDate time.Time `json:"tenancy_end_date"`
}

func GetRoom(id int) (Room, error) {
	room := Room{ID: uint(id)}
	err := db.First(&room).Error
	if err != nil {
		return Room{}, err
	}

	return room, nil
}

func CreateRoom(r *Room) error {
	return db.Create(r).Error
}

func UpdateRoom(r *Room) error {
	return db.Save(r).Error
}
