package router

import (
	"github.com/laoyutang/laoyutang-server/controllers/admin/menu"
	"github.com/laoyutang/laoyutang-server/controllers/admin/role"
	"github.com/laoyutang/laoyutang-server/controllers/admin/user"
	"github.com/laoyutang/laoyutang-server/middlewares"
)

func initAdminRouter() {
	adminRouter := Router.Group("/admin")
	adminRouter.Use(middlewares.LoginAuth())

	userRouter := adminRouter.Group("/user")
	{
		userRouter.GET("", user.Read)
		userRouter.POST("/get-menus", user.GetMenus)
	}

	roleRouter := adminRouter.Group("/role")
	{
		roleRouter.GET("", role.Read)
	}

	menuRouter := adminRouter.Group("/menu")
	{
		menuRouter.GET("", menu.Read)
	}
}
