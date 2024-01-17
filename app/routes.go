package app

import "github.com/gin-gonic/gin"

func (a *App) RegisterRoutes() {
	rg := a.engine.RouterGroup.Group("/")
	rg.GET("health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	orderRg := a.engine.RouterGroup.Group("order")
	orderRg.POST("/", a.orderHr.CreateOrderHr)
	orderRg.POST("/delay", a.orderHr.CreateOrderDelayReportHr)
}
