package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	cfg "github.com/vcoromero/gymstration-backend/config"
)

var DB *sql.DB

func Connect() {
	databaseConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)
	databaseConnectionUri := databaseConnectionString

	db, err := sql.Open("mysql", databaseConnectionUri)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	DB = db
	log.Println("Connected to the database!")
}
