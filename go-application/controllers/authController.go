package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mintesnotAfework/NoteKeepingApplication-using-go-and-react/dto"
	"github.com/mintesnotAfework/NoteKeepingApplication-using-go-and-react/helpers"
	"github.com/mintesnotAfework/NoteKeepingApplication-using-go-and-react/models"
	"github.com/mintesnotAfework/NoteKeepingApplication-using-go-and-react/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := &models.User{}
	utils.PareseBody(r, user)
	if !helpers.ValidateUser(user) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user.UserType = "USER"
	u := user.CreateUser()
	if u.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hashedPassword, err := helpers.HashPassword(user.Password)
	if hashedPassword == "" || err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	refresh_token, token, err := helpers.GenerateToken(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(dto.Combiner(&refresh_token, &token))
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func AdminRegister(w http.ResponseWriter, r *http.Request, adminUser *models.User) {
	w.Header().Set("Content-Type", "application/json")

	user := &models.User{}
	utils.PareseBody(r, user)
	if !helpers.ValidateUser(user) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u := user.CreateUser()
	if u.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hashedPassword, err := helpers.HashPassword(user.Password)
	if hashedPassword == "" || err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	u.Password = "XXX-XXX-XXX"
	res, _ := json.Marshal(u)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := &models.User{}
	utils.PareseBody(r, user)
	if user.Email == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u := models.GetUserByEmail(user.Email)
	if u.ID == 0 {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if u.Password != user.Password {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	refresh, token, err := helpers.GenerateToken(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(dto.Combiner(&refresh, &token))

	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
