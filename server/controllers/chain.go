package controllers

import (
	"net/http"

	"hostname-chainer/client"
	"hostname-chainer/types"

	"github.com/gin-gonic/gin"
)

func Chain(c *gin.Context) {
	currentHost := c.MustGet("currentHost").(types.CurrentHost)

	var requestBody client.Hosts
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
	}

	if !currentHost.HasNextHostInfo() {
		resBody := requestBody.AddCurrentHostInfo(currentHost.GetHostName(), "")
		c.JSON(http.StatusOK, resBody)
		return
	}

	requester := client.NewClient(currentHost.GetNextHostAddress(), currentHost.GetNextHostPort())
	reqBody := requestBody.AddCurrentHostInfo(currentHost.GetHostName(), currentHost.GetNextHostAddress())
	resp, err := requester.Post(reqBody)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}
