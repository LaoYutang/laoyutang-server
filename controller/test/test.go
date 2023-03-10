package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laoyutang/laoyutang-server/modules/structs"
)

func Test(c *gin.Context) {
	res := &structs.Response{
		Success: true,
		Data: &gin.H{
			"test": c.Request.Method,
		},
	}
	c.JSON(http.StatusOK, res)
}
