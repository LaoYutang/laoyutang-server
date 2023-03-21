package router

import (
	"github.com/laoyutang/laoyutang-server/controllers/test"
	"github.com/laoyutang/laoyutang-server/middlewares"
)

func initTestRouter() {
	testRouter := Router.Group("/test").Use(middlewares.LoginAuth())
	{
		testRouter.Any("/ok", test.Test)
	}
}
