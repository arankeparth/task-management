package main

import (
	"log"
	"net/http"
	"os"

	customerApi "task-management/server/customerservice"
	"task-management/server/spec/customerspec"

	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(os.Stdout, "[CustomerService Main] ", log.LstdFlags)
	logger.Printf("Starting customer service")
	logger.Printf("Initializing customer service")
	cs := customerApi.NewCustomerService()
	logger.Printf("Initializing customer service handler")
	Handler := cs.InitCustomerServiceHandler()
	Mux := mux.NewRouter()
	Mux.Handle(customerspec.BasePath, Handler)
	errs := make(chan error, 100)
	logger.Printf("Listening on %s", customerspec.Port)
	go func() {
		errs <- http.ListenAndServe(customerspec.Port, accessControl(Handler))
	}()
	err := <-errs
	logger.Printf("Error in customer service: %s", err.Error())
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, public_key")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
