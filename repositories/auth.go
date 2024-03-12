package repositories

import (
	"github.com/che-ict/DEV-DT-Microblog/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username, password, displayName string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 11) // foutieve functie

	if err != nil {
		return err
	}

	user := models.User{
		Username:    username,
		Password:    string(hash),
		DisplayName: displayName,
	}

	db := DB()

	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func LoginUser(username, password string) bool {
	var user models.User
	DB().Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false
	}

	return true
}
