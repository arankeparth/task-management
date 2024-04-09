package customerApi

import (
	"fmt"
	"errors"
	"task-management/server/spec/authspec"
	"task-management/server/spec/customerspec"
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
	
	fmt.Println(resp)
	if resp.IsError() {
		fmt.Println("err")
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
