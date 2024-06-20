package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	cfg "github.com/vcoromero/gymstration-backend/config"
)

func Connect() {
	databaseConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)
	databaseConnectionUri := databaseConnectionString

	db, err := sql.Open("mysql", databaseConnectionUri)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the database!")
}
