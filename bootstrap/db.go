package bootstrap

import (
	"app/app"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InItDB() {
	var err error
	// 建立数据库连接池
	conf := mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
			app.Config.GetString("db.username", "root"),
			app.Config.GetString("db.password", ""),
			app.Config.GetString("db.hostname", "127.0.0.1"),
			app.Config.GetString("db.port", "3306"),
			app.Config.GetString("db.database", ""),
			app.Config.GetString("db.charset", "utf8"),
		),
	})

	// 准备数据库连接池
	db, err := gorm.Open(conf)
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}

	sqlDB, _ := db.DB()

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(app.Config.GetInt("db.mysql.max_open_connections"))
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(app.Config.GetInt("db.mysql.max_idle_connections"))
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(time.Duration(app.Config.GetInt("db.mysql.max_life_seconds")) * time.Second)
	app.DB = db
}
