package dbconn

import (
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	taskManagementContainer = "task-management-database-1"
)

func NewDB(dbName string) (*sqlx.DB, error) {
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		Net:                  "tcp",
		Addr:                 taskManagementContainer,
		DBName:               dbName,
		AllowNativePasswords: true,
	}
	db, err := sqlx.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println("error:", err.Error())
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("error:", err.Error())
		log.Fatal(err)
	}
	if err != nil {
		fmt.Println("error:", err.Error())
		return nil, err
	}
	db.SetMaxIdleConns(0)
	return db, nil
}
