package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laoyutang/laoyutang-server/modules/structs"
	"github.com/sirupsen/logrus"
)

func Test(c *gin.Context) {
	test := map[string]interface{}{
		"a": 1,
	}
	logrus.Debug(test)

	res := &structs.Response{
		Success: true,
		Data: &gin.H{
			"test": c.Request.Method,
		},
	}
	c.JSON(http.StatusOK, res)
}
