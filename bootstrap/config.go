package bootstrap

import (
	"app/app"
	"fmt"

	"github.com/spf13/viper"
)

type StrMap map[string]interface{}

func InitConfig() {

	conf := viper.New()
	// 设置配置文件名
	conf.SetConfigName("config")
	// 设置配置文件类型
	conf.SetConfigType("yaml")
	// 查找配置文件的路径
	conf.AddConfigPath(".")
	err := conf.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败：%w", err))
	}
	app.Viper = conf
}
