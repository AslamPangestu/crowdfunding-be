package config

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

// AuthService : JWT Service Contract
type AuthService interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type authService struct {
}

// NewAuthService Initiation
func NewAuthService() *authService {
	return &authService{}
}

var SECRET_KEY = []byte(os.Getenv("SECRET_JWT"))

// GenerateToken : Service Generate JWT Token
func (s *authService) GenerateToken(userID int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID

	signedToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

// ValidateToken : Service Validate JWT Token
func (s *authService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid token")
		}
		return SECRET_KEY, nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
