package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanlinhua/go-web-admin/app/lang"
)

type Result struct {
	Ctx *gin.Context
}

// 返回的结果
type ResultData struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Total int64       `json:"count"`
	Data  interface{} `json:"data"`
}

func New(ctx *gin.Context) *Result {
	return &Result{Ctx: ctx}
}

// 成功
func (r *Result) Success(data interface{}, total int64) {
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
	res.Msg = lang.Get(msg)
	res.Total = 0
	res.Data = gin.H{}
	r.Ctx.JSON(http.StatusOK, res)
}
