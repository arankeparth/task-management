package authservice

import (
	"context"
	authdl "plantrip-backend/server/authservice/dl"
	"plantrip-backend/server/spec/authspec"

	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	AuthDl *authdl.AuthDl
}

func NewAuthHandler(authDl *authdl.AuthDl) *AuthHandler {
	authHandler := &AuthHandler{
		AuthDl: authDl,
	}
	return authHandler
}

func (h *AuthHandler) Login(ctx context.Context, username string, password string) (*authspec.LoginResponse, error) {
	encryptedPass, customerId := h.AuthDl.GetInfo(username)
	isLoginSuccessful := bcrypt.CompareHashAndPassword([]byte(encryptedPass), []byte(password)) == nil
	if isLoginSuccessful {
		privateKeyString, publicKey, err := h.CreatePrivatePublicKeyPair(ctx, customerId)
		if err != nil {
			return &authspec.LoginResponse{}, err
		}

		tokenString, err := getJwtToken(privateKeyString, customerId)
		if err != nil {
			return &authspec.LoginResponse{
				IsLoggedIn:   false,
				ErrorMessage: err.Error(),
			}, err
		}
		return &authspec.LoginResponse{
			IsLoggedIn: true,
			PublicKey:  publicKey,
			AuthToken:  tokenString,
			CustomerId: customerId,
		}, nil
	}
	return &authspec.LoginResponse{
		IsLoggedIn:   false,
		ErrorMessage: "Failed to login",
	}, nil
}

func (h *AuthHandler) CreateUser(ctx context.Context, username string, password string, customerid int64) error {
	err := h.AuthDl.CreateUser(username, password, customerid)
	if err != nil {
		return err
	}
	return nil
}
