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
	fmt.Println("DB_USER", os.Getenv("DB_USER"))
	fmt.Println("DB_PASS", os.Getenv("DB_PASS"))
	fmt.Println("DB_HOST", os.Getenv("DB_HOST"))
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               "parth123",
		Net:                  "tcp",
		Addr:                 "task-management-mysql-1",
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
	fmt.Println("hiii")
	return db, nil
}
