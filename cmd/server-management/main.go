package main

import (
	"flag"

	"sm/internal/database"
)

func main() {
	host := flag.String("host", "localhost", "Enter host (localhost): ")
	port := flag.Int("port", 5432, "Enter port (5432): ")
	username := flag.String("user", "postgres", "Enter user (postgres): ")
	password := flag.String("password", "password", "Enter password: ")
	dbname := flag.String("db", "database", "Enter database name: ")

	flag.Parse()

	db := database.ConnectToDB(*host, *port, *username, *password, *dbname)

	defer db.Close()

}
