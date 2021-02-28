package config

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

// JwtService : JWT Service Contract
type JwtService interface {
	GenerateJwtToken(userID int) (string, error)
}

type jwtService struct {
}

// JwtServiceInit Initiation
func JwtServiceInit() *jwtService {
	return &jwtService{}
}

// GenerateJwtToken : Service Generate JWT Token
func (s *jwtService) GenerateJwtToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	secretKey := os.Getenv("SECRET_JWT")
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claim)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}
