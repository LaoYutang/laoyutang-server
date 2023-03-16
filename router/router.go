package router

import (
	"github.com/gin-gonic/gin"
	"github.com/laoyutang/laoyutang-server/middlewares"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
	// 注册中间件
	Router.Use(middlewares.Logger())

	// 初始化路由组
	initTestRouter()
	initUserRouter()
}
