package models

import (
	"github.com/mintesnotAfework/NoteKeepingApplication-using-go-and-react/config"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Note{})
}
