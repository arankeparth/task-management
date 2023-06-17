package bookingservice

import (
	"errors"
	"log"
	bookingspec "plantrip-backend/server/spec/bookingservice"
)

type BookingHandler struct {
	remoteSdk *RemoteSdk
}

func (bh *BookingHandler) GetSchedules(req *bookingspec.GetScheduleRequest) (*bookingspec.GetScheduleResponse, error) {
	if req.TransportMode == bookingspec.TRANSPORT_MODE_AIRPLANE {
		resp, err := bh.remoteSdk.GetAirPlaneSchedule(req.Source, req.Destination, req.TravelDate)
		log.Printf("[INFO] CustomerId: %d Method: %s Getting schedules for airplane transport mode.", req.CustomerId, "GetSchedules")
		if err != nil {
			log.Printf("[ERROR] CustomerId: %d Method: %s ErrorMessage: %s", req.CustomerId, "GetSchedules", err.Error())
			return &bookingspec.GetScheduleResponse{}, err
		}
		return resp, nil
	} else if req.TransportMode == bookingspec.TRANSPORT_MODE_TRAIN {
		return &bookingspec.GetScheduleResponse{}, nil
	} else if req.TransportMode == bookingspec.TRANSPORT_MODE_BUS {
		return &bookingspec.GetScheduleResponse{}, nil
	}

	log.Printf("[ERROR] CustomerId: %d Method: %s ErrorMessage: %s", req.CustomerId, "GetSchedules", "invalid transport mode")
	return &bookingspec.GetScheduleResponse{}, errors.New("invalid transport mode")
}
