package customerApi

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"task-management/server/spec/customerspec"
)

// logError function to log errors
func logError(err error, functionName string) {
	if err != nil {
		log.Printf("Error in %s: %v\n", functionName, err)
	}
}

func decodeCreateCustomerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body := customerspec.CreateCustomerRequest{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logError(err, "decodeCreateCustomerRequest")
		return nil, err
	}
	return body, nil
}

func decodeGetCustomerRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	customerId, err := GetCustomerId(r)
	if err != nil {
		logError(err, "decodeGetCustomerRequest")
		return nil, err
	}
	return customerspec.GetCustomerRequest{
		CustomerId: customerId,
	}, nil
}

func decodeDeleteCustomerRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	customerId, err := GetCustomerId(r)
	if err != nil {
		logError(err, "decodeDeleteCustomerRequest")
		return nil, err
	}

	TokenString := r.Header.Get("Authorization")
	PublicKey := r.Header.Get("public_key")
	ctx = context.WithValue(ctx, "token_string", TokenString)
	ctx = context.WithValue(ctx, "public_key", PublicKey)

	return &customerspec.DeleteCustomerRequest{
		CustomerId: customerId,
	}, nil
}

func decodeGetOffersRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	customerId, err := GetCustomerId(r)
	if err != nil {
		logError(err, "decodeGetOffersRequest")
		return nil, err
	}

	req := customerspec.GetOffersRequest{}
	req.TokenString = r.Header.Get("Authorization")
	req.PublicKey = r.Header.Get("public_key")
	req.CustomerId = customerId
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
