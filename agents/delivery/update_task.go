package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/saaitt/snappfood_test/agents/domain"
)

func (hr *handler) UpdateTask(c *gin.Context) {
	var req domain.TaskUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(404, gin.H{
			"message": "wrong input",
		})
		return
	}
	trip, err := hr.uc.UpdateTask(&req)
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
