package api

import (
	"log"

	"github.com/technodeguy/real-estate/api/config"
	"github.com/technodeguy/real-estate/api/controllers"
)

func Start() {
	cnf, err := config.LoadConfig("config")

	if err != nil {
		log.Fatalf("Unable to initialize config %s", err)
	}

	log.Printf("CONFIG %#v", *cnf)

	server := controllers.NewServer(cnf)

	server.Initialize()

	server.RunServer()
}
