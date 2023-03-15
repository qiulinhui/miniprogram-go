package wechat

import (
	"app/app"

	"github.com/silenceper/wechat/cache"
	wechat "github.com/silenceper/wechat/v2"
	wxConfig "github.com/silenceper/wechat/v2/miniprogram/config"
)

func GetWechat() *wechat.Wechat {
	return wechat.NewWechat()
}

func GetMiniprogramCfg() *wxConfig.Config {
	var memory = cache.NewMemory()
	return &wxConfig.Config{
		AppID:     app.Config.GetString("miniprogram.appid"),
		AppSecret: app.Config.GetString("miniprogram.appsecret"),
		Cache:     memory,
	}
}
