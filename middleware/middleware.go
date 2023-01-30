package middleware

import (
	"github.com/gin-gonic/gin"
)

func Instance(r *gin.Engine) *gin.Engine {
	r.Use(Cors())
	r.Use(Filter())
	//r.Use(Pretreatment())
	return r
}
