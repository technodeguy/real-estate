package validators

import (
	"encoding/json"
	"net/http"
)

type InputValidation interface {
	Validate() error
}

func DecodeAndValidate(r *http.Request, v InputValidation) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}

	return v.Validate()
}
