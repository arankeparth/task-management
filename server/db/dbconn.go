package dbconn

import (
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDB(dbName string) (*sqlx.DB, error) {
	cfg := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASS"),
		Net:    "tcp",
		Addr:   "localhost:3306",
		//Addr:                 os.Getenv("DB_HOST"),
		DBName:               dbName,
		AllowNativePasswords: true,
	}
	db, err := sqlx.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		fmt.Println("error:", err.Error())
		return nil, err
	}
	db.SetMaxIdleConns(0)
	return db, nil
}
