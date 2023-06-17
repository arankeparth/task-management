package authservice

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func FormatPubKey(pubKeyString string) string {
	return "-----BEGIN PUBLIC KEY-----\n" + pubKeyString + "\n-----END PUBLIC KEY-----\n"
}

func encodePrivateKeyToPem(privateKey *rsa.PrivateKey) []byte {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPem := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	return pem.EncodeToMemory(privateKeyPem)
}

func encodePublicKeyToPem(publicKey *rsa.PublicKey) []byte {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		fmt.Println("Failed to encode public key:", err)
		return nil
	}

	publicKeyPem := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	return pem.EncodeToMemory(publicKeyPem)
}

func getJwtToken(privateKeyString string, customerId int64) (string, error) {
	block, _ := pem.Decode([]byte(privateKeyString))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		fmt.Println("Failed to decode PEM block containing private key")
		return "", errors.New("internal error")
	}

	// Parse the private key
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("Failed to parse private key:", err)
		return "", errors.New("internal error")
	}
	claims := jwt.MapClaims{
		"customerid": fmt.Sprint(customerId),
		"iat":        time.Now().Unix(),
		"exp":        time.Now().Add(time.Hour * 2).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
