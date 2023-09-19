package models

import (
	"errors"
	"time"

	"github.com/tofu345/Building-mgmt-backend/src"
	"gorm.io/gorm"
)

type User struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	Email       string    `json:"email" gorm:"unique" validate:"required,email"`
	Password    string    `json:"-" validate:"required,pswd"`
	FirstName   string    `json:"first_name" validate:"required"`
	LastName    string    `json:"last_name" validate:"required"`
	IsSuperuser bool      `json:"is_superuser"`
	OwnedRooms  []Room    `json:"owned_rooms" gorm:"foreignKey:OwnerID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (user *User) Name() string {
	return user.FirstName + " " + user.LastName
}

func GetUserByEmail(email string) (User, error) {
	var user User
	err := db.Model(&User{}).Preload("OwnedRooms").First(&user, "email = ?", email).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, src.ErrUserNotFound
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
	err := db.Model(&User{}).Preload("OwnedRooms").First(&users).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return users, nil
	}
	return users, err
}

func CreateUser(user *User) error {
	return db.Create(user).Error
}

func UpdateUser(user *User) error {
	return db.Save(user).Error
}
