package domain

import "github.com/gin-gonic/gin"

type Handler interface {
	AssignTrip(c *gin.Context)
	UpdateTrip(c *gin.Context)
}
