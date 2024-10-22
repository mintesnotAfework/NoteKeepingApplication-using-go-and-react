package routes

import (
	"github.com/gorilla/mux"
	"github.com/mintesnotAfework/NoteKeepingApplication-using-go-and-react/controllers"
)

func NoteRouteRegisteration(r *mux.Router) {
	router := r.PathPrefix("/note").Subrouter()

	router.HandleFunc("/", controllers.GetAllNote).Methods("GET")
	router.HandleFunc("/{id}", controllers.GetNoteById).Methods("GET")
	router.HandleFunc("/{id}/user", controllers.GetNoteByUserId).Methods("GET")
	router.HandleFunc("/", controllers.CreateNote).Methods("POST")
	router.HandleFunc("/{id}", controllers.UpdateNoteById).Methods("PUT")
	router.HandleFunc("/{id}", controllers.DeleteNoteById).Methods("DELETE")
}
