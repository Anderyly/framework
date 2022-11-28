/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package lib

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var _ BeJson = (*beJson)(nil)

type BeJson interface {
	Msg(code int, msg string, data interface{})
	Success(data interface{})
	Fail(msg string)
	Code(code int)
}

type beJson struct {
	Ctx *gin.Context
}

type restbeJson struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewJson(c *gin.Context) BeJson {
	return &beJson{Ctx: c}
}

func (con *beJson) Code(code int) {
	con.Msg(code, "success", map[string]interface{}{})
}

func (con *beJson) Success(data interface{}) {
	con.Msg(http.StatusOK, "success", data)
}

func (con *beJson) Fail(msg string) {
	con.Msg(http.StatusBadRequest, msg, map[string]interface{}{})
}

func (con *beJson) Msg(code int, msg string, data interface{}) {
	con.Ctx.JSON(http.StatusOK, restbeJson{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
