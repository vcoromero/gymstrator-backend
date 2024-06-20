package main

import (
	"github.com/vcoromero/gymstration-backend/config"
	"github.com/vcoromero/gymstration-backend/database"
)

func main() {
	config.Init()
	database.Connect()
}
