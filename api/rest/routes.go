package rest

import (
	"github.com/gin-gonic/gin"
)

func bindRoutes(router *gin.Engine, handlers *HandlerRegistry) {
	router.GET("/ping", handlers.PingHandler.Handle)

	account := router.Group("/account")
	account.POST("", handlers.CreateAccountHandler.Handle)
}
