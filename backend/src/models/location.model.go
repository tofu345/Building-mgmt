package models

type Location struct {
	ID      uint   `gorm:"primarykey" json:"id"`
	Name    string `json:"name" gorm:"unique" validate:"required"`
	Address string `json:"address" gorm:"unique" validate:"required"`
	AdminID uint   `json:"admin"`
	Rooms   []Room `json:"rooms" gorm:"foreignKey:ID"`
}
