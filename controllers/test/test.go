package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laoyutang/laoyutang-server/modules/structs"
	"github.com/laoyutang/laoyutang-server/utils"
)

func Test(c *gin.Context) {
	type test struct {
		A string `json:"a"`
	}
	str := test{}
	utils.ParseJson("{\"a\":\"sdfsdf胜多负少\"}", &str)

	res := &structs.Response{
		Success: true,
		Data: &gin.H{
			"test": c.Request.Method,
		},
	}
	c.JSON(http.StatusOK, res)
}
