package routes

import (
	"github.com/gorilla/mux"
	"github.com/mintesnotAfework/NoteKeepingApplication-using-go-and-react/controllers"
)

func UserRouteRegisteration(r *mux.Router) {
	router := r.PathPrefix("/user").Subrouter()

	router.HandleFunc("/", controllers.GetAllUser).Methods("GET")
	router.HandleFunc("/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/{id}", controllers.UpdateUserById).Methods("PUT")
	router.HandleFunc("/{id}", controllers.DeleteUserById).Methods("DELETE")
}
