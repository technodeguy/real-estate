package controllers

import (
	"fmt"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, request *http.Request) {
	fmt.Fprint(w, "Hello world")
}
