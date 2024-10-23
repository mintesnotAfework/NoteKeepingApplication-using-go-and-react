package routes

import (
	"github.com/gorilla/mux"
	"github.com/mintesnotAfework/NoteKeepingApplication-using-go-and-react/controllers"
	"github.com/mintesnotAfework/NoteKeepingApplication-using-go-and-react/middleware"
)

func NoteRouteRegisteration(r *mux.Router) {
	router := r.PathPrefix("/note").Subrouter()

	router.HandleFunc("/", middleware.AdminAutheticateJWT(controllers.GetAllNote)).Methods("GET")
	router.HandleFunc("/{id}", middleware.UserAutheticateJWT(controllers.GetNoteById)).Methods("GET")
	router.HandleFunc("/{id}/user", middleware.UserAutheticateJWT(controllers.GetNoteByUserId)).Methods("GET")
	router.HandleFunc("/", middleware.UserAutheticateJWT(controllers.CreateNote)).Methods("POST")
	router.HandleFunc("/{id}", middleware.UserAutheticateJWT(controllers.UpdateNoteById)).Methods("PUT")
	router.HandleFunc("/{id}", middleware.UserAutheticateJWT(controllers.DeleteNoteById)).Methods("DELETE")
}
