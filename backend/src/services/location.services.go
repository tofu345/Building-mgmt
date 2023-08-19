package services

import (
	"errors"

	"github.com/tofu345/Building-mgmt-backend/src/constants"
	m "github.com/tofu345/Building-mgmt-backend/src/models"
	"gorm.io/gorm"
)

func GetLocations() ([]m.Location, error) {
	rows := []m.Location{}
	err := db.Model(&m.Location{}).Preload("Rooms").First(&rows).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return []m.Location{}, constants.ErrObjectNotFound
	}

	return rows, nil
}

func GetLocation(id int) (m.Location, error) {
	obj := m.Location{ID: uint(id)}
	err := db.Model(&m.Location{}).Preload("Rooms").First(&obj).Error
	if err != nil {
		return m.Location{}, err
	}

	return obj, nil
}
