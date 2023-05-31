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
	Msg(code int, msg string, data interface{}) // 自定义返回
	Success(data interface{})                   // 返回 200 data信息
	Fail(msg string)                            // 返回400 无data信息
	Code(code int, msg string)                  // 返回自定义状态、消息 无data信息
}

type beJson struct {
	Ctx *gin.Context
}

type respJson struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewJson(c *gin.Context) BeJson {
	return &beJson{Ctx: c}
}

func (con *beJson) Code(code int, msg string) {
	con.Msg(code, msg, map[string]interface{}{})
}

func (con *beJson) Success(data interface{}) {
	con.Msg(http.StatusOK, "success", data)
}

func (con *beJson) Fail(msg string) {
	con.Msg(http.StatusBadRequest, msg, map[string]interface{}{})
}

func (con *beJson) Msg(code int, msg string, data interface{}) {
	con.Ctx.JSON(http.StatusOK, respJson{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
