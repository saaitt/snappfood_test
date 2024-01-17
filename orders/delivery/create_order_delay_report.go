package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/saaitt/snappfood_test/orders/domain"
)

func (hr *handler) CreateOrderDelayReportHr(c *gin.Context) {
	var req domain.OrderDelayReportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(404, gin.H{
			"message": "wrong input",
		})
		return
	}
	order, err := hr.uc.CreateOrderDelayReport(&req)
	if err != nil {
		if err == domain.ErrorOrderDelayReportAlreadyExists ||
			err == domain.ErrorOrderIsOnTheWay {
			c.JSON(409, gin.H{
				"message": err.Error(),
			})
			return
		}
		if err == domain.ErrorWrongOrderID {
			c.JSON(422, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(500, gin.H{
			"message": "there was an unexpected error submitting the order delay report",
		})
		return
	}
	c.JSON(200, gin.H{
		"order": order,
	})
}
