package customerApi

import (
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

	if resp.IsError() {
		return errors.New("failed to create a new user")
	}

	return nil
}
