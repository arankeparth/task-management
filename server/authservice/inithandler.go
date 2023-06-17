package authservice

import (
	"log"
	"net/http"
	authdl "plantrip-backend/server/authservice/dl"
	dbconn "plantrip-backend/server/db"
)

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (as *AuthService) InitAuthServiceHandler() http.Handler {
	AuthDB := "authorization"
	Db, err := dbconn.NewDB(AuthDB)
	if err != nil {
		print(err.Error())
		log.Fatal(err)
		return nil
	}
	AuthDl := authdl.NewAuthDl(Db)
	AuthBl := NewAuthHandler(AuthDl)
	AuthEps, err := NewCustomerEndpoints(AuthBl, &log.Logger{})
	if err != nil {
		log.Fatal("Failed to create authentication endpoints", err.Error())
		return nil
	}

	HttpHandler := MakeHandler(AuthEps, AuthBl)
	return HttpHandler
}
