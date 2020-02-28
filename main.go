package main

import (
	"github.com/joho/godotenv"
	"github.com/technodeguy/real-estate/api"
)

func init() {
	godotenv.Load()
}

func main() {
	api.Start()
}
