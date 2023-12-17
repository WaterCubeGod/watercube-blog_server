package res

import (
	"github.com/gin-gonic/gin"
	"gvb_server/utils"
	"net/http"
)

type Response struct {
	Code    CodeType `json:"code"`
	Data    any      `json:"data"`
	Message string   `json:"message"`
}

type ListResponse[T any] struct {
	Count int64 `json:"count"`
	List  T     `json:"list"`
}

func Result(code CodeType, data any, message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Data:    data,
		Message: message,
	})
}

func OK(data any, msg string, c *gin.Context) {
	Result(SUCCESS, data, msg, c)
}

func OKWith(c *gin.Context) {
	Result(SUCCESS, "", "成功", c)
}

func OKWithData(data any, c *gin.Context) {
	Result(SUCCESS, data, "成功", c)
}

func OKWithList(list any, count int64, c *gin.Context) {
	OKWithData(ListResponse[any]{
		List:  list,
		Count: count,
	}, c)
}

func OKWithMessage(msg string, c *gin.Context) {
	Result(SUCCESS, map[string]any{}, msg, c)
}

func Fail(data any, msg string, c *gin.Context) {
	Result(Error, data, msg, c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(Error, map[string]any{}, msg, c)
}

func FailWithError(err error, obj any, field string, c *gin.Context) {
	msg := utils.GetValidMsg(err, obj, field)
	FailWithMessage(msg, c)
}

func FailWithCode(code CodeType, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(code, map[string]any{}, msg, c)
		return
	}
	Result(Error, map[string]any{}, "未知错误", c)
}
