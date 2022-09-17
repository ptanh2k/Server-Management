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

// Get all servers info
func GetAllServers(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		servers := make([]APIServer, 0)

		db.Table("servers").Model(&s.Server{}).Find(&servers, &APIServer{})

		c.IndentedJSON(http.StatusOK, gin.H{"servers": servers})
	}
	return gin.HandlerFunc(fn)
}

// Add new server
func CreateNewServer(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		type CreateServerInput struct {
			Name     string `json:"name"`
			Ip       string `json:"ip"`
			Port     uint16 `json:"port"`
			Status   bool   `json:"status"`
			Password string `json:"password"`
		}

		var input CreateServerInput

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
