package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mintesnotAfework/NoteKeepingApplication-using-go-and-react/models"
	"github.com/mintesnotAfework/NoteKeepingApplication-using-go-and-react/utils"
)

func CreateNote(w http.ResponseWriter, r *http.Request, u *models.User) {
	w.Header().Set("Content-Type", "application/json")

	note := &models.Note{}
	utils.PareseBody(r, note)
	if u.UserType != "ADMIN" {
		note.UserId = int64(u.ID)
	}
	n, err := note.CreateNote()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(n)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func GetNoteByUserId(w http.ResponseWriter, r *http.Request, u *models.User) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if u.UserType != "ADMIN" && int64(u.ID) != int64(id) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	note := models.GetAllNoteByUserId(int64(id))
	res, err := json.Marshal(note)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetNoteById(w http.ResponseWriter, r *http.Request, u *models.User) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	note := models.GetNoteById(int64(id))
	if u.UserType != "ADMIN" && int64(u.ID) != int64(note.UserId) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	res, err := json.Marshal(note)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllNote(w http.ResponseWriter, r *http.Request, u *models.User) {
	w.Header().Set("Content-Type", "application/json")
	if u.UserType != "ADMIN" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	notes := models.GetAllNote()
	res, err := json.Marshal(notes)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateNoteById(w http.ResponseWriter, r *http.Request, u *models.User) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	temp := models.GetNoteById(int64(id))
	if u.UserType != "ADMIN" && int64(u.ID) != int64(temp.UserId) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	note := &models.Note{}
	utils.PareseBody(r, note)
	n, err := note.UpdateNoteById(int64(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, err := json.Marshal(n)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteNoteById(w http.ResponseWriter, r *http.Request, u *models.User) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	temp := models.GetNoteById(int64(id))
	if u.UserType != "ADMIN" && int64(u.ID) != int64(temp.UserId) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	note := models.DeleteNoteById(int64(id))
	res, err := json.Marshal(note)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
