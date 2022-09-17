package rest

import (
	s "sm/internal/model"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllServers(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		servers := make([]s.Server, 0)

		db.Table("servers").Find(&servers)

		c.IndentedJSON(http.StatusOK, gin.H{"servers": servers})
	}
	return gin.HandlerFunc(fn)
}
