package bookingservice

import (
	"net/http"
	bookingspec "plantrip-backend/server/spec/bookingservice"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandler(eps *BookingEps) http.Handler {

	getScheduledHandler := kithttp.NewServer(
		eps.GetSchedulesEP,
		decodeGetSchedulesRequest,
		encodeResponse,
	)

	r := mux.NewRouter()
	r.Handle(bookingspec.GetSchedulesPath, getScheduledHandler).Methods("GET")
	return r
}
