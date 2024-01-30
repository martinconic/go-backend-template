package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStatusInfo(c *gin.Context) {
	err := server.db.GetDBStatus()
	if err != nil {
		log.Fatalf("Unable to reach database: %v", err)
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "Database not found",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": "Server and database up and running",
	})

}
