package router

import (
	"github.com/laoyutang/laoyutang-server/controllers/admin/user"
	"github.com/laoyutang/laoyutang-server/middlewares"
)

func initUserRouter() {
	userRouter := Router.Group("/user")
	{
		userRouter.POST("/sign-in", user.SignIn)
		userRouter.POST("/login", user.Login)
		userRouter.POST("/get-menus", user.GetMenus).Use(middlewares.LoginAuth())
	}
}
