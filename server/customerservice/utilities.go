package customerApi

import (
	"errors"
	"net/http"
	customerspec "task-management/server/spec/customerspec"

	"github.com/gorilla/mux"
)

func GetCustomerId(r *http.Request) (string, error) {
	vars := mux.Vars(r)
	customerId, _ := vars[customerspec.CustomerIdKey]

	if customerId == "" {
		return "", errors.New("invalid customerid")
	}
	return customerId, nil
}
