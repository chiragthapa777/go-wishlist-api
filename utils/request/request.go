package utils

import (
	"errors"
	"log"
	"strings"

	"github.com/chiragthapa777/wishlist-api/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetUser() {

}

// ParseAndValidate parses the request body into the given struct and validates it using the provided validator.
func ParseAndValidate[T any](c *fiber.Ctx) (*T, error) {
	// Parse the body
	var body T
	if err := c.BodyParser(&body); err != nil {
		return nil, err
	}

	validate := validator.New()
	err := validate.Struct(body)
	if err != nil {
		// Validation failed, handle the error
		validationErrors := err.(validator.ValidationErrors)
		errorString := make([]string, len(validationErrors))
		for _, err := range validationErrors {
			errorString = append(errorString, err.Error())
		}
		log.Println(errorString)
		formattedSting := strings.Join(utils.Filter(errorString, func(s string) bool { return s != "" }), ", ") + "."
		return nil, errors.New("Validation Error : " + formattedSting)
	}
	return &body, nil
}
