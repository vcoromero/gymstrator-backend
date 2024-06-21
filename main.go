package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vcoromero/gymstration-backend/config"
	"github.com/vcoromero/gymstration-backend/database"
	"github.com/vcoromero/gymstration-backend/src/routes"
)

func main() {
	config.Init()
	database.Connect()

	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	port := ":" + config.ServerPort
	log.Printf("Starting server on %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
