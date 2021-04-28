package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Ctx *gin.Context
}

// 返回的结果
type ResultData struct {
	Code  int         `json:"code"`  //提示代码
	Msg   string      `json:"msg"`   //提示信息
	Total int         `json:"count"` //data total,count for layui
	Data  interface{} `json:"data"`  //数据
}

func New(ctx *gin.Context) *Result {
	return &Result{Ctx: ctx}
}

// 成功
func (r *Result) Success(data interface{}, total int) {
	if data == nil {
		data = gin.H{}
	}
	res := ResultData{}

	res.Code = 0
	res.Msg = "SUCCESS"
	res.Total = total
	res.Data = data
	r.Ctx.JSON(http.StatusOK, res)
}

// 失败
func (r *Result) Error(code int, msg string) {
	res := ResultData{}

	res.Code = code
	res.Msg = msg
	res.Total = 0
	res.Data = gin.H{}
	r.Ctx.JSON(http.StatusOK, res)
}
