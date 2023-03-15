package main

import (
	"app/bootstrap"
)

func main() {
	bootstrap.InitConfig() // 初始化配置文件
	bootstrap.InItDB()     // 初始化数据库连接
	bootstrap.InitRoutes() // 初始化路由
	bootstrap.InItGin()    // 初始化Gin
}
