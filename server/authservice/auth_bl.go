package authservice

import (
	"context"
	"os"
	authdl "task-management/server/authservice/dl"
	"task-management/server/spec/authspec"
	"time"
	"fmt"

	"github.com/dgrijalva/jwt-go"
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

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (h *AuthHandler) Login(ctx context.Context, username string, password string) (*authspec.LoginResponse, error) {
	encryptedPass, userID := h.AuthDl.GetInfo(username)
	isLoginSuccessful := bcrypt.CompareHashAndPassword([]byte(encryptedPass), []byte(password)) == nil
	if isLoginSuccessful {
		expirationTime := time.Now().Add(time.Hour * 1)

		claims := &Claims{
			Username: username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
		if err != nil {
			return nil, err
		}

		return &authspec.LoginResponse{
			IsLoggedIn: true,
			AuthToken:  tokenString,
			CustomerId: userID,
		}, nil
	}
	return &authspec.LoginResponse{
		IsLoggedIn:   false,
		ErrorMessage: "Failed to login",
	}, nil
}

func (h *AuthHandler) CreateUser(ctx context.Context, username string, password string, customerid string) error {
	fmt.Println("getting here")
	err := h.AuthDl.CreateUser(username, password, customerid)
	if err != nil {
		return err
	}
	return nil
}
