package main

import (
	"Axiss_server/api"
	"Axiss_server/config"
	"github.com/gin-gonic/gin"
)

func main() {
	globalConfig := config.GetGlobalConfig()

	router := gin.Default()
	// 处理跨域请求
	router.Use(api.CORSMiddleware())

	md := router.Group("/md")
	md.Use(api.CORSMiddleware())
	{
		md.GET("/get_website_title", api.GetWebsiteTitle)
	}

	//wechatInstance := wechat.InitWechat()
	//offAccount := wechat.NewMyOfficialAccount(wechatInstance)
	//wc := router.Group("/wc")
	//{
	//	wc.Any("/serve", offAccount.Serve)
	//}

	rss := router.Group("rss")
	rss.Use(api.CORSMiddleware())
	{
		rss.POST("/add", api.AddFeed)
		rss.GET("/feed", api.GetFeed)
	}

	router.Run(globalConfig.Listen)
}
