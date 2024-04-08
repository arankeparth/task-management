package main

import (
	"fmt"
	"net/http"

	"task-management/server/authservice"
	"task-management/server/spec/authspec"

	"github.com/gorilla/mux"
)

func main() {
	As := authservice.NewAuthService()
	AuthHandler := As.InitAuthServiceHandler()
	Mux := mux.NewRouter()
	Mux.Handle(authspec.BasePath, AuthHandler)
	errs := make(chan error, 100)

	go func() {
		errs <- http.ListenAndServe(authspec.Host, accessControl(AuthHandler))
	}()
	err := <-errs
	fmt.Println(err.Error())
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, public_key, Access-Control-Allow-Origin")
		w.Header().Set("Referrer-Policy", "same-origin")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
