package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mintesnotAfework/NoteKeepingApplication-using-go-and-react/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("can not fine the .env file")
	}

	r := mux.NewRouter()
	router := r.PathPrefix("/api").Subrouter()
	routes.AuthRouteRegisteration(router)
	routes.UserRouteRegisteration(router)
	routes.NoteRouteRegisteration(router)

	port := os.Getenv("PORT")
	hostname := os.Getenv("HOSTNAME")

	fmt.Printf("Starting port at %s:%s", hostname, port)
	fmt.Println("")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(hostname+":"+port, r))
}
