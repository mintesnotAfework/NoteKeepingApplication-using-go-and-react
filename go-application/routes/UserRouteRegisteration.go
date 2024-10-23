package routes

import (
	"github.com/gorilla/mux"
	"github.com/mintesnotAfework/NoteKeepingApplication-using-go-and-react/controllers"
	"github.com/mintesnotAfework/NoteKeepingApplication-using-go-and-react/middleware"
)

func UserRouteRegisteration(r *mux.Router) {
	router := r.PathPrefix("/user").Subrouter()

	router.HandleFunc("/", middleware.AdminAutheticateJWT(controllers.GetAllUser)).Methods("GET")
	router.HandleFunc("/{id}", middleware.UserAutheticateJWT(controllers.GetUserById)).Methods("GET")
	router.HandleFunc("/{id}", middleware.UserAutheticateJWT(controllers.UpdateUserById)).Methods("PUT")
	router.HandleFunc("/{id}", middleware.UserAutheticateJWT(controllers.DeleteUserById)).Methods("DELETE")
}
