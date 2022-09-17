package utils

import "github.com/gin-gonic/gin"

func RespondContentType(c *gin.Context, status int, body interface{}) {
	if c.GetHeader("accept") == "application/xml" {
		c.XML(status, body)
		return
	}
	c.JSON(status, body)
}

func RespondWithError(c *gin.Context, err *MsMvcError) {
	if c.GetHeader("accept") == "application/xml" {
		c.XML(err.StatusCode, err)
		return
	}
	c.JSON(err.StatusCode, err)
}
