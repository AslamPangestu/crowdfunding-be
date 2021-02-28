package config

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

// JwtService : JWT Service Contract
type JwtService interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

// JwtServiceInit Initiation
func JwtServiceInit() *jwtService {
	return &jwtService{}
}

var secretKey = os.Getenv("SECRET_JWT")

// GenerateToken : Service Generate JWT Token
func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claim)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

// ValidateToken : Service Validate JWT Token
func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid token")
		}
		return secretKey, nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
