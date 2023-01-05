package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

type Service interface {
	GenerateToken(userId int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
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

func (s *jwtService) ValidateToken(encodeToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodeToken, func(token *jwt.Token) (any, error) {
		// checkin has method sama atau engga
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid token!")
		}

		// kalo sama kasihin secret key buat di validasi sama method jwt.Parse()
		return SecretKey, nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
