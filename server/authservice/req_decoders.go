package authservice

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"task-management/server/spec/authspec"
)

// logError function to log errors
func logError(err error, functionName string) {
	if err != nil {
		log.Printf("Error in %s: %v\n", functionName, err)
	}
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func createSessionCookie(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	return json.NewEncoder(w).Encode(resp)
}

func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body := &authspec.LoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logError(err, "decodeLoginRequest")
		return nil, err
	}
	return body, nil
}

func decodeCreateUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	body := &authspec.CreateUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logError(err, "decodeCreateUserRequest")
		return nil, err
	}
	return body, nil
}

func decodeValidateJwtRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	body := &authspec.VerifyJwtReq{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logError(err, "decodeValidateJwtRequest")
		return nil, err
	}
	return body, nil
}
