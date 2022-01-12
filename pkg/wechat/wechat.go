package wechat

import (
	"app/config"

	"github.com/silenceper/wechat/cache"
	wechat "github.com/silenceper/wechat/v2"
	wxConfig "github.com/silenceper/wechat/v2/miniprogram/config"
)

var (
	MiniprogramCfg *wxConfig.Config
	WeChat         *wechat.Wechat
)

func init() {
	var memory = cache.NewMemory()
	MiniprogramCfg = &wxConfig.Config{
		AppID:     config.GetString("miniprogram.appid"),
		AppSecret: config.GetString("miniprogram.appsecret"),
		Cache:     memory,
	}
	WeChat = wechat.NewWechat()
}
