package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var dbConn *sql.DB

func init() {

}

func GetDBConn() *sql.DB {
	return dbConn
}
