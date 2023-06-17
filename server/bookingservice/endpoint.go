package bookingservice

import (
	"context"
	"errors"
	"log"
	bookingspec "plantrip-backend/server/spec/bookingservice"

	"github.com/go-kit/kit/endpoint"
)

type BookingEps struct {
	GetSchedulesEP endpoint.Endpoint
}

func NewCustomerEndpoints(h *BookingHandler, logger *log.Logger) (*BookingEps, error) {
	getSchedulesEp := makeGetSchedulesEP(h)
	return &BookingEps{
		GetSchedulesEP: getSchedulesEp,
	}, nil
}

func makeGetSchedulesEP(h *BookingHandler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*bookingspec.GetScheduleRequest)
		if !ok {
			return nil, errors.New("error while converting types")
		}
		resp, err := h.GetSchedules(req)
		return resp, err
	}
}
