package domain

import "github.com/gin-gonic/gin"

type Handler interface {
	CreateOrderHr(c *gin.Context)
	CreateOrderDelayReportHr(c *gin.Context)
	GetVendorDelayReportHr(c *gin.Context)
}
