package models

import (
	"errors"

	"github.com/tofu345/Building-mgmt-backend/src/constants"
	"gorm.io/gorm"
)

type Location struct {
	ID      uint   `gorm:"primarykey" json:"id"`
	Name    string `json:"name" gorm:"unique" validate:"required"`
	Address string `json:"address" gorm:"unique" validate:"required"`
	AdminID uint   `json:"admin"`
	Rooms   []Room `json:"rooms" gorm:"foreignKey:ID"`
}

func GetLocations() ([]Location, error) {
	rows := []Location{}
	err := db.Model(&Location{}).Preload("Rooms").First(&rows).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return []Location{}, constants.ErrObjectNotFound
	}

	return rows, nil
}

func GetLocation(id int) (Location, error) {
	obj := Location{ID: uint(id)}
	err := db.Model(&Location{}).Preload("Rooms").First(&obj).Error
	if err != nil {
		return Location{}, err
	}

	return obj, nil
}

func CreateLocation(l *Location) error {
	return db.Create(l).Error
}

func UpdateLocation(l *Location) error {
	return db.Save(l).Error
}
