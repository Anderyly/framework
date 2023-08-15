package biz

import "github.com/gin-gonic/gin"

type Biz struct {
	Demo *demo
}

type base struct {
	ctx *gin.Context
	biz *Biz
}

func NewBiz(ctx *gin.Context) (b *Biz) {
	b = &Biz{}
	base := &base{ctx: ctx, biz: b}
	b.Demo = &demo{base}

	return
}
