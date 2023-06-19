package wechat

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
)

func InitWechat() *wechat.Wechat {
	wc := wechat.NewWechat()
	memoryCache := cache.NewMemory()
	wc.SetCache(memoryCache)
	return wc
}
