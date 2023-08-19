package services

import m "github.com/tofu345/Building-mgmt-backend/src/models"

func GetUserByEmail(email string) (m.User, error) {
	var user m.User
	err := db.First(&user, "email = ?", email).Error
	return user, err
}

func GetUserList() ([]m.User, error) {
	var users []m.User
	err := db.Find(&users).Error
	return users, err
}
