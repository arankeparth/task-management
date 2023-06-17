package bookingspec

type GetScheduleResponse struct {
	Source        string         `json:"source"`
	Destination   string         `json:"destination"`
	TravelDate    string         `json:"travel_date"`
	TransportMode string         `json:"transport_mode"`
	Schedule      []ScheduleItem `json:"schedule"`
}

type ScheduleItem struct {
	Time     string  `json:"time"`
	Duration string  `json:"duration"`
	Route    string  `json:"route"`
	Price    float64 `json:"price"`
}
