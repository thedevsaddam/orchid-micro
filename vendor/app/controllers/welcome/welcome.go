package welcome

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to Orchid-micro for GO lovers", "status": 2000})
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "You are requesting for some api..", "status": 2000})
}
