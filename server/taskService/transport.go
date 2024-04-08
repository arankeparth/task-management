package taskservice

import (
	"net/http"
	"task-management/server/middleware"
	"task-management/server/spec/taskspec"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandler(eps *TaskEps) http.Handler {
	jwtMiddleWare := middleware.JWTMiddleware
	createTaskHandler := kithttp.NewServer(
		eps.CreateTaskEP,
		decodeCreateTaskRequest,
		encodeResponse,
	)

	updateTaskHandler := kithttp.NewServer(
		eps.UpdateTaskEP,
		decodeUpdateTaskRequest,
		encodeResponse,
	)
	getTasksHandler := kithttp.NewServer(
		eps.GetTasksEP,
		decodeGetTasksRequest,
		encodeResponse,
	)

	getTaskByIDHandler := kithttp.NewServer(
		eps.GetTaskByIDEP,
		decodeGetTaskByIDRequest,
		encodeResponse,
	)

	createCommentHandler := kithttp.NewServer(
		eps.CreateCommentsEP,
		decodeCreateCommentRequest,
		encodeResponse,
	)

	getCommentsHandler := kithttp.NewServer(
		eps.GetCommentsEP,
		decodeGetCommentsRequest,
		encodeResponse,
	)

	r := mux.NewRouter()
	r.Handle(taskspec.CreateTaskPath, jwtMiddleWare(createTaskHandler)).Methods("POST")
	r.Handle(taskspec.UpdateTaskPath, updateTaskHandler).Methods("PUT")
	r.Handle(taskspec.GetTasksPath, jwtMiddleWare(getTasksHandler)).Methods("POST")
	r.Handle(taskspec.GetTaskByIDPath, getTaskByIDHandler).Methods("GET")
	r.Handle(taskspec.CreateCommentPath, jwtMiddleWare(createCommentHandler)).Methods("POST")
	r.Handle(taskspec.GetCommentsPath, jwtMiddleWare(getCommentsHandler)).Methods("GET")

	return r
}
