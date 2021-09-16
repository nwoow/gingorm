package Router

import (
	controllers "DT/Controllers"
	"DT/Middleware"

	"github.com/gin-gonic/gin"
)

// Router ...
func Router() {
	Router := gin.Default()
	Router.Use(Middleware.TokenAuthMiddleware())
	r := Router.Group("/api")
	// Session Cotroll API end
	r.GET("/logs", controllers.Listerror) // new
	// r.GET("/logconfiglist", controllers.Logconfiglist) // new
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	Router.Run(":8080")
}
