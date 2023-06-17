package authservice

import (
	"context"
	"errors"
	"fmt"
	"log"
	"plantrip-backend/server/spec/authspec"

	"github.com/go-kit/kit/endpoint"
)

type AuthEps struct {
	LoginEP      endpoint.Endpoint
	CreateUserEP endpoint.Endpoint
	VerifyJwtEP  endpoint.Endpoint
}

func NewCustomerEndpoints(h *AuthHandler, logger *log.Logger) (*AuthEps, error) {
	loginEP := makeLoginEP(h)
	if loginEP == nil {
		log.Printf("Failed to create login endpoint method=%s", "NewCustomerEndpoints")
		return nil, errors.New("failed to create Login Endpoint")
	}

	createUserEP := makeCreateUserEP(h)
	if createUserEP == nil {
		log.Printf("Failed to create createUser endpoint method=%s", "NewCustomerEndpoints")
		return nil, errors.New("failed to create createUser endpoint")
	}

	verifyJwtEP := makeVerifyJwtEP(h)
	if verifyJwtEP == nil {
		log.Printf("Failed to create verifyJwt endpoint method=%s", "NewCustomerEndpoints")
		return nil, errors.New("failed to create verifyJwt endpoint")
	}
	return &AuthEps{
		LoginEP:      loginEP,
		CreateUserEP: createUserEP,
		VerifyJwtEP:  verifyJwtEP,
	}, nil
}

func makeLoginEP(h *AuthHandler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*authspec.LoginRequest)
		if !ok {
			return nil, errors.New("error while converting types")
		}
		resp, err := h.Login(ctx, req.Username, req.Password)
		return resp, err
	}
}

func makeCreateUserEP(h *AuthHandler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*authspec.CreateUserRequest)
		if !ok {
			return nil, errors.New("error while converting types")
		}
		err = h.CreateUser(ctx, req.Username, req.Password, req.CustomerId)
		if err != nil {
			return &authspec.CreateUserResponse{
				Status:       1,
				ErrorMessage: err.Error(),
			}, err
		}

		return &authspec.CreateUserResponse{
			Status: 0,
		}, nil
	}
}

func makeVerifyJwtEP(h *AuthHandler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*authspec.VerifyJwtReq)
		if !ok {
			return nil, errors.New("error while converting types")
		}
		parsedToken, err := h.VerifyJwt(req.TokenString, req.PublicKeyString)
		if err != nil {
			errMsg := fmt.Sprintf("invalid jwt token: %s", err.Error())
			return nil, errors.New(errMsg)
		}

		return &authspec.VerifyJwtResponse{
			Status:      0,
			ParsedToken: parsedToken,
		}, nil
	}
}
