package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/technodeguy/real-estate/api/controllers"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("Running on staging env")
	} else {
		log.Print("Running on dev env without docker")
	}
}

func Start() {
	server := controllers.NewServer()

	server.Initialize(os.Getenv("DB_URI"))

	server.RunServer(os.Getenv("API_PORT"))
}
