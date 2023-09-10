package api

import (
	"Axiss_server/db"
	"Axiss_server/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var Logger = util.Logger

type BenchMarkTO struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	Icon string `json:"icon"`
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
	page, err1 := strconv.Atoi(pageStr)
	size, err2 := strconv.Atoi(sizeStr)
	Logger.Errorln("testtesttest")
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page or size value"})
		return
	}

	offset := (page - 1) * size
	limit := size

	var benchmarks []db.BenchMark
	// 查询数据
	db.GlobalDb.Offset(offset).Limit(limit).Find(&benchmarks)

	c.JSON(http.StatusOK, gin.H{"data": benchmarks})
}
