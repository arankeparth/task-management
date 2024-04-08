package taskservice

import (
	"errors"
	"log"
	"os"
	"task-management/server/spec/taskspec"
	"task-management/server/taskservice/dl"
)

type TaskService struct {
	dl     *dl.TaskDL
	logger *log.Logger
}

func NewTaskBL(DL *dl.TaskDL) *TaskService {
	logger := log.New(os.Stdout, "[TaskService] ", log.LstdFlags)
	return &TaskService{
		dl:     DL,
		logger: logger,
	}
}

func (ts *TaskService) CreateTask(task *taskspec.CreateTaskRequest) error {
	err := ts.dl.CreateTask(task)
	if err != nil {
		ts.logger.Printf("Error creating task: %s", err.Error())
		return err
	}
	return nil
}

func (ts *TaskService) GetTasks(req *taskspec.GetTasksRequest) (taskspec.GetTaskResponse, error) {
	tasks, err := ts.dl.GetTasks(req)
	if err != nil {
		ts.logger.Printf("Error getting tasks: %s", err.Error())
		return nil, err
	}

	return tasks, nil
}

func (ts *TaskService) GetTaskByID(id int) (*taskspec.Task, error) {
	task, err := ts.dl.GetTaskByID(id)
	if err != nil {
		ts.logger.Printf("Error getting task by ID: %s", err.Error())
		return nil, err
	}

	return task, nil
}

func (ts *TaskService) UpdateTask(task *taskspec.Task) error {
	err := ts.dl.UpdateTask(task)
	if err != nil {
		ts.logger.Printf("Error updating task: %s", err.Error())
		return err
	}
	return nil
}

func (ts *TaskService) DeleteTask(id int) error {
	return errors.New("not implemented")
}

func (ts *TaskService) GetComments(taskID int) ([]taskspec.Comment, error) {
	comments, err := ts.dl.GetComments(taskID)
	if err != nil {
		ts.logger.Printf("Error getting comments: %s", err.Error())
		return nil, err
	}

	return comments, nil
}

func (ts *TaskService) CreateComment(comment *taskspec.Comment) error {
	err := ts.dl.CreateComment(comment)
	if err != nil {
		ts.logger.Printf("Error creating comment: %s", err.Error())
		return err
	}

	return nil
}

func (ts *TaskService) CreateNotification(notification *taskspec.Notification) error {
	err := ts.dl.CreateNotification(notification)
	if err != nil {
		ts.logger.Printf("Error creating notification: %s", err.Error())
		return err
	}
	return nil
}
