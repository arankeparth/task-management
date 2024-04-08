package main

import (
	"fmt"
	"net/http"

	"task-management/server/spec/taskspec"
	"task-management/server/taskservice"

	"github.com/gorilla/mux"
)

func main() {
	taskservice := taskservice.NewTaskService()
	Handler := taskservice.InitTaskServiceHandler()
	Mux := mux.NewRouter()
	Mux.Handle(taskspec.BasePath, Handler)
	errs := make(chan error, 100)
	go func() {
		errs <- http.ListenAndServe(taskspec.Host, accessControl(Handler))
	}()
	err := <-errs
	fmt.Println(err.Error())
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, public_key")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
