package main

import (
	"app/bootstrap"
)

func main() {
	bootstrap.SetUpDB()    // 初始化数据库
	bootstrap.SetUpRoute() // 初始化路由
	bootstrap.Start()      // 启动服务
}
