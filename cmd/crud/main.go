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
	r.GET("/servers/:id", rest.GetServerWithId(db))
	r.POST("/servers", rest.CreateNewServer(db))
	// r.POST("/_bulk/servers", rest.CreateMultipleServers(db))
	r.PATCH("/servers/:id", rest.UpdateServer(db))
	r.DELETE("/servers/:id", rest.DeleteServer(db))

	r.Run(":8080")
}
