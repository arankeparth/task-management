package authdl

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type AuthDl struct {
	DB *sqlx.DB
}

func NewAuthDl(db *sqlx.DB) *AuthDl {
	return &AuthDl{
		DB: db,
	}
}

type CustomerInfo struct {
	CustomerId int64  `db:"customerid"`
	Password   string `db:"password"`
}

const (
	credsTable           = "customer_creds"
	publicKeyTable       = "customer_pubs"
	pubKeyColumnName     = "pub_key"
	customerIdColumnName = "customerid"
)

func (dl *AuthDl) GetInfo(username string) (string, int64) {
	query := fmt.Sprintf("SELECT customerid,  password from %s WHERE username='%s';", credsTable, username)
	var password string
	var customerId int64
	dl.DB.QueryRow(query).Scan(&customerId, &password)
	return password, customerId
}

func (dl *AuthDl) DeleteInfo(customerId int64) error {
	query := fmt.Sprintf("DELETE from %s WHERE cutomerid=%d;", credsTable, customerId)
	_, err := dl.DB.Query(query)
	if err != nil {
		return err
	}
	return nil
}

func (dl *AuthDl) SetPubKey(publicKey string, customerid int64) error {
	query := fmt.Sprintf("INSERT into %s values(%d, '%s');", publicKeyTable, customerid, publicKey)
	_, err := dl.DB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (dl *AuthDl) UpdatePubKey(publicKey string, customerId int64) error {
	query := fmt.Sprintf("UPDATE %s SET %s ='%s' WHERE %s = %d;", publicKeyTable, pubKeyColumnName, publicKey, customerIdColumnName, customerId)
	_, err := dl.DB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (dl *AuthDl) DeletePubKey(customerID int64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = %d;", publicKeyTable, customerIdColumnName, customerID)
	_, err := dl.DB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (dl *AuthDl) GetPubKey(customerid int64) string {
	query := fmt.Sprintf("SELECT %s from %s where %s = %d;", pubKeyColumnName, publicKeyTable, customerIdColumnName, customerid)
	var pubKey string
	dl.DB.QueryRow(query).Scan(&pubKey)
	return pubKey
}

func (dl *AuthDl) CreateUser(username string, password string, customerid int64) error {
	query := fmt.Sprintf("INSERT into %s values(%d, '%s', '%s');", credsTable, customerid, username, password)
	_, err := dl.DB.Exec(query)
	if err != nil {
		print(err.Error())
		return err
	}
	return nil
}
