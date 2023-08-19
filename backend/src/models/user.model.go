package models

import "time"

type User struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Email       string    `json:"email" gorm:"unique" validate:"required,email"`
	Password    string    `json:"-" validate:"required,pswd"`
	FirstName   string    `json:"first_name" validate:"required"`
	LastName    string    `json:"last_name" validate:"required"`
	IsSuperuser bool      `json:"is_superuser"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (user *User) Name() string {
	return user.FirstName + " " + user.LastName
}
