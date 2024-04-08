package customerApi

import (
	"log"
	"net/http"
	dl "task-management/server/customerservice/dl"
	dbconn "task-management/server/db"

	"github.com/go-resty/resty/v2"
)

type CustomerService struct {
}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (cs *CustomerService) InitCustomerServiceHandler() http.Handler {
	CustomerDB := "customerservice"
	client := resty.New()
	client.BaseURL = "http://localhost:8080"
	Db, err := dbconn.NewDB(CustomerDB)
	if err != nil {
		log.Print(err.Error())
		return nil
	}

	Dl := dl.NewCustomerDL(Db)
	Bl := NewCustomerHandler(Dl, client)

	CustomerEps, err := NewCustomerEndpoints(Bl)
	if err != nil {
		log.Printf("Failed to create authentication endpoints %s", err.Error())
		return nil
	}

	HttpHandler := MakeHandler(CustomerEps)
	return HttpHandler
}
