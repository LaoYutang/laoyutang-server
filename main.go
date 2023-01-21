package main

import (
	"github.com/laoyutang/laoyutang-server/router"
)

func main() {
	// 启动gin监听9000端口
	router.Router.Run(":9000")
}
