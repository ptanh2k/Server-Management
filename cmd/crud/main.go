package main

import (
	"fmt"
	"log"
	"os"

	"sm/internal/database"

	"sm/api/rest"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"

	"strconv"
)

func main() {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("SM_HOST")
	dbport := os.Getenv("SM_PORT")
	username := os.Getenv("SM_USERNAME")
	password := os.Getenv("SM_PASSWORD")
	dbname := os.Getenv("SM_DB_NAME")

	port, err := strconv.Atoi(dbport)

	if err != nil {
		fmt.Println("Cannot convert string to int")
	}

	db, err := database.ConnectToDB(host, port, username, password, dbname)

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	// HTTP Request for user
	// r.POST("/register", )

	// HTTP Request for Server
	r.GET("/servers", rest.GetAllServers(db))
	r.GET("/servers/:id", rest.GetServerWithId(db))
	r.POST("/servers", rest.CreateNewServer(db))
	r.PATCH("/servers/:id", rest.UpdateServer(db))
	r.DELETE("/servers/:id", rest.DeleteServer(db))

	r.Run(":8080")
}
