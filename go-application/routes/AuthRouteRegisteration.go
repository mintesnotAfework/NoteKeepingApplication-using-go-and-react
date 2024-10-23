package routes

import (
	"github.com/gorilla/mux"
	"github.com/mintesnotAfework/NoteKeepingApplication-using-go-and-react/controllers"
	"github.com/mintesnotAfework/NoteKeepingApplication-using-go-and-react/middleware"
)

func AuthRouteRegisteration(r *mux.Router) {
	router := r.PathPrefix("/admin/auth").Subrouter()

	router.HandleFunc("/register", middleware.AdminAutheticateJWT(controllers.AdminRegister)).Methods("POST")
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/new/token/refresh", nil).Methods("POST")
}
