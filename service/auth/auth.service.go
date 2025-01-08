package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/chiragthapa777/wishlist-api/config"
	"github.com/chiragthapa777/wishlist-api/model"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the password using bcrypt.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPassword verifies if the hashed password matches the plain-text password.
func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Struct to hold custom claims for the JWT payload
type Claims struct {
	UserId uint `json:"userId"`
	jwt.StandardClaims
}

// GenerateJWT generates a new JWT token with the specified user data and expiration time
func GenerateJWT(userID uint, expirationTime time.Duration) (string, error) {
	// Set up the claims
	claims := Claims{
		UserId: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expirationTime).Unix(),
			Issuer:    "go-whish-list",
		},
	}

	// Create a new JWT token with the claims and sign it with the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := config.GetConfig().JwtSecret // string
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetUserJwtToken(user *model.User) (string, int, error) {
	JwtAccessTokenExpiryMinutes, err := strconv.Atoi(config.GetConfig().JwtAccessTokenExpiryMinute)
	if err != nil {
		return "", 0, err
	}

	jwtToken, err := GenerateJWT(user.ID, time.Duration(JwtAccessTokenExpiryMinutes)*time.Minute)
	if err != nil {
		return "", 0, err
	}

	return jwtToken, JwtAccessTokenExpiryMinutes, nil
}

// ValidateJWT validates the JWT token and returns the decoded claims
func ValidateJWT(tokenString string) (*Claims, error) {
	secretKey := []byte(config.GetConfig().JwtSecret)
	// Parse the token
	token, err := jwt.ParseWithClaims((tokenString), &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure that the signing method matches
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	// If the token is invalid or expired
	if err != nil {
		return nil, err
	}

	// Validate the claims and return the claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
