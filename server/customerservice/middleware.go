package customerApi

import (
	"context"
	"errors"
	"log"
	"task-management/server/spec/customerspec"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/endpoint"
)

func JWTMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			req := request.(customerspec.GetOffersRequest)
			_, err := verifyJwt(req.TokenString, req.PublicKey)
			if err != nil {
				return nil, err
			}
			return next(ctx, request)
		}
	}
}

func verifyJwt(tokenString string, pubKeyString string) (jwt.MapClaims, error) {
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(FormatPubKey(pubKeyString)))
	if err != nil {
		return nil, err
	}

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		log.Printf("Invalid jwt token: %v", err)
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token claims")
}

func FormatPubKey(pubKeyString string) string {
	return "-----BEGIN PUBLIC KEY-----\n" + pubKeyString + "\n-----END PUBLIC KEY-----\n"
}
