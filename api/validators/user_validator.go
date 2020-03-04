package validators

import (
	"errors"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateSignedUrlRequest struct {
	Filename string `json:"filename"`
}

func (u CreateSignedUrlRequest) Validate() error {
	var err error

	err = validation.ValidateStruct(&u,
		validation.Field(&u.Filename, validation.Required, validation.Length(1, 64)),
	)

	if !strings.Contains(u.Filename, ".") {
		err = errors.New("")
	}

	return err
}

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

type SaveUserAvatar struct {
	Id     int    `json:"id"`
	Avatar string `json:"avatar"`
}

func (u SaveUserAvatar) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Id, validation.Required),
		validation.Field(&u.Avatar, validation.Required, validation.Length(6, 255)),
	)
}

type LoginRequest struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

func (u LoginRequest) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Nickname, validation.Required),
		validation.Field(&u.Password, validation.Required),
	)
}
