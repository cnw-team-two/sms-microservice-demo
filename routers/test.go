package routers

import (
	"github.com/gin-gonic/gin"
)

func TestGET(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "UP", "cluster": "Edge",
	})
}
