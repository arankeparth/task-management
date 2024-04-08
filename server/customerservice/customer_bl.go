package customerApi

import (
	"errors"
	customerdl "task-management/server/customerservice/dl"
	"task-management/server/spec/customerspec"

	"github.com/go-resty/resty/v2"
)

type CustomerHandler struct {
	Dl     *customerdl.CustomerDL
	Client *resty.Client
}

func NewCustomerHandler(dl *customerdl.CustomerDL, client *resty.Client) *CustomerHandler {
	return &CustomerHandler{
		Dl:     dl,
		Client: client,
	}
}

func (h *CustomerHandler) Create(req *customerspec.CreateCustomerRequest) error {
	err := h.Dl.CreateCustomer(req)
	if err != nil {
		return err
	}
	err = h.CreateUserCreds(req)
	if err != nil {
		return err
	}
	return nil
}

func (h *CustomerHandler) Destroy(req *customerspec.DeleteCustomerRequest) error {
	err := h.Dl.DeleteCustomer(req)
	if err != nil {
		return err
	}
	return nil
}

func (h *CustomerHandler) Update(req *customerspec.UpdateCustomerRequest) error {
	err := h.Dl.UpdateCustomer(req)
	if err != nil {
		return err
	}
	return nil
}

func (h *CustomerHandler) GetInfo(req *customerspec.GetCustomerRequest) (*customerspec.GetCustomerResponse, error) {
	info := h.Dl.GetCustomer(req)
	if info == nil {
		return nil, errors.New("failed to get info")
	}
	return info, nil
}

func (h *CustomerHandler) GetOffers(customerId int64) ([]customerspec.GetOffersResponse, error) {
	resp, err := h.Dl.GetOffers(customerId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
