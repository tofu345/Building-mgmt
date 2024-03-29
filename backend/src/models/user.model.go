package models

import (
	"errors"
	"time"

	"github.com/tofu345/Building-mgmt-backend/src/constants"
	"gorm.io/gorm"
)

type User struct {
	ID          uint      `json:"-" gorm:"primarykey"`
	Email       string    `json:"email" gorm:"unique" validate:"required,email"`
	Password    string    `json:"-" validate:"required,pswd"`
	FirstName   string    `json:"first_name" validate:"required"`
	LastName    string    `json:"last_name" validate:"required"`
	IsSuperuser bool      `json:"is_superuser"`
	OwnedRooms  []Room    `json:"owned_rooms" gorm:"foreignKey:OwnerID"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

func (user *User) Name() string {
	return user.FirstName + " " + user.LastName
}

func GetUserByEmail(email string) (User, error) {
	var user User
	err := db.Model(&User{}).Preload("OwnedRooms").First(&user, "email = ?", email).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, constants.ErrUserNotFound
	}
	return user, err
}

func GetUserByID(id uint) (User, error) {
	user := User{ID: id}
	err := db.First(&user).Error
	return user, err
}

func GetAdmins() ([]User, error) {
	var users []User
	err := db.Where("is_superuser = ?", true).Find(&users).Error
	return users, err
}

func GetUserList() ([]User, error) {
	var users []User
	model := db.Model(&User{}).Preload("OwnedRooms").First(&users)
	return users, model.Error
}

func CreateUser(user *User) error {
	return db.Create(user).Error
}

func UpdateUser(user *User) error {
	return db.Save(user).Error
}
