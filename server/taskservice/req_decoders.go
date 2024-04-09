package taskservice

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"task-management/server/spec/taskspec"

	"github.com/gorilla/mux"
)

// Function to log errors
func logError(err error, functionName string) {
	if err != nil {
		// You can customize the logging format as needed
		log.Printf("Error in %s: %v\n", functionName, err)
	}
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeCreateTaskRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body := taskspec.CreateTaskRequest{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logError(err, "decodeCreateTaskRequest")
		return nil, err
	}
	return body, nil
}

func decodeUpdateTaskRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body := taskspec.Task{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logError(err, "decodeUpdateTaskRequest")
		return nil, err
	}
	return body, nil
}

func decodeGetTasksRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	body := taskspec.GetTasksRequest{}
	ctx = context.WithValue(ctx, "Authorization", r.Header.Get("Authorization"))
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logError(err, "decodeGetTasksRequest")
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
		logError(err, "decodeCreateCommentRequest")
		return nil, err
	}
	return body, nil
}

func decodeGetCommentsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	taskID, err := strconv.Atoi(mux.Vars(r)["taskID"])
	if err != nil {
		logError(err, "decodeGetCommentsRequest")
		return nil, err
	}
	if taskID == 0 {
		errMsg := "taskID is required"
		logError(errors.New(errMsg), "decodeGetCommentsRequest")
		return nil, errors.New(errMsg)
	}
	return taskspec.Comment{
		TicketID: taskID,
	}, nil
}
