package rest

import (
	"fmt"
	s "sm/internal/model"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type APIServer struct {
	Id     uint16 `json:"id,omitempty"`
	Name   string `json:"name"`
	Ip     string `json:"ip"`
	Port   uint16 `json:"port"`
	Status bool   `json:"status"`
}

type ServerInput struct {
	Name     string `json:"name"`
	Ip       string `json:"ip"`
	Port     uint16 `json:"port"`
	Status   bool   `json:"status"`
	Password string `json:"password"`
}

// GET /servers
// Get all servers info
func GetAllServers(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		servers := make([]APIServer, 0)

		db.Table("servers").Find(&servers, &APIServer{})

		c.IndentedJSON(http.StatusOK, servers)
	}
	return gin.HandlerFunc(fn)
}

// GET /servers/{id}
// Get server with id
func GetServerWithId(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		sid := c.Param("id")

		var servers []s.Server
		var server APIServer

		db.Table("servers").Where("id = ?", sid).Find(&servers).Scan(&server)

		c.IndentedJSON(http.StatusOK, server)
	}

	return gin.HandlerFunc(fn)
}

// POST /servers
// Add new server
func CreateNewServer(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var input ServerInput

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newServer := s.Server{Name: input.Name, Ip: input.Ip, Port: input.Port, Status: input.Status, Password: input.Password}

		result := db.Table("servers").Select("Name", "Ip", "Port", "Status", "Password").Create(&newServer)

		if err := result.Error; err != nil {
			panic(err)
		}

		fmt.Printf("%d row(s) affected\n", result.RowsAffected)

		c.JSON(http.StatusCreated, gin.H{"server": newServer})

	}

	return gin.HandlerFunc(fn)
}

// PATCH /servers/{id}
// Update to a single server in the system
func UpdateServer(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		sid := c.Param("id")

		var servers []s.Server

		find := db.Table("servers").Where("id = ?", sid).Find(&servers)

		if err := find.Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Server not found"})
			return
		}

		var input ServerInput

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updateServer := s.Server{Name: input.Name, Ip: input.Ip, Port: input.Port, Status: input.Status, Password: input.Password}

		result := db.Table("servers").Model(&s.Server{}).Where("id = ?", sid).Updates(&updateServer)

		if err := result.Error; err != nil {
			panic(err)
		}

		fmt.Printf("%d row(s) affected\n", result.RowsAffected)

		c.JSON(http.StatusOK, gin.H{"server": updateServer})
	}

	return gin.HandlerFunc(fn)
}

// DELETE /servers/{id}
// Delete server from system
func DeleteServer(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		sid := c.Param("id")

		result := db.Table("servers").Where("id = ?", sid).Delete(&s.Server{})

		fmt.Printf("%d row(s) affected\n", result.RowsAffected)

		c.JSON(http.StatusOK, gin.H{"data": "deleted"})
	}

	return gin.HandlerFunc(fn)
}

// Insert multiple servers
// func CreateMultipleServers(db *gorm.DB) gin.HandlerFunc {
// 	fn := func(c *gin.Context) {
// 		var input []APIServer

// 		if err := c.ShouldBindJSON(&input); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err})
// 		}

// 		var inputToDB []APIServer

// 		for i := 0; i < len(input); i++ {
// 			info := APIServer{Id: input[i].Id, Name: input[i].Name, Ip: input[i].Ip, Port: input[i].Port, Status: input[i].Status}
// 			inputToDB = append(inputToDB, info)
// 		}
// 	}

// 	return gin.HandlerFunc(fn)
//}
