package model

import (
	"bookstore/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	var err error

	config := mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
			config.GetString("db.username", "root"),
			config.GetString("db.password", ""),
			config.GetString("db.hostname", "127.0.0.1"),
			config.GetString("db.port", "3306"),
			config.GetString("db.database", ""),
			config.GetString("db.charset", "utf8"),
		),
	})

	// 准备数据库连接池
	DB, err = gorm.Open(config)
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}
	return DB
}
