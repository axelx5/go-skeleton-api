package account

import (
	"github.com/axelx5/go-skeleton-api/application/usecases"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	CreateUseCase usecases.UseCase
}

func (h Handler) Handle(c *gin.Context) {

	h.CreateUseCase.Execute()

	c.JSON(200, gin.H{
		"message": "pon111g",
	})
}
