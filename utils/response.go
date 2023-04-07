package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laoyutang/laoyutang-server/modules/structs"
)

// 失败响应
func ResponseFail(c *gin.Context, errNo int, message string) {
	c.JSON(http.StatusOK, &structs.Response{
		Success: false,
		ErrNo:   errNo,
		Message: message,
	})
}

// 默认失败响应
func ResponseFailDefault(c *gin.Context) {
	c.JSON(http.StatusOK, &structs.Response{
		Success: false,
		ErrNo:   http.StatusInternalServerError,
		Message: "服务器异常",
	})
}

// 成功响应
func ResponseSuccess(c *gin.Context, data any) {
	c.JSON(http.StatusOK, &structs.Response{
		Success: true,
		Data:    data,
		Message: "操作成功",
	})
}
