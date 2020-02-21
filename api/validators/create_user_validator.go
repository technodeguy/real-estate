package validators

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateUserRequest struct {
	Nickname    string `json:"nickname"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
}

func (u CreateUserRequest) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Nickname, validation.Required, validation.Length(5, 45)),
		validation.Field(&u.Password, validation.Required, validation.Length(6, 60)),
		validation.Field(&u.PhoneNumber, validation.Required, validation.Length(10, 15)),
	)
}
