package users

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/models/user"
)

func FetchUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"payload": user.FetchUsers(), "status": http.StatusOK})
}
