package controllers

import (
	"net/http"

	"github.com/technodeguy/real-estate/api/responses"
)

func (server *Server) GetUsers(w http.ResponseWriter, request *http.Request) {
	responses.Json(w, http.StatusOK, map[string]string{"usersCount": "11"})
}
