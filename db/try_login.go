package db

import (
	"github.com/cristhianforerod/udemy-go-course/models"
	"golang.org/x/crypto/bcrypt"
)

func TryLogin(email string, password string) (models.User, bool) {
	user, found, _ := CheckUser(email)
	if !found {
		return user, false
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, false
	}
	return user, true
}
