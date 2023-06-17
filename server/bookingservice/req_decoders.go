package bookingservice

import (
	"context"
	"encoding/json"
	"net/http"
	"plantrip-backend/server/spec/authspec"
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeGetSchedulesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body := &authspec.LoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}
	return body, nil
}
