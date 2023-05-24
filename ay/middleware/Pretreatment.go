/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package middleware

import (
	"github.com/gin-gonic/gin"
)

func Pretreatment() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
