package sqlConnect

import (
	"database/sql"
	"log"
)

var dB *sql.DB

func ConnectDB() *sql.DB {
	//Get the database connection

	//Get credentials from the environment
	connectionString := "username:password@tcp(localhost:3306)/dbname"
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
