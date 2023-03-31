package router

import (
	"github.com/laoyutang/laoyutang-server/controllers/admin/user"
	"github.com/laoyutang/laoyutang-server/middlewares"
)

func initAdminRouter() {
	adminRouter := Router.Group("/admin")

	userRouter := adminRouter.Group("/user")
	userRouter.Use(middlewares.LoginAuth())
	{
		userRouter.POST("/get-menus", user.GetMenus)
	}
}
