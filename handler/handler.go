package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Format(c *gin.Context) {
	var (
		req  FormatRequest
		resp FormatResponse
	)

	err := c.BindJSON(&req)
	if err != nil {
		c.Error(err)
		return
	}

	resp.Content = req.Content

	c.JSON(http.StatusOK, resp)
}
