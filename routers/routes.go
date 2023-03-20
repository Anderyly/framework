package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Instance(r *gin.Engine) {

	ApiRouter(r)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusNotFound,
			"msg":  "访问错误",
		})
	})

}
