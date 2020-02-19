package responses

import (
	"net/http"
)

func Error(w http.ResponseWriter, statusCode int, err error) {
	Json(w, statusCode, struct {
		ErrorCode string `json:"errorCode"`
	}{
		ErrorCode: err.Error(),
	})
}
