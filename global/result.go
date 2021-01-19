package global

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Ctx *gin.Context
}

type ResultCont struct {
	Code	int  `json:"code"` // 自增
	Msg string  `json:"msg"` //
	Data interface{} `json:"data"`
}

func NewResult(ctx *gin.Context) *Result {
	return &Result{Ctx: ctx}
}

func (r *Result) Success(data interface{}) {
//func Succ(data interface{}) {
	if (data == nil) {
		data = gin.H{}
	}
	res := ResultCont{}
	res.Code = 0
	res.Msg = ""
	res.Data = data
	r.Ctx.JSON(http.StatusOK,res)
}

func (r *Result)Error(code int,msg string) {
	res := ResultCont{}
	res.Code = code
	res.Msg = msg
	res.Data = gin.H{}
	r.Ctx.JSON(http.StatusOK,res)
	r.Ctx.Abort()
}
