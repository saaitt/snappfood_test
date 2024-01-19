package delivery

import (
	"github.com/gin-gonic/gin"
)

func (hr *handler) GetVendorDelayReportHr(c *gin.Context) {
	report, err := hr.uc.VendorDelaysReport()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "there was an error getting the report",
			"error":   err,
		})
		return
	}
	c.JSON(200, gin.H{
		"report": report,
	})
}
