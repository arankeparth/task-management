package customerdl

import (
	"fmt"
	"plantrip-backend/server/spec/customerspec"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerDL struct {
	DB *sqlx.DB
}

func NewCustomerDL(db *sqlx.DB) *CustomerDL {
	return &CustomerDL{
		DB: db,
	}
}

const (
	CustomerInfoTable  = "customer_info"
	CustomerCredsTable = "customer_creds"
	OffersTable        = "offers"
)

func (DL *CustomerDL) CreateCustomer(req *customerspec.CreateCustomerRequest) error {
	query := fmt.Sprintf("INSERT into %s values(%d, '%s', '%s', '%s', '%d', '%s')", CustomerInfoTable,
		req.CustomerId,
		req.FirstName,
		req.LastName,
		req.Email,
		req.Age,
		req.Gender)
	_, err := DL.DB.Exec(query)

	if err != nil {
		print(err.Error())
		return err
	}
	return nil
}

func (DL *CustomerDL) DeleteCustomer(req *customerspec.DeleteCustomerRequest) error {
	query := fmt.Sprintf("DELETE from %s WHERE customerid=%d", CustomerInfoTable, req.CustomerId)
	_, err := DL.DB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (DL *CustomerDL) UpdateCustomer(req *customerspec.UpdateCustomerRequest) error {
	query := fmt.Sprintf("UPDATE %s SET values(%d, '%s', '%s', '%s', '%s', '%s') WHERE cusomterid = %d", CustomerInfoTable,
		req.CustomerId,
		req.FirstName,
		req.LastName,
		req.Email,
		req.Age,
		req.Gender,
		req.CustomerId)
	_, err := DL.DB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (DL *CustomerDL) GetCustomer(req *customerspec.GetCustomerRequest) *customerspec.GetCustomerResponse {
	query := fmt.Sprintf("SELECT * from %s  WHERE customerid=%d", CustomerInfoTable, req.CustomerId)
	resp := &customerspec.GetCustomerResponse{}
	DL.DB.QueryRow(query).Scan(resp)
	return resp
}

func (DL *CustomerDL) GetOffers(customerId int64) ([]customerspec.GetOffersResponse, error) {
	query := fmt.Sprintf("SELECT * from %s where customerid=%d", OffersTable, customerId)
	rows, err := DL.DB.Query(query)
	if err != nil {
		fmt.Println("error")
		return nil, err
	}
	resp := []customerspec.GetOffersResponse{}
	for rows.Next() {
		var r customerspec.GetOffersResponse
		err := rows.Scan(&r.Title, &r.Desc, &r.CustomerId)
		if err != nil {
			return nil, err
		}
		resp = append(resp, r)
	}

	return resp, nil
}
