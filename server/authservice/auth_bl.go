package authservice

import (
	"context"
	"log"
	"os"
	authdl "task-management/server/authservice/dl"
	"task-management/server/spec/authspec"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	AuthDl *authdl.AuthDl
	logger *log.Logger
}

func NewAuthHandler(authDl *authdl.AuthDl) *AuthHandler {
	logger := log.New(os.Stdout, "[AuthService] ", log.LstdFlags)
	authHandler := &AuthHandler{
		AuthDl: authDl,
		logger: logger,
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
		h.logger.Printf("User logged in successfully for email-id: %s", username)
		return &authspec.LoginResponse{
			IsLoggedIn: true,
			AuthToken:  tokenString,
			CustomerId: userID,
		}, nil
	}
	h.logger.Printf("Failed to login for email-id: %s. Invalid creds", username)
	return &authspec.LoginResponse{
		IsLoggedIn:   false,
		ErrorMessage: "Failed to login",
	}, nil
}

func (h *AuthHandler) CreateUser(ctx context.Context, username string, password string, customerid string) error {
	err := h.AuthDl.CreateUser(username, password, customerid)
	if err != nil {
		h.logger.Printf("Error creating user creds: %s", err.Error())
		return err
	}
	h.logger.Printf("User creds created successfully for email-id: %s", username)
	return nil
}
