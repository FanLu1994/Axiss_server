package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type APIResponse struct {
	ErrorCode    int         `json:"code"`
	ErrorMessage string      `json:"message"`
	Result       interface{} `json:"result"`
}

func Success(c *gin.Context, message string, result interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		ErrorCode:    200,
		ErrorMessage: message,
		Result:       result,
	})
}

func Fail(c *gin.Context, code int, message string, result interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		ErrorCode:    code,
		ErrorMessage: message,
		Result:       result,
	})
}
