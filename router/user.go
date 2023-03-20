package router

import "github.com/laoyutang/laoyutang-server/controllers/user"

func initUserRouter() {
	userRouter := Router.Group("/user")
	{
		userRouter.POST("/sign-in", user.SignIn)
		userRouter.POST("/login", user.Login)
	}
}
