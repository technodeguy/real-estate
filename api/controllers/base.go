package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	db     *sql.DB
	router *mux.Router
}

func (server *Server) Initialize(DbUri string) {
	var err error
	server.db, err = sql.Open("mysql", DbUri)

	if err != nil {
		log.Fatalf("Unable to connect to db %v", err.Error())
	}

	if err = server.db.Ping(); err != nil {
		log.Fatalf("Unable to ping to db %v", err.Error())
	}

	log.Println("DB connected successfully")

	server.router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) RunServer(port string) {
	fmt.Println(fmt.Sprintf("Listening on port %s", port))
	log.Fatal(http.ListenAndServe(":"+port, server.router))
}
