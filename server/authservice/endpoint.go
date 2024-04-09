package authservice

import (
	"context"
	"errors"
	"log"
	"task-management/server/spec/authspec"

	"github.com/go-kit/kit/endpoint"
)

type AuthEps struct {
	LoginEP      endpoint.Endpoint
	CreateUserEP endpoint.Endpoint
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

	return &AuthEps{
		LoginEP:      loginEP,
		CreateUserEP: createUserEP,
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
