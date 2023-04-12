package router

import (
	"github.com/laoyutang/laoyutang-server/controllers/prompter/aiApiKeys"
	"github.com/laoyutang/laoyutang-server/middlewares"
)

func initPrompterRouter() {
	prompterRouter := Router.Group("/prompter")
	prompterRouter.Use(middlewares.LoginAuth())

	keysRouter := prompterRouter.Group("/aiApiKeys")
	{
		keysRouter.GET("", middlewares.FuncAuth("aiApiKeys"), aiApiKeys.Read)
		keysRouter.POST("", middlewares.FuncAuth("aiApiKeys:operation"), aiApiKeys.Create)
		keysRouter.DELETE("", middlewares.FuncAuth("aiApiKeys:operation"), aiApiKeys.Delete)
	}
}
