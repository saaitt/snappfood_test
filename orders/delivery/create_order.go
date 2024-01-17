package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/saaitt/snappfood_test/orders/domain"
)

func (hr *handler) CreateOrderHr(c *gin.Context) {
	var orderReq domain.OrderRequest
	if err := c.ShouldBindJSON(&orderReq); err != nil {
		c.JSON(404, gin.H{
			"message": "wrong input",
		})
		return
	}
	order, err := hr.uc.CreateOrder(&orderReq)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "there was an error creating the order",
		})
		return
	}
	c.JSON(200, gin.H{
		"order": order,
	})
}
