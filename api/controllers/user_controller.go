package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/technodeguy/real-estate/api/consts"
	"github.com/technodeguy/real-estate/api/validators"

	"github.com/technodeguy/real-estate/api/models"
	"github.com/technodeguy/real-estate/api/responses"
)

func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	users, _ := user.FindAllUsers(server.db)

	responses.Json(w, http.StatusOK, users)
}

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	userInput := &validators.CreateUserRequest{}

	if err := validators.DecodeAndValidate(r, userInput); err != nil {
		responses.Error(w, http.StatusBadRequest, errors.New(consts.BAD_REQUEST))
		return
	}

	user := models.User{
		Nickname:    userInput.Nickname,
		Password:    userInput.Password,
		PhoneNumber: userInput.PhoneNumber,
	}

	id, err := user.CreateUser(server.db)

	log.Print("Error occured", err)

	responses.Json(w, http.StatusCreated, map[string]uint32{"id": id})
}
