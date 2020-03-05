package controllers

import (
	"github.com/gin-gonic/gin"
	"recursive-echo/types"
)

func Echo(c *gin.Context) {
	host := c.MustGet("host").(types.Host)

	if host.HasNextHostInfo() {
		resp, _ := host.GetNextHostClient().RequestToNextHost()
		c.JSON(200, host.GetName() + " " + string(resp.Body()))
	} else {
		c.JSON(200, host.GetName())
	}
}
