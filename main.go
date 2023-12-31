package main

import (
	"Axiss_server/api"
	"Axiss_server/config"
	"Axiss_server/util"
	"github.com/gin-gonic/gin"
)

func main() {
	globalConfig := config.GetGlobalConfig()

	router := gin.Default()
	// 处理跨域请求
	router.Use(api.CORSMiddleware())
	// 日志模块
	router.Use(util.MyLogger())

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

	benchmark := router.Group("benchmark")
	benchmark.Use(api.CORSMiddleware())
	{
		benchmark.POST("/add", api.AddBenchMark)
		benchmark.GET("/get", api.GetBenchmarkList)
		benchmark.GET("/tags", api.GetTags)
	}

	router.Run(globalConfig.Listen)
}
