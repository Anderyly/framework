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

func Header() gin.HandlerFunc {
	return func(c *gin.Context) {
		for name, values := range c.Request.Header {
			// Loop over all values for the name.
			for _, value := range values {
				if name == "Authorization" {
					_ = value
				}
			}
		}
		c.Next()
	}
}
