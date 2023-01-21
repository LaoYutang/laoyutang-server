package router

import "github.com/laoyutang/laoyutang-server/controller/test"

func initTestRouter() {
	testRouter := Router.Group("/test")
	{
		testRouter.GET("/ok", test.Test)
	}
}
