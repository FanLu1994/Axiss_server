package api

import (
	"Axiss_server/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Feed struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

func AddFeed(c *gin.Context) {
	var feed Feed
	if err := c.ShouldBindJSON(&feed); err != nil {
		c.JSON(http.StatusOK, APIResponse{
			ErrorCode:    400,
			ErrorMessage: "参数错误",
		})
		return
	}

	newFeed := db.Feed{
		Title:       feed.Title,
		Link:        feed.Link,
		Description: feed.Description,
		Author:      feed.Author,
	}

	db.GlobalDb.Create(newFeed)

	c.JSON(http.StatusOK, APIResponse{
		ErrorCode:    http.StatusOK,
		ErrorMessage: "",
	})
}
