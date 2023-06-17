package main

import (
	"Axiss_server/api"
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

	router.Run(":7777")
}
