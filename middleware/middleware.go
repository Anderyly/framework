package middleware

import (
	"framework/ay"
	"github.com/gin-gonic/gin"
)

func Instance(r *gin.Engine) {
	r.Use(Logger(ay.Logger))
	r.Use(Cors())
	//r.Use(Header())
	//r.Use(Pretreatment())
}
