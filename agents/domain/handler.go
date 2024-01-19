package domain

import "github.com/gin-gonic/gin"

type Handler interface {
	GetAssigned(c *gin.Context)
	UpdateTask(c *gin.Context)
}
