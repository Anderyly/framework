package api

import (
	"framework/ay/lib"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (con IndexController) Index(c *gin.Context) {
	lib.NewJson(c).Code(200, "")
	//	c.JSON(http.StatusOK, gin.H{
	//	"code": 200,
	//	"msg":  "success",
	//	"data": gin.H{},
	//})
}
