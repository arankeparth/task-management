package authservice

import (
	"net/http"
	"task-management/server/spec/authspec"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandler(eps *AuthEps, jwtHandler *AuthHandler) http.Handler {

	loginHandler := kithttp.NewServer(
		eps.LoginEP,
		decodeLoginRequest,
		createSessionCookie,
	)

	createUserHandler := kithttp.NewServer(
		eps.CreateUserEP,
		decodeCreateUserRequest,
		encodeResponse,
	)

	r := mux.NewRouter()
	r.Handle(authspec.LoginPath, loginHandler).Methods("POST")
	r.Handle(authspec.CreateUserPath, createUserHandler).Methods("POST")

	return r
}
