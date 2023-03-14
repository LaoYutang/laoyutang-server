package main

import (
	_ "github.com/laoyutang/laoyutang-server/modules/db"
	_ "github.com/laoyutang/laoyutang-server/modules/loger"
	"github.com/laoyutang/laoyutang-server/router"
)

func main() {
	// 启动gin监听9000端口
	router.Router.Run(":9000")
}
