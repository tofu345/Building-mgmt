package internal

import (
	"time"
)

var models = []any{
	User{},
	Location{},
	Room{},
}

type User struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Email       string    `json:"email" gorm:"unique"`
	Password    string    `json:"-"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	IsSuperuser bool      `json:"is_superuser"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (user *User) Name() string {
	return user.FirstName + " " + user.LastName
}

func (u *User) requiredFields() []string {
	return []string{"email", "first_name", "last_name"}
}

type Location struct {
	ID      uint   `gorm:"primarykey" json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	AdminID uint   `json:"admin"`
	Rooms   []Room `json:"rooms" gorm:"foreignKey:ID"`
}

func (l *Location) requiredFields() []string {
	return []string{"name", "address"}
}

type Room struct {
	ID             uint      `gorm:"primarykey" json:"id"`
	Name           string    `json:"name"`
	OwnerID        uint      `json:"owner"`
	TenantID       uint      `json:"tenant"`
	TenancyEndDate time.Time `json:"tenancy_end_date"`
}

func (r *Room) requiredFields() []string {
	return []string{"name"}
}
