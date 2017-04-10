package example

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func Example() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Authorization") == "" {
			c.JSON(http.StatusUnauthorized, response{Message: "This route is protected by example middleware..", Status: 4001})
			c.Abort()
			return
		}
	}
}
