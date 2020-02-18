package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DbURI string
)

func init() {
	e := godotenv.Load()

	if e != nil {
		log.Fatalf("Env error %s", e.Error())
	}

	DbURI = os.Getenv("DB_URI")

	log.Print("Initialized env variables")
}
