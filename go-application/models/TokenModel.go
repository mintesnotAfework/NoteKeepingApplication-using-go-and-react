package models

import (
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	Refresh string `gorm:"unique not null;size::64" validate:"required,min=4,max=64"`
	UserId  uint   `gorm:"not null;size::64" validate:"required,min=4,max=64"`
}

func (t *Token) CreateToken() *Token {
	db.Create(&t)
	return t
}

func GetRefreshTokenByUserId(id uint) *Token {
	var token Token
	db.Where("user_id=?", id).Find(&token)
	return &token
}

func DeleteRefreshTokenByUserId(user_id uint64) {
	db.Where("user_id=?", user_id).Delete(&Token{})
}

func DeleteTokenById(id uint64) {
	db.Where("ID=?", id).Delete(&Token{})
}

func checkTokenExists(refresh string) bool {
	var count int64

	err := db.Model(&Token{}).Where("refresh = ?", refresh).Count(&count).Error
	if err != nil {
		return false
	}

	return count != 1
}

func checkUserRefreshTokenExists(id uint) bool {
	var count int64

	err := db.Model(&Token{}).Where("user_id = ?", id).Count(&count).Error
	if err != nil {
		return false
	}

	return count != 1
}
