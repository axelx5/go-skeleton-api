package rest

import "github.com/gin-gonic/gin"

func Start() {
	r := gin.Default()

	wiredHandlers := wireHandlersDependencies()
	bindRoutes(r, wiredHandlers)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
