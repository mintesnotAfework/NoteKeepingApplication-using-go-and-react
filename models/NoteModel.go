package models

import (
	"errors"

	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Content string `gorm:"not null;size::64" validate:"required,min=4,max=64" json:"content"`
	UserId  int64  `gorm:"not null;size::64" validate:"required,min=4,max=64" json:"user_id"`
}

func (n *Note) CreateNote() (*Note, error) {
	if checkUserExists(n.UserId) {
		db.Create(&n)
		return n, nil
	}
	return n, errors.New("User Not found")
}

func GetAllNote() []Note {
	var notes []Note
	db.Find(&notes)
	return notes
}

func GetAllNoteByUserId(user_id int64) []Note {
	var notes []Note
	db.Where("user_id=?", user_id).Find(&notes)
	return notes
}

func GetNoteById(id int64) Note {
	var note Note
	db.Where("ID=?", id).Find(&note)
	return note
}
func (n *Note) UpdateNoteById(id int64) (*Note, error) {
	var note Note
	if checkNoteExists(id) {
		responseDb := db.Where("ID=?", id).Find(&note)
		if n.Content != "" {
			note.Content = n.Content
		}
		responseDb.Save(&note)
	} else {
		return nil, errors.New("no such data")
	}
	return &note, nil
}

func UpdateNoteUserId(id int64, user_id int64) (*Note, error) {
	var note Note
	if checkNoteExists(id) && checkUserExists(user_id) {
		responseDb := db.Where("ID=?", id).Find(&note)
		note.UserId = user_id
		responseDb.Save(&note)
	} else {
		return nil, errors.New("no such data or user")
	}
	return &note, nil
}

func DeleteNoteById(id int64) *Note {
	var note Note

	db.Where("ID=?", id).Find(&note)
	db.Where("ID=?", id).Delete(&Note{})

	return &note
}

func checkNoteExists(id int64) bool {
	var count int64

	err := db.Model(&Note{}).Where("ID = ?", id).Count(&count).Error
	if err != nil {
		return false
	}

	return count > 0
}
