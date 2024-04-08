package taskservice

import (
	"context"
	"errors"
	"fmt"
	"task-management/server/spec/taskspec"

	"github.com/go-kit/kit/endpoint"
)

type TaskEps struct {
	CreateTaskEP     endpoint.Endpoint
	GetTasksEP       endpoint.Endpoint
	UpdateTaskEP     endpoint.Endpoint
	DeleteTaskEP     endpoint.Endpoint
	GetTaskByIDEP    endpoint.Endpoint
	GetCommentsEP    endpoint.Endpoint
	CreateCommentsEP endpoint.Endpoint
}

func NewTaskServiceEndPoints(t *TaskService) (*TaskEps, error) {
	createTasksEP := makeCreateTaskEP(t)
	getTasksEP := makeGetTasksEP(t)
	updateTaskEP := makeUpdateTaskEP(t)
	deleteTaskEP := makeDeleteTaskEP(t)
	getTaskByIDEP := makeGetTaskByIDEP(t)
	getCommentsEP := makeGetCommentsEP(t)
	createCommentsEP := makeCreateCommentsEP(t)

	return &TaskEps{
		createTasksEP,
		getTasksEP,
		updateTaskEP,
		deleteTaskEP,
		getTaskByIDEP,
		getCommentsEP,
		createCommentsEP,
	}, nil
}

func makeCreateTaskEP(t *TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(taskspec.CreateTaskRequest)
		fmt.Println(request)
		if !ok {
			return nil, errors.New("error while converting types")
		}
		err = t.CreateTask(&req)
		return nil, err
	}
}

func makeGetTasksEP(t *TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		fmt.Println("create task endpoint")
		fmt.Println(request)
		req, ok := request.(taskspec.GetTasksRequest)
		if !ok {
			return nil, errors.New("error while converting types")
		}
		tasks, err := t.GetTasks(&req)
		if err != nil {
			return nil, err
		}
		return tasks, nil
	}
}

func makeUpdateTaskEP(t *TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(taskspec.Task)
		if !ok {
			return nil, errors.New("error while converting types")
		}
		err = t.UpdateTask(&req)
		return nil, err
	}
}

func makeDeleteTaskEP(t *TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*taskspec.Task)
		if !ok {
			return nil, errors.New("error while converting types")
		}
		err = t.DeleteTask(req.ID)
		return nil, err
	}
}

func makeGetTaskByIDEP(t *TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(int)
		if !ok {
			return nil, errors.New("error while converting types")
		}
		task, err := t.GetTaskByID(req)
		if err != nil {
			return nil, err
		}
		return task, nil
	}
}

func makeGetCommentsEP(t *TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(taskspec.Comment)
		if !ok {
			return nil, errors.New("error while converting types")
		}
		comments, err := t.GetComments(req.TicketID)
		if err != nil {
			return nil, err
		}
		return comments, nil
	}
}

func makeCreateCommentsEP(t *TaskService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		fmt.Printf("%+v\n", request)
		req, ok := request.(taskspec.Comment)
		if !ok {
			return nil, errors.New("error while converting types")
		}
		err = t.CreateComment(&req)
		return nil, err
	}
}
