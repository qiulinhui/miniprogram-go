package bootstrap

import (
	"bookstore/config"
	"bookstore/model"
	"time"
)

func SetUpDB() {
	// 建立数据库连接池
	db := model.ConnectDB()

	sqlDB, _ := db.DB()

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(config.GetInt("db.mysql.max_open_connections"))
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(config.GetInt("db.mysql.max_idle_connections"))
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("db.mysql.max_life_seconds")) * time.Second)
}
