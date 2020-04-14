package validators

import validation "github.com/go-ozzo/ozzo-validation/v4"

type CreateEstateRequest struct {
	Title  string `json:"title"`
	Price  int    `json:"price"`
	IsRent bool   `json:"isRent"`
	CityId int    `json:"cityId"`
}

func (e CreateEstateRequest) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Title, validation.Required, validation.Length(3, 128)),
		validation.Field(&e.Price, validation.Required),
		validation.Field(&e.IsRent, validation.Required),
		validation.Field(&e.CityId, validation.Required),
	)
}
