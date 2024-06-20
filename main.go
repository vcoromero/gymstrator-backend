package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vcoromero/gymstration-backend/config"
	"github.com/vcoromero/gymstration-backend/database"
	"github.com/vcoromero/gymstration-backend/src/handlers"
)

func main() {
	config.Init()
	database.Connect()

	r := mux.NewRouter()
	r.HandleFunc("/roles", handlers.GetRoles).Methods("GET")

	port := ":" + config.ServerPort
	log.Printf("Starting server on %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
