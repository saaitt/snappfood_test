package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (hr *handler) GetAssigned(c *gin.Context) {
	fmt.Println("hi")
	agentHistory, err := hr.uc.GetAssigned()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "there was an unexpected error assigning request",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"agent_history": agentHistory,
	})
}
