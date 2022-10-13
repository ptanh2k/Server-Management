package main

import (
	"fmt"
	"log"
	"os"

	"sm/internal/database"

	"sm/api/rest"

	"sm/middleware"

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
	auth := r.Group("/auth")
	{
		auth.POST("/register", rest.Register(db))
		auth.POST("/login", rest.Login(db))
	}
	// Admin
	protected := r.Group("/admin")
	{
		protected.Use(middleware.JwtCheckMiddleware())
		protected.GET("/user", rest.CurrentUser(db))
	}
	// HTTP Request for Server

	server := r.Group("/servers")
	{
		server.GET("/all", middleware.JwtCheckMiddleware(), rest.GetAllServers(db))
		server.GET("/:id", middleware.JwtCheckMiddleware(), rest.GetServerWithId(db))
		server.POST("/", middleware.JwtCheckMiddleware(), rest.CreateNewServer(db))
		server.PATCH("/:id", middleware.JwtCheckMiddleware(), rest.UpdateServer(db))
		server.DELETE("/:id", middleware.JwtCheckMiddleware(), rest.DeleteServer(db))
	}
	r.Run(":8080")
}
