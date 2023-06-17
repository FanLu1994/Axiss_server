package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"net/http"
	"time"
)

func GetWebsiteTitle(c *gin.Context) {
	url := c.Query("url")

	if url == "" {
		c.JSON(http.StatusBadRequest, APIResponse{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "URL parameter is missing",
			Result:       "",
		})
		return
	}

	// 使用Colly爬取网页标题
	collector := colly.NewCollector()
	collector.SetRequestTimeout(1 * time.Second)
	var title string

	collector.OnHTML("title", func(e *colly.HTMLElement) {
		title = e.Text
	})

	err := collector.Visit(url)
	if err != nil {
		c.JSON(http.StatusOK, APIResponse{
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: "Failed to fetch website data",
			Result:       "",
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		ErrorCode:    http.StatusOK,
		ErrorMessage: "",
		Result:       title,
	})
}
