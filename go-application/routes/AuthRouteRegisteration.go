package routes

import (
	"github.com/gorilla/mux"
	"github.com/mintesnotAfework/NoteKeepingApplication-using-go-and-react/controllers"
)

func AuthRouteRegisteration(r *mux.Router) {
	router := r.PathPrefix("/auth").Subrouter()

	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/new/token/refresh", nil).Methods("POST")
}
