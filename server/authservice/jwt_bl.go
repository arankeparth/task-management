package authservice

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func (h *AuthHandler) VerifyJwt(tokenStr string) (jwt.MapClaims, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			// TODO: Get the key from the environment
			return []byte(os.Getenv("JWT_KEY")), nil
		})

	if err != nil {
		return nil, err
	}

	if !tkn.Valid {
		return nil, errors.New("invalid token")
	}

	return jwt.MapClaims{}, nil
}
