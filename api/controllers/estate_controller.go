package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/technodeguy/real-estate/api/consts"
	"github.com/technodeguy/real-estate/api/models"
	"github.com/technodeguy/real-estate/api/responses"
	"github.com/technodeguy/real-estate/api/utils"
	"github.com/technodeguy/real-estate/api/validators"
)

func (server *Server) CreateEstate(w http.ResponseWriter, r *http.Request) {
	userId := utils.ExtractIdFromHeaders(r)

	userInput := &validators.CreateEstateRequest{}

	if err := validators.DecodeAndValidate(r, userInput); err != nil {
		responses.Error(w, http.StatusBadRequest, errors.New(consts.BAD_REQUEST))
		return
	}

	estate := models.Estate{
		Title:  userInput.Title,
		Price:  userInput.Price,
		IsRent: userInput.IsRent,
		CityId: userInput.CityId,
		UserId: userId,
	}

	id, err := estate.CreateEstate(server.db)

	if driverErr, ok := err.(*mysql.MySQLError); ok {
		log.Print("Error occured", err)
		if driverErr.Number == 1452 { // Constraint failes
			responses.Error(w, http.StatusUnprocessableEntity, errors.New(consts.CONSTRAINT_NOT_FOUND))
		} else {
			responses.Error(w, http.StatusInternalServerError, errors.New(consts.INTERNAL))
		}

		return
	}

	responses.Json(w, http.StatusCreated, map[string]uint32{"id": id})
}
