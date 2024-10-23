package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName  string `gorm:"not null;size::64" validate:"required,min=4,max=64" json:"first_name"`
	LastName   string `gorm:"not null;size::64" validate:"required,min=4,max=64" json:"last_name"`
	MiddleName string `gorm:"size::64" validate:"min=4,max=64" json:"middle_name"`
	Email      string `gorm:"unique;not null;size:64" validate:"email,required,min=4,max=64" json:"email"`
	Password   string `gorm:"not null" validate:"required" json:"password"`
	Phone      string `gorm:"not null" validate:"required,min=10,max=15" json:"phone"`
	UserType   string `gorm:"not null" validate:"required,eq=ADMIN|eq=USER" json:"user_type"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (u *User) CreateUser() *User {
	db.Create(&u)
	return u
}

func GetAllUsers() []User {
	var users []User
	db.Find(&users)
	return users
}

func GetUserById(id int64) *User {
	var user User
	db.Where("ID=?", id).Find(&user)
	return &user

}

func GetUserByEmail(email string) *User {
	var user User
	db.Where("email=?", email).Find(&user)
	return &user

}

func (u *User) UpdateUserById(id int64) (*User, error) {
	var user User
	if checkUserExists(id) {
		responseDb := db.Where("ID=?", id).Find(&user)
		if u.FirstName != "" {
			user.FirstName = u.FirstName
		}
		if u.LastName != "" {
			user.LastName = u.LastName
		}
		if u.MiddleName != "" {
			user.MiddleName = u.MiddleName
		}
		if u.Email != "" {
			user.Email = u.Email
		}
		if u.Password != "" {
			user.Password = u.Password
		}
		if u.Phone != "" {
			user.Phone = u.Phone
		}
		if u.UserType != "" {
			user.UserType = u.UserType
		}
		responseDb.Save(&user)
	} else {
		return nil, errors.New("no such data")
	}
	return &user, nil
}

func DeleteUserById(id int64) *User {
	var user User

	db.Where("ID=?", id).Find(&user)
	db.Where("ID=?", id).Delete(&User{})

	return &user
}

func checkUserExists(id int64) bool {
	var count int64

	err := db.Model(&User{}).Where("ID = ?", id).Count(&count).Error
	if err != nil {
		return false
	}

	return count > 0
}
