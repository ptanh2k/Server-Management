package main

import (
	"flag"

	"sm/internal/database"

	"sm/internal/rest"

	"github.com/gin-gonic/gin"
)

func main() {
	host := flag.String("host", "localhost", "Enter host (localhost): ")
	port := flag.Int("port", 5432, "Enter port (5432): ")
	username := flag.String("user", "postgres", "Enter user (postgres): ")
	password := flag.String("password", "password", "Enter password: ")
	dbname := flag.String("db", "database", "Enter database name: ")

	flag.Parse()

	db, err := database.ConnectToDB(*host, *port, *username, *password, *dbname)

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/servers", rest.GetAllServers(db))

	r.Run(":8080")
}
