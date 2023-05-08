package main

import (
	"b5gocmf/router"
	"b5gocmf/utils/core"
	"fmt"
)

// fresh
//
//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download
func main() {
	//初始化引导
	engine := core.Load()
	fmt.Println("哈哈哈啊哈哈")
	//加载路由
	router.LoadRouter(engine)

	//启动服务
	core.Run(engine)
}
