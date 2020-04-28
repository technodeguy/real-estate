package validators

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
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

type IdRequest struct {
	Id int `json:"id"`
}

func (e IdRequest) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Id, validation.Required),
	)
}
