package lib

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const RequestBodyKey = "request_body_for_sentry"

func CaptureRequestBody() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip body capture for GET, HEAD, and OPTIONS requests
		if c.Request.Method == http.MethodGet ||
			c.Request.Method == http.MethodHead ||
			c.Request.Method == http.MethodOptions {
			c.Next()
			return
		}

		// Check content type
		contentType := c.GetHeader("Content-Type")

		// Only process specific content types
		if !strings.Contains(contentType, "application/json") &&
			!strings.Contains(contentType, "application/x-www-form-urlencoded") &&
			!strings.Contains(contentType, "multipart/form-data") {
			c.Next()
			return
		}

		// Read the request body
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.Next()
			return
		}

		// Restore the request body
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// Process the body based on content type
		var bodyData interface{}
		if strings.Contains(contentType, "application/json") {
			// Parse JSON
			if err := json.Unmarshal(bodyBytes, &bodyData); err == nil {
				c.Set(RequestBodyKey, bodyData)
			} else {
				c.Set(RequestBodyKey, string(bodyBytes))
			}
		} else if strings.Contains(contentType, "multipart/form-data") {
			// For multipart form data, store form values
			if form, err := c.MultipartForm(); err == nil {
				c.Set(RequestBodyKey, form.Value)
			}
		} else {
			// For other content types, store as string
			c.Set(RequestBodyKey, string(bodyBytes))
		}

		c.Next()
	}
}
