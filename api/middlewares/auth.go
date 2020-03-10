package middlewares

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/technodeguy/real-estate/api/consts"
	"github.com/technodeguy/real-estate/api/responses"
	"github.com/technodeguy/real-estate/api/services"
)

func SetMiddlewareAuth(pTokenService services.ITokenService, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			log.Print("Missed auth header")
			responses.Error(w, http.StatusUnauthorized, errors.New(consts.UNAUTHORIZED))
			return
		}

		data, err := pTokenService.Decode(token)

		if err != nil {
			log.Print("Error by decoding", err)

			responses.Error(w, http.StatusUnauthorized, errors.New(consts.UNAUTHORIZED))
			return
		}

		originalToken, err := pTokenService.Get("AccessToken", strconv.Itoa(data.Id))

		if err != nil {
			log.Print("Error by retrieving original", err)

			responses.Error(w, http.StatusUnauthorized, errors.New(consts.UNAUTHORIZED))
			return
		}

		if originalToken != token {
			responses.Error(w, http.StatusUnauthorized, errors.New(consts.UNAUTHORIZED))
			return
		}

		r.Header.Set("authorization", strconv.Itoa(data.Id))
		next(w, r)
	}
}
