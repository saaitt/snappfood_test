package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/saaitt/snappfood_test/trips/domain"
)

func (hr *handler) UpdateTrip(c *gin.Context) {
	var req domain.TripUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(404, gin.H{
			"message": "wrong input",
		})
		return
	}
	trip, err := hr.uc.UpdateTrip(&req)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"trip": trip,
	})
}
