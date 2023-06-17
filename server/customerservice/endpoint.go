package customerApi

import (
	"context"
	"errors"
	"plantrip-backend/server/spec/customerspec"

	"github.com/go-kit/kit/endpoint"
)

type CustomerEps struct {
	CreateCustomerEP endpoint.Endpoint
	DeleteCustomerEP endpoint.Endpoint
	GetCustomerEP    endpoint.Endpoint
	UpdateCustomerEP endpoint.Endpoint
	GetOffersEP      endpoint.Endpoint
}

func NewCustomerEndpoints(ch *CustomerHandler) (*CustomerEps, error) {
	var err error
	jwtMiddleWare := ch.JWTMiddleware()
	createCustomerEP := makeCreateCustomerEP(ch)
	if createCustomerEP == nil {
		err = errors.New("failed to create createCustomerEP")
		return nil, err
	}

	deleteCustomerEP := makeDeleteCustomerEP(ch)
	if deleteCustomerEP == nil {
		err = errors.New("failed to create deleteCustomerEP")
		return nil, err
	}

	getCustomerEP := makeGetCustomerEP(ch)
	if getCustomerEP == nil {
		err = errors.New("failed to create getCustomerEP")
		return nil, err
	}

	updateCustomerEP := makeUpdateCustomerEP(ch)
	if updateCustomerEP == nil {
		err = errors.New("failed to create updateCustomerEP")
		return nil, err
	}

	getOffersEP := makeGetOffersEP(ch)

	return &CustomerEps{
		CreateCustomerEP: createCustomerEP,
		DeleteCustomerEP: jwtMiddleWare(deleteCustomerEP),
		GetCustomerEP:    jwtMiddleWare(getCustomerEP),
		UpdateCustomerEP: jwtMiddleWare(updateCustomerEP),
		GetOffersEP:      jwtMiddleWare(getOffersEP),
	}, nil
}

func makeCreateCustomerEP(ch *CustomerHandler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(customerspec.CreateCustomerRequest)
		if !ok {
			return nil, errors.New("error while converting types")
		}
		err = ch.Create(&req)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}

func makeGetCustomerEP(ch *CustomerHandler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(customerspec.GetCustomerRequest)
		if !ok {
			return nil, errors.New("error while converting types")
		}
		info, err := ch.GetInfo(&req)
		if err != nil {
			return nil, err
		}
		return info, nil
	}
}

func makeUpdateCustomerEP(ch *CustomerHandler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(customerspec.UpdateCustomerRequest)
		if !ok {
			return nil, errors.New("error while converting types")
		}
		err = ch.Update(&req)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}

func makeDeleteCustomerEP(ch *CustomerHandler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(customerspec.DeleteCustomerRequest)
		if !ok {
			return nil, errors.New("error while converting types")
		}
		err = ch.Destroy(&req)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}

func makeGetOffersEP(ch *CustomerHandler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(customerspec.GetOffersRequest)
		resp, err := ch.GetOffers(req.CustomerId)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}
