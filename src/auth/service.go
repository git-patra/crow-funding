package auth

import "github.com/golang-jwt/jwt/v4"

type Service interface {
	GenerateToken(userId int) (string, error)
}

type jwtService struct {
}

func NewJWTService() *jwtService {
	return &jwtService{}
}

var SecretKey = []byte("rahasia")

func (s *jwtService) GenerateToken(userId int) (string, error) {
	payload := jwt.MapClaims{}
	payload["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token.SignedString(SecretKey)
}
