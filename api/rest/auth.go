package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"sm/internal/model"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var input RegisterInput

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		u := model.User{}

		u.Username = input.Username
		u.Password = input.Password

		_, err := u.Register(db)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	return gin.HandlerFunc(fn)
}
