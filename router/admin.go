package router

import (
	"github.com/laoyutang/laoyutang-server/controllers/admin/user"
	"github.com/laoyutang/laoyutang-server/middlewares"
)

func initAdminRouter() {
	adminRouter := Router.Group("/admin")

	userRouter := adminRouter.Group("/user")
	{
		userRouter.POST("/sign-in", user.SignIn)
		userRouter.POST("/login", user.Login)
		userRouter.POST("/get-menus", user.GetMenus).Use(middlewares.LoginAuth())
	}
}
