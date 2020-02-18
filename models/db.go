package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/technodeguy/real-estate/config"
)

var dbConn *sql.DB

func init() {
	var err error
	dbConn, err = sql.Open("mysql", config.DbURI)

	if err != nil {
		log.Fatalf("Unable to connect to db %v", err.Error())
	}

	if err = dbConn.Ping(); err != nil {
		log.Fatalf("Unable to ping to db %v", err.Error())
	}

	log.Println("DB connected successfully")
}

func GetDBConn() *sql.DB {
	return dbConn
}
