package controller

import (
	"database/sql"
	"errors"

	"github.com/chiragthapa777/wishlist-api/dto"
	"github.com/chiragthapa777/wishlist-api/model"
	authService "github.com/chiragthapa777/wishlist-api/service/auth"
	userService "github.com/chiragthapa777/wishlist-api/service/user"
	requestUtils "github.com/chiragthapa777/wishlist-api/utils/request"
	responseUtils "github.com/chiragthapa777/wishlist-api/utils/response"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {

	body, err := requestUtils.ParseAndValidate[dto.RegisterUserDto](c)

	if err != nil {
		return responseUtils.SendErrorResponse(c, responseUtils.ResponseParam{Error: err, Status: 400})
	}

	user := model.User{
		Email:    body.Email,
		Password: body.Password,
		Name: sql.NullString{
			String: body.Name.Data,
			Valid:  body.Name.Valid,
		},
	}

	err = userService.CreateUser(&user, nil)
	if err != nil {
		return responseUtils.SendErrorResponse(c, responseUtils.ResponseParam{Error: err, Status: 400})
	}

	jwtToken, expiryMinutes, err := authService.GetUserJwtToken(&user)
	if err != nil {
		return responseUtils.SendErrorResponse(c, responseUtils.ResponseParam{Error: err, Status: 500})
	}

	return responseUtils.SendResponse(c, responseUtils.ResponseParam{
		Message: "Registration successful!",
		Data:    map[string]any{"user": user, "token": jwtToken, "expiryMinutes": expiryMinutes},
	})
}

func Login(c *fiber.Ctx) error {
	body, err := requestUtils.ParseAndValidate[dto.LoginDto](c)

	if err != nil {
		return responseUtils.SendErrorResponse(c, responseUtils.ResponseParam{Error: err, Status: 400})
	}

	user, _ := userService.FindUserByEmail(body.Email, nil)

	if user == nil {
		return responseUtils.SendErrorResponse(c, responseUtils.ResponseParam{Error: errors.New("invalid credentials"), Status: 400})
	}

	err = authService.CheckPassword(user.Password, body.Password)
	if err != nil {
		return responseUtils.SendErrorResponse(c, responseUtils.ResponseParam{Error: errors.New("invalid credentials"), Status: 400})

	}

	jwtToken, expiryMinutes, err := authService.GetUserJwtToken(user)
	if err != nil {
		return responseUtils.SendErrorResponse(c, responseUtils.ResponseParam{Error: err, Status: 500})
	}

	return responseUtils.SendResponse(c, responseUtils.ResponseParam{
		Message: "Login successful!",
		Data:    map[string]any{"user": user, "expiryMinutes": expiryMinutes, "token": jwtToken},
	})
}

func GetCurrentUser(c *fiber.Ctx) error {
	user, err := requestUtils.GetCurrentUser(c)
	if err != nil {
		return responseUtils.SendErrorResponse(c, responseUtils.ResponseParam{Error: err, Status: 400})
	}
	return responseUtils.SendResponse(c, responseUtils.ResponseParam{
		Message: "Current user detail",
		Data:    user,
	})
}
