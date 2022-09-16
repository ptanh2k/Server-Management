package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectToDB(host string, port int, user string, password string, dbname string) *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected")

	return db
}
