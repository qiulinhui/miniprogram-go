package app

import (
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type config struct {
}

var Viper *viper.Viper

func newConfig() *config {
	return &config{}
}

var Config = newConfig()

// Get 获取配置项
func (conf *config) Get(path string, defaultValue ...interface{}) interface{} {
	if !Viper.IsSet(path) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return Viper.Get(path)
}

func (conf *config) GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(conf.Get(path, defaultValue...))
}

// GetString 获取 String 类型的配置信息
func (conf *config) GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(conf.Get(path, defaultValue...))
}

// GetInt64 获取 Int64 类型的配置信息
func (conf *config) GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(conf.Get(path, defaultValue...))
}

// GetUint 获取 Uint 类型的配置信息
func (conf *config) GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(conf.Get(path, defaultValue...))
}

// GetBool 获取 Bool 类型的配置信息
func (conf *config) GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(conf.Get(path, defaultValue...))
}
