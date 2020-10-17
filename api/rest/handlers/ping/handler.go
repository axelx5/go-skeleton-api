package ping

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (Handler) Handle(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
