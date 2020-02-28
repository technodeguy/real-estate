package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/technodeguy/real-estate/api/consts"
	"github.com/technodeguy/real-estate/api/validators"

	"github.com/technodeguy/real-estate/api/models"
	"github.com/technodeguy/real-estate/api/responses"
	"github.com/technodeguy/real-estate/api/utils"
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

	if driverErr, ok := err.(*mysql.MySQLError); ok {
		log.Print("Error occured", err)
		if driverErr.Number == 1062 { // ER_DUP_ENTRY
			responses.Error(w, http.StatusUnprocessableEntity, errors.New(consts.USER_ALREADY_EXISTS))
		} else {
			responses.Error(w, http.StatusInternalServerError, errors.New(consts.INTERNAL))
		}

		return
	}

	responses.Json(w, http.StatusCreated, map[string]uint32{"id": id})
}

func (server *Server) GetPresignedUrl(w http.ResponseWriter, r *http.Request) {
	userInput := &validators.CreateSignedUrlRequest{}

	if err := validators.DecodeAndValidate(r, userInput); err != nil {
		responses.Error(w, http.StatusBadRequest, errors.New(consts.BAD_REQUEST))
		return
	}

	mimeType := utils.GetFileMimeType(userInput.Filename)

	log.Println("GENERATED", mimeType, " and ", userInput.Filename)

	data, err := server.s3Service.GetPresignedUrl(mimeType, userInput.Filename)

	if err != nil {
		log.Println("Error by creating presigned url", err)

		responses.Error(w, http.StatusUnprocessableEntity, errors.New(consts.GET_PRESIGNED_URL_ERROR))

		return
	}

	responses.Json(w, http.StatusCreated, data)
}

// func (server *Server) SaveUserAvatar(w http.ResponseWriter, r *http.Request) {

// }
