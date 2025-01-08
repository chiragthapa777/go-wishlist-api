package service

import (
	"errors"

	"github.com/chiragthapa777/wishlist-api/database"
	"github.com/chiragthapa777/wishlist-api/model"
	authService "github.com/chiragthapa777/wishlist-api/service/auth"
	"gorm.io/gorm"
)

func CreateUser(user *model.User, tx *gorm.DB) error {
	db := tx

	if db == nil {
		db = database.GetDatabase()
	}

	found, _ := FindUserByEmail(user.Email, tx)

	if found != nil {
		return errors.New("Email already in use.")
	}

	hashedPassword, err := authService.HashPassword(user.Password)

	if err != nil {
		return err
	}

	user.Password = hashedPassword

	result := db.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func FindUserByEmail(email string, tx *gorm.DB) (*model.User, error) {
	db := tx

	if db == nil {
		db = database.GetDatabase()
	}

	user := model.User{}

	result := db.Where("email = ?", email).Take(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
