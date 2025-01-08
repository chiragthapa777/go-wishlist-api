package dto

import utils "github.com/chiragthapa777/wishlist-api/utils/json/optional-nullable-property"

type RegisterUserDto struct {
	Name utils.OptionalNullableProperty[string] `json:"name" validate:""`
	LoginDto
}

type LoginDto struct {
	Password string `json:"password" validate:"required,min=5,max=20"`
	Email    string `json:"email" validate:"required,email"`
}
