package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mintesnotAfework/NoteKeepingApplication-using-go-and-react/models"
	"github.com/mintesnotAfework/NoteKeepingApplication-using-go-and-react/utils"
)

func GetUserById(w http.ResponseWriter, r *http.Request, u *models.User) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := models.GetUserById(int64(id))
	res, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllUser(w http.ResponseWriter, r *http.Request, u *models.User) {
	w.Header().Set("Content-Type", "application/json")

	users := models.GetAllUsers()
	res, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUserById(w http.ResponseWriter, r *http.Request, u *models.User) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := &models.User{}
	utils.PareseBody(r, user)
	result, err := user.UpdateUserById(int64(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUserById(w http.ResponseWriter, r *http.Request, u *models.User) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := models.DeleteUserById(int64(id))
	res, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
