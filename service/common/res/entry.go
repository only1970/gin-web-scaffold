package res

import (
	"gin-web-scaffold/utils/vaild"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

type ListResponse[T any] struct {
	List  []T `json:"list"`
	Count int `json:"count"`
}

const (
	Success = 200
	Failed  = 500
)

func OK(data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: Success,
		Data: data,
		Msg:  msg,
	})
}

func OkWithMsg(msg string, c *gin.Context) {
	OK(map[string]any{}, msg, c)
}

func OkWithData(data any, c *gin.Context) {
	OK(data, "query sucess", c)
}

func OkWithList[T any](list []T, count int, c *gin.Context) {
	if len(list) == 0 {
		list = []T{}
	}
	OK(ListResponse[T]{
		List:  list,
		Count: count,
	}, "List query sucess", c)
}

func Fail(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func FailWithMgs(msg string, c *gin.Context) {
	Fail(Failed, map[string]any{}, msg, c)
}

func FailWithData(data any, c *gin.Context) {
	Fail(Failed, data, "系统错误", c)
}

func FailErrMgs(err error, obj any, c *gin.Context) {
	FailWithMgs(vaild.GetVaildMsg(err, obj), c)
}

func FailWithDetailed(data interface{}, mgs string, c *gin.Context) {
	Fail(Failed, data, mgs, c)
}
