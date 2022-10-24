package routers

import (
	"github.com/gin-gonic/gin"
)

func Instance(r *gin.Engine) *gin.Engine {

	r = ApiRouter(r)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  "访问错误",
		})
	})

	return r
}
