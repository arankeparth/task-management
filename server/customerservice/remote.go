package customerApi

import (
	"errors"
	"plantrip-backend/server/spec/authspec"
	"plantrip-backend/server/spec/customerspec"
)

func (h *CustomerHandler) CreateUserCreds(req *customerspec.CreateCustomerRequest) error {
	payload := map[string]interface{}{
		"username":   req.Email,
		"password":   req.Password,
		"customerid": req.CustomerId,
	}

	resp, _ := h.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(h.Client.BaseURL + authspec.CreateUserPath)

	if resp.IsError() {
		return errors.New("failed to create a new user")
	}

	return nil
}

func (h *CustomerHandler) VerifyJwt(tokenString string, publicKey string) error {
	payload := map[string]interface{}{
		"token_string":   tokenString,
		"pub_key_string": publicKey,
	}
	resp, _ := h.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(h.Client.BaseURL + authspec.VerifyJwtPath)

	if resp.IsError() {
		return errors.New("authentication failed")
	}

	return nil
}
