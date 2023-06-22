package api

import (
	"Axiss_server/db"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/feeds"
	"net/http"
	"time"
)

type Feed struct {
	Title       string    `json:"title"`
	Link        string    `json:"link"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
	Created     time.Time `json:"-"`
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

	db.GlobalDb.Create(&newFeed)

	c.JSON(http.StatusOK, APIResponse{
		ErrorCode:    http.StatusOK,
		ErrorMessage: "",
	})
}

func GetFeed(c *gin.Context) {
	var feedList []db.Feed
	db.GlobalDb.Order("id desc").Limit(100).Find(&feedList)
	feedXml := &feeds.Feed{
		Title:       "Axiss的收藏",
		Link:        &feeds.Link{Href: "https://fanlu.top/"},
		Description: "Axiss的浏览器收藏",
		Author:      &feeds.Author{Name: "Axiss", Email: "1066580119@qq.com"},
		Created:     time.Now(),
		Items:       []*feeds.Item{},
	}
	for _, feed := range feedList {
		item := feeds.Item{
			Title:       feed.Title,
			Link:        &feeds.Link{Href: feed.Link},
			Description: feed.Description,
			Author:      &feeds.Author{Name: feed.Author, Email: ""},
			Created:     feed.CreatedAt,
		}
		feedXml.Items = append(feedXml.Items, &item)
	}

	rss, err := feedXml.ToRss()
	if err != nil {
		Fail(c, 400, "格式转换失败", err)
		return
	}

	c.Header("Content-Type", "application/xml")
	c.String(http.StatusOK, rss)

}
