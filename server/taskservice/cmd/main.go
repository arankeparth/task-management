package main

import (
	"log"
	"net/http"
	"os"

	"task-management/server/spec/taskspec"
	"task-management/server/taskservice"

	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(os.Stdout, "[TaskService Main] ", log.LstdFlags)
	logger.Printf("Starting task service")
	logger.Printf("Initializing task service")
	taskservice := taskservice.NewTaskService()
	logger.Printf("Initializing task service handler")
	Handler := taskservice.InitTaskServiceHandler()
	Mux := mux.NewRouter()
	Mux.Handle(taskspec.BasePath, Handler)
	errs := make(chan error, 100)
	logger.Printf("Listening on %s", taskspec.Port)
	go func() {
		errs <- http.ListenAndServe(taskspec.Port, accessControl(Handler))
	}()
	err := <-errs
	logger.Printf("Error in task service: %s", err.Error())
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
