package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/saaitt/snappfood_test/trips/domain"
)

func (hr *handler) AssignTrip(c *gin.Context) {
	var req domain.TripRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(404, gin.H{
			"message": "wrong input",
		})
		return
	}
	trip, err := hr.uc.AssignTrip(&req)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "there was an unexpected error assigning the trip",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"trip": trip,
	})
}
