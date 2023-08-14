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

func (user *User) validate() (map[string]string, bool) {
	errors := map[string]string{}

	if user.Email == "" {
		errors["email"] = RequiredField
	}

	if user.FirstName == "" {
		errors["first_name"] = RequiredField
	}

	if user.LastName == "" {
		errors["last_name"] = RequiredField
	}

	if user.Password == "" {
		errors["password"] = RequiredField
	}

	return errors, len(errors) == 0
}

type Location struct {
	ID      uint   `gorm:"primarykey" json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	AdminID uint   `json:"admin"`
	Rooms   []Room `json:"rooms" gorm:"foreignKey:ID"`
}

func (location *Location) validate() (map[string]string, bool) {
	errors := map[string]string{}

	if location.Name == "" {
		errors["name"] = RequiredField
	}

	if location.Address == "" {
		errors["address"] = RequiredField
	}

	return errors, len(errors) == 0
}

type Room struct {
	ID             uint      `gorm:"primarykey" json:"id"`
	Name           string    `json:"name"`
	OwnerID        uint      `json:"owner"`
	TenantID       uint      `json:"tenant"`
	TenancyEndDate time.Time `json:"tenancy_end_date"`
}
