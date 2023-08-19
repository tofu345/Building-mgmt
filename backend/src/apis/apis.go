package apis

import (
	"github.com/tofu345/Building-mgmt-backend/src/models"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = models.GetDB()
}
