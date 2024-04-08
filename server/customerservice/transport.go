package customerApi

import (
	"net/http"

	"task-management/server/middleware"
	"task-management/server/spec/customerspec"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandler(eps *CustomerEps) http.Handler {
	opts := []kithttp.ServerOption{}
	JWTMiddleware := middleware.JWTMiddleware
	createCustomerHandler := kithttp.NewServer(
		eps.CreateCustomerEP,
		decodeCreateCustomerRequest,
		encodeResponse,
		opts...,
	)
	getCustomerHandler := kithttp.NewServer(
		eps.GetCustomerEP,
		decodeGetCustomerRequest,
		encodeResponse,
		opts...,
	)
	deleteCustomerHandler := kithttp.NewServer(
		eps.DeleteCustomerEP,
		decodeDeleteCustomerRequest,
		encodeResponse,
		opts...,
	)
	getOffersHadler := kithttp.NewServer(
		eps.GetOffersEP,
		decodeGetOffersRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()
	r.Handle(customerspec.CreateCustomerPath, createCustomerHandler).Methods("POST")
	r.Handle(customerspec.GetCustomerPath, JWTMiddleware(getCustomerHandler)).Methods("GET")
	r.Handle(customerspec.DeleteCustomerPath, JWTMiddleware(deleteCustomerHandler)).Methods("DELETE")
	r.Handle(customerspec.GetOffersPath, JWTMiddleware(getOffersHadler)).Methods("GET")

	return r
}
