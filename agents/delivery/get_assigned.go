package delivery

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func (hr *handler) GetAssigned(c *gin.Context) {
	id := c.Param("id")
	agentID, err := strconv.ParseUint(id, 0, 64)
	if err != nil {
		c.JSON(403, gin.H{
			"message": "invalid input",
			"error":   err.Error(),
		})
		return
	}
	task, err := hr.uc.GetAssigned(agentID)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "there was an unexpected error assigning request",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"task": task,
	})
}
