package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"

	"github.com/technodeguy/real-estate/api/services"
)

type Server struct {
	db        *sql.DB
	router    *mux.Router
	s3Service *services.AwsS3Service
}

func NewServer() *Server {
	return &Server{}
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

	awsS3Service := services.NewAwsS3Service(
		os.Getenv("AWS_ACCESS_KEY_ID"),
		os.Getenv("AWS_SECRET_ACCESS_KEY"),
		os.Getenv("AWS_BUCKET_NAME"),
	)

	awsS3Service.Initialize()

	server.s3Service = awsS3Service

	log.Println("Services initialized successfully")
}

func (server *Server) RunServer(port string) {
	fmt.Println(fmt.Sprintf("Listening on port %s", port))
	log.Fatal(http.ListenAndServe(":"+port, server.router))
}
