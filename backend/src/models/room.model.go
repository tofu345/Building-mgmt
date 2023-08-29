package models

import "time"

type Room struct {
	ID             uint      `json:"id" gorm:"primarykey"`
	Name           string    `json:"name" validate:"required"`
	OwnerID        uint      `json:"owner_id" validate:"required"`
	LocationID     uint      `json:"-" validate:"required"`
	TenantID       *uint     `json:"tenant_id" validate:"-"`
	Tenant         User      `json:"tenant" validate:"-" gorm:"foreignKey:TenantID"`
	TenancyEndDate time.Time `json:"tenancy_end_date" gorm:"autoCreateTime:false"`
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
