package customerApi

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCustomerId(r *http.Request) (string, error) {
	vars := mux.Vars(r)
	customerId, _ := vars["customerid"]

	if customerId == "" {
		return "", errors.New("invalid customerid")
	}
	return customerId, nil
}
