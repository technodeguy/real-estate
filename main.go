package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/technodeguy/real-estate/models"
)

const (
	routePrefix = "/api/v1"
)

func handler(w http.ResponseWriter, request *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func initServer() {
	r := mux.NewRouter()

	r.HandleFunc(routePrefix, handler).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", r))
}

func main() {
	models.GetDBConn()
	initServer()
	log.Println("Test application")
}
