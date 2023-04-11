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
		userRouter.GET("", middlewares.FuncAuth("user:operation"), user.Read)
		userRouter.POST("/get-menus", user.GetMenusAndPerms)
	}

	roleRouter := adminRouter.Group("/role")
	{
		roleRouter.GET("", middlewares.FuncAuth("role:operation"), role.Read)
		roleRouter.POST("", middlewares.FuncAuth("role:operation"), role.Create)
		roleRouter.PUT("", middlewares.FuncAuth("role:operation"), role.Update)
		roleRouter.DELETE("", middlewares.FuncAuth("role:operation"), role.Delete)
		// roleRouter.GET("/list", role.GetList)
	}

	menuRouter := adminRouter.Group("/menu")
	{
		menuRouter.GET("", middlewares.FuncAuth("menu:operation"), menu.Read)
		menuRouter.POST("", middlewares.FuncAuth("menu:operation"), menu.Create)
		menuRouter.PUT("", middlewares.FuncAuth("menu:operation"), menu.Update)
		menuRouter.DELETE("", middlewares.FuncAuth("menu:operation"), menu.Delete)
		// menuRouter.GET("/list", menu.GetList)
	}
}
