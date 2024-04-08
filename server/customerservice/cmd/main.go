package main

import (
	"fmt"
	"net/http"

	customerApi "task-management/server/customerservice"
	"task-management/server/spec/customerspec"

	"github.com/gorilla/mux"
)

func main() {
	cs := customerApi.NewCustomerService()
	Handler := cs.InitCustomerServiceHandler()
	Mux := mux.NewRouter()
	Mux.Handle(customerspec.BasePath, Handler)
	errs := make(chan error, 100)
	go func() {
		errs <- http.ListenAndServe(customerspec.Host, accessControl(Handler))
	}()
	err := <-errs
	fmt.Println(err.Error())
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
