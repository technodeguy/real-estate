package controllers

import (
	"net/http"

	"github.com/technodeguy/real-estate/api/models"
	"github.com/technodeguy/real-estate/api/responses"
)

func (server *Server) GetUsers(w http.ResponseWriter, request *http.Request) {
	user := models.User{}

	users, _ := user.FindAllUsers(server.db)

	responses.Json(w, http.StatusOK, users)
}
