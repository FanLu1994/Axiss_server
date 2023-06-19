package main

import (
	"Axiss_server/api"
	"Axiss_server/wechat"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 处理跨域请求
	router.Use(api.CORSMiddleware())

	md := router.Group("/md")
	{
		md.GET("/get_website_title", api.GetWebsiteTitle)
	}

	wechatInstance := wechat.InitWechat()
	offAccount := wechat.NewMyOfficialAccount(wechatInstance)
	wc := router.Group("/wc")
	{
		wc.Any("/wc/serve", offAccount.Serve)
	}

	router.Run(":7777")
}
