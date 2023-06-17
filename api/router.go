package api

import "github.com/gin-gonic/gin"

func router() {
	router := gin.Default()

	md := router.Group("/md")
	{
		md.GET("/get_website_title", GetWebsiteTitle)
	}

	router.Run(":8080")
}
