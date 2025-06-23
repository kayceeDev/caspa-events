package lib

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// JSON serializes the api response properly to json
func JSON(c *gin.Context, message string, status int, data any, cl *CustomLogger, request ...interface{}) {
	var responseData gin.H
	switch v := data.(type) {
	case error:
		responseData = gin.H{
			"message": message,
			"errors":  v.Error(),
			"status":  http.StatusText(status),
		}

		if cl != nil {
			cl.Error(c, message, v.Error())
		}

	default:
		responseData = gin.H{
			"message": message,
			"data":    v,
			"status":  http.StatusText(status),
		}

		// if cl != nil {
		// 	cl.Info(message, v)
		// }
	}

	if val, ok := c.Request.Context().Deadline(); ok {
		if val.Before(time.Now()) {
			c.Abort()
			return
		}
	}

	c.JSON(status, responseData)
}

func HtmlRes(c *gin.Context, template string, status int, data gin.H) {
	c.HTML(status, template, data)
}
