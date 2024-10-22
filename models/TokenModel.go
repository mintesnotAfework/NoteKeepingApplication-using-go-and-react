package models

import (
	"errors"

	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	value   string `gorm:"unique not null;size::64" validate:"required,min=4,max=64" json:"value"`
	refresh string `gorm:"unique not null;size::64" validate:"required,min=4,max=64" json:"refresh"`
	userId  uint64 `gorm:"not null;size::64" validate:"required,min=4,max=64" json:"user_id"`
}

func (t *Token) CreateToken() *Token {
	db.Create(&t)
	return t
}

func GetUserId(value string) (uint64, error) {
	if checkTokenExists(value) {
		var token Token
		db.Model(&Token{}).Where("value = ?", value).First(&token)
		return token.userId, nil
	} else {
		return 0, errors.New("token not found")
	}
}

func DeleteTokenByUserId(user_id uint64) {
	db.Where("user_id=?", user_id).Delete(&Token{})
}

func DeleteTokenById(id uint64) {
	db.Where("ID=?", id).Delete(&Token{})
}

func checkTokenExists(value string) bool {
	var count int64

	err := db.Model(&Token{}).Where("value = ?", value).Count(&count).Error
	if err != nil {
		return false
	}

	return count != 1
}
