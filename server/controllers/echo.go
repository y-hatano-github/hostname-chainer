package controllers

import "github.com/gin-gonic/gin"

func Echo(c *gin.Context) {
	host := c.MustGet("host").(string)
	c.JSON(200, gin.H{"message": host})
}
