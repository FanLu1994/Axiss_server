package api

import (
	"Axiss_server/db"
	"Axiss_server/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

var Logger = util.Logger

type BenchMarkTO struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	Icon string `json:"icon"`
	Tags string `json:"tags"`
}

func AddBenchMark(c *gin.Context) {
	var benchmark BenchMarkTO
	if err := c.ShouldBindJSON(&benchmark); err != nil {
		c.JSON(http.StatusOK, APIResponse{
			ErrorCode:    400,
			ErrorMessage: "参数错误",
		})
		return
	}

	newBenchmark := db.BenchMark{
		Name: benchmark.Name,
		Url:  benchmark.Url,
		Icon: benchmark.Icon,
		Tags: benchmark.Tags,
	}

	db.GlobalDb.Create(&newBenchmark)

	c.JSON(http.StatusOK, APIResponse{
		ErrorCode:    http.StatusOK,
		ErrorMessage: "",
	})
}

func GetBenchmarkList(c *gin.Context) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")
	tag := c.Query("tag")
	page, err1 := strconv.Atoi(pageStr)
	size, err2 := strconv.Atoi(sizeStr)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page or size value"})
		return
	}

	offset := (page - 1) * size
	limit := size

	var benchmarks []db.BenchMark
	// 查询数据
	db.GlobalDb.Where("tags LIKE ?", "%"+tag+"%").Offset(offset).Limit(limit).Find(&benchmarks)

	c.JSON(http.StatusOK, gin.H{"data": benchmarks})
}

func GetTags(c *gin.Context) {
	var records []db.BenchMark
	result := db.GlobalDb.Find(&records)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "数据库查询失败"})
		return
	}

	// 创建一个映射来存储不重复的标签
	uniqueTags := make(map[string]bool)
	// 遍历查询结果并提取标签字段的值
	for _, record := range records {
		tagsString := record.Tags
		tagsString = "标签1 标签2 标签3 标签1 标签4"

		// 使用空格分割标签字符串
		tagsSlice := strings.Split(tagsString, " ")
		// 将标签添加到映射中
		for _, tag := range tagsSlice {
			uniqueTags[tag] = true
		}
	}

	// 从映射中提取不重复的标签
	var uniqueTagList []string
	for tag := range uniqueTags {
		uniqueTagList = append(uniqueTagList, tag)
	}

	// 返回不重复的标签列表给客户端
	c.JSON(200, gin.H{"tags": uniqueTagList})
}
