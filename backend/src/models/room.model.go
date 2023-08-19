package models

import "time"

type Room struct {
	ID   uint   `gorm:"primarykey" json:"id"`
	Name string `json:"name" validate:"required"`
	// Owner          User      `json:"owner"  gorm:"foreignKey:ID"`
	// Tenant         User      `json:"tenant"  gorm:"foreignKey:ID"`
	TenancyEndDate time.Time `json:"tenancy_end_date"`
}
