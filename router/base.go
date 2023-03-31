package router

import "github.com/laoyutang/laoyutang-server/controllers/base"

func initBaseRouter() {
	Router.POST("/sign-in", base.SignIn)
	Router.POST("/login", base.Login)
}
