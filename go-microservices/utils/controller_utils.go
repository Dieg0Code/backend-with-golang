package utils

import "github.com/gin-gonic/gin"

func Respong(c *gin.Context, status int, body interface{}) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(status, body)
		return
	}

	c.JSON(status, body)
}

func RespondError(c *gin.Context, err *ApplicationError) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(err.Status, err)
		return
	}

	c.JSON(err.Status, err)
}
