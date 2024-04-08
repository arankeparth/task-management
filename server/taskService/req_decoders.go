package taskservice

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"task-management/server/spec/taskspec"

	"github.com/gorilla/mux"
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeCreateTaskRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body := taskspec.CreateTaskRequest{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}
	return body, nil
}

func decodeUpdateTaskRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body := taskspec.Task{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}
	return body, nil
}

func decodeGetTasksRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	body := taskspec.GetTasksRequest{}
	ctx = context.WithValue(ctx, "Authorization", r.Header.Get("Authorization"))
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}
	return body, nil
}

func decodeGetTaskByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	taskID := params["taskID"]
	return taskID, nil
}

func decodeCreateCommentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body := taskspec.Comment{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}
	return body, nil
}

func decodeGetCommentsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	taskID, err := strconv.Atoi(mux.Vars(r)["taskID"])
	if err != nil {
		return nil, err
	}
	if taskID == 0 {
		return nil, errors.New("taskID is required")
	}
	return taskspec.Comment{
		TicketID: taskID,
	}, nil
}
