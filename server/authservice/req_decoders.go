package authservice

import (
	"context"
	"encoding/json"
	"net/http"
	"plantrip-backend/server/spec/authspec"
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body := &authspec.LoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}
	return body, nil
}

func decodeCreateUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	body := &authspec.CreateUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}
	return body, nil
}

func decodeValidateJwtRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	body := &authspec.VerifyJwtReq{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}
	return body, nil
}
