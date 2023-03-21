package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laoyutang/laoyutang-server/modules/structs"
	"github.com/laoyutang/laoyutang-server/utils"
	"github.com/sirupsen/logrus"
)

func Test(c *gin.Context) {
	user := c.MustGet("UserName")

	logrus.Debug(utils.ToJson(user))

	res := &structs.Response{
		Success: true,
		Data: &gin.H{
			"test": c.Request.Method,
		},
	}
	c.JSON(http.StatusOK, res)
}
