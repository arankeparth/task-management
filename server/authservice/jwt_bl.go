package authservice

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
)

func (h *AuthHandler) VerifyJwt(tokenString string, pubKeyString string) (jwt.MapClaims, error) {
	fmt.Println(pubKeyString)
	fmt.Println("hi")
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(FormatPubKey(pubKeyString)))
	fmt.Println(pubKeyString)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		log.Printf("Invalid jwt token: %v", err)
		fmt.Println(err.Error())
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	}
	fmt.Println("invalid token")
	return nil, errors.New("invalid token claims")
}

func (h *AuthHandler) CreatePrivatePublicKeyPair(ctx context.Context, customerId int64) (string, string, error) {

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Failed to generate private key:", err)
		return "", "", nil
	}
	publicKey := privateKey.PublicKey

	privateKeyPem := encodePrivateKeyToPem(privateKey)
	publicKeyPem := encodePublicKeyToPem(&publicKey)
	return string(privateKeyPem), string(publicKeyPem), nil
}
