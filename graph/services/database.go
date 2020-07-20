package services

import (
	"database/sql"
	// Get mysql library
	_ "github.com/go-sql-driver/mysql"
)

// DB - export db instance
var DB *sql.DB

// CreateDatabase assign db instance
func CreateDatabase() {
	db, err := sql.Open("mysql", "user:test@tcp(127.0.0.1:3306)/mydb?parseTime=true")

	if err != nil {
		panic(err)
	}

	DB = db
}
