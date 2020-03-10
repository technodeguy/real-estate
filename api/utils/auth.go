package utils

import (
	"net/http"
	"strconv"
)

func ExtractIdFromHeaders(r *http.Request) int {
	id, _ := strconv.Atoi(r.Header.Get("authorization"))

	return id
}
