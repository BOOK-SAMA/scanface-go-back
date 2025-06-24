package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/testme")
	if err != nil {
		panic(err)
	}
	// Open doesn't open a connection. Validate DSN data:
	if err := pingdb(DB); err != nil {
		// พิมพ์ข้อความ error
		fmt.Println("Error connecting to DB:", err)
		panic(err) // หรือจะไม่ panic ก็ได้ ขึ้นอยู่กับการจัดการของโปรแกรมคุณ
	}
}

func pingdb(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return fmt.Errorf("%s ไม่สามารถ ping database ได้", err.Error())
	}
	return nil
}
