package bookingspec

type GetScheduleRequest struct {
	CustomerId    int32  `json:"customerid"`
	Source        string `json:"source"`
	Destination   string `json:"destination"`
	TravelDate    string `json:"travel_date"`
	TransportMode int32  `json:"transport_mode"`
}
