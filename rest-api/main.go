package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", getData)

	router.Run("127.0.0.1:3000")
}

func getData(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, map[string]string{
		"status": "Ok",
	})
}
