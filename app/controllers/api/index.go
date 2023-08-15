package api

import (
	"framework/app/biz"
	"framework/ay"
	"framework/ay/lib"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
}

func (con IndexController) Index(c *gin.Context) {
	b := biz.NewBiz(c)
	res, err := ay.IgnoreNotFoundReturn(b.Demo.Get(1))
	if err != nil {
		lib.NewJson(c).Fail(err.Error())
		return
	}
	if res == nil {
		lib.NewJson(c).Fail("数据不存在")
		return
	}

	lib.NewJson(c).Msg(http.StatusOK, "success", res)
	//	c.JSON(http.StatusOK, gin.H{
	//	"code": 200,
	//	"msg":  "success",
	//	"data": gin.H{},
	//})
}
