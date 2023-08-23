package routers

import (
	"framework/app/controllers/api"
	"github.com/gin-gonic/gin"
)

func ApiRouter(r *gin.Engine) *gin.Engine {

	router := r.Group("/api")
	router.Any("", api.IndexController{}.Index)

	return r
}
