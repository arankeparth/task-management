package authservice

import (
	"net/http"
	"plantrip-backend/server/spec/authspec"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandler(eps *AuthEps, jwtHandler *AuthHandler) http.Handler {

	loginHandler := kithttp.NewServer(
		eps.LoginEP,
		decodeLoginRequest,
		encodeResponse,
	)

	createUserHandler := kithttp.NewServer(
		eps.CreateUserEP,
		decodeCreateUserRequest,
		encodeResponse,
	)

	verifyJwtHandler := kithttp.NewServer(
		eps.VerifyJwtEP,
		decodeValidateJwtRequest,
		encodeResponse,
	)

	r := mux.NewRouter()
	r.Handle(authspec.LoginPath, loginHandler).Methods("POST")
	r.Handle(authspec.CreateUserPath, createUserHandler).Methods("POST")
	r.Handle(authspec.VerifyJwtPath, verifyJwtHandler).Methods("POST")

	return r
}
