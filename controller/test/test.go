package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	c.String(http.StatusOK, "成功！")
}
