package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/technodeguy/real-estate/api/controllers"
)

var server = controllers.Server{}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Env file loading error")
	} else {
		log.Print("Env variables successfully loaded")
	}
}

func Start() {
	server.Initialize(os.Getenv("DB_URI"))

	server.RunServer(os.Getenv("API_PORT"))
}
