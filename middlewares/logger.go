package middlewares

import (
	"bytes"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		body, _ := io.ReadAll(c.Request.Body)
		logrus.WithFields(logrus.Fields{
			"Ip":    c.ClientIP(),
			"Path":  c.Request.URL.Path,
			"Query": c.Request.URL.RawQuery,
			"Body":  string(body),
		}).Info("Request")
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		c.Next()
	}
}
