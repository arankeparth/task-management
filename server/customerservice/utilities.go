package customerApi

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetCustomerId(r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	customerId, err := strconv.Atoi(vars["customerid"])
	if err != nil {
		return 0, err
	}
	if customerId == 0 {
		return 0, errors.New("invalid customerid")
	}
	return int64(customerId), nil
}
