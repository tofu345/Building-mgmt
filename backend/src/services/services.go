package services

import (
	m "github.com/tofu345/Building-mgmt-backend/src/models"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = m.GetDB()
}
