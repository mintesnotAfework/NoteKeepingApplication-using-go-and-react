package helpers

import (
	"github.com/mintesnotAfework/NoteKeepingApplication-using-go-and-react/models"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	value, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		panic("The hashing is not working")
	}
	return string(value), err
}

func ValidateUser(user *models.User) bool {
	if user.Password == "" || user.FirstName == "" || user.LastName == "" || user.Email == "" || (user.UserType != "ADMIN" && user.UserType != "USER") {
		return false
	}
	return true
}
