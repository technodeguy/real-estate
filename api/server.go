package api

import (
	"os"

	"github.com/technodeguy/real-estate/api/controllers"
)

func Start() {
	server := controllers.NewServer()

	server.Initialize(os.Getenv("DB_URI"))

	server.RunServer(os.Getenv("API_PORT"))
}
