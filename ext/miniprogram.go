package ext

import (
	"bookstore/config"

	"github.com/silenceper/wechat/cache"
	wechat "github.com/silenceper/wechat/v2"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
)

var memory = cache.NewMemory()
var MiniprogramCfg = &miniConfig.Config{

	AppID:     config.GetString("miniprogram.appid"),
	AppSecret: config.GetString("miniprogram.appsecret"),
	Cache:     memory,
}

func WeChat() *wechat.Wechat {
	wc := wechat.NewWechat()
	return wc
}
