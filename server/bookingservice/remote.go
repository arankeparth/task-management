package bookingservice

import (
	"errors"
	bookingspec "plantrip-backend/server/spec/bookingservice"

	"github.com/go-resty/resty/v2"
)

type RemoteSdk struct {
	sdkClient *resty.Client
}

func (r *RemoteSdk) GetAirPlaneSchedule(src, dest, date string) (*bookingspec.GetScheduleResponse, error) {
	params := map[string]string{
		"From": src,
		"TO":   dest,
		"Date": date,
	}
	url := "https://timetable-lookup.p.rapidapi.com/TimeTable/{From}/{To}/{Date}"
	resp, _ := r.sdkClient.R().
		SetHeader("X-RapidAPI-Key", "698fb1ad51msh8459d1a25c2f750p15373cjsn00ad8f332954").
		SetHeader("X-RapidAPI-Host", "timetable-lookup.p.rapidapi.com").
		SetPathParams(params).
		Post(url)

	if resp.IsError() {
		return &bookingspec.GetScheduleResponse{}, errors.New("failed to get the airplace schedule")
	}
	parsedResponse, err := r.parseResponse(resp)
	if err != nil {
		return &bookingspec.GetScheduleResponse{}, err
	}
	return parsedResponse, nil
}

func (r *RemoteSdk) parseResponse(rawResponse *resty.Response) (*bookingspec.GetScheduleResponse, error) {
	return &bookingspec.GetScheduleResponse{}, nil
}
