package router

import "github.com/laoyutang/laoyutang-server/controllers/test"

func initTestRouter() {
	testRouter := Router.Group("/test")
	{
		testRouter.Any("/ok", test.Test)
	}
}
