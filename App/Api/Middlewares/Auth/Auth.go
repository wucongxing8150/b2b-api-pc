package Auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// VerifyAuth
// @Description: 鉴权中间件
// @return gin.HandlerFunc
func VerifyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		RequestURI := c.Request.RequestURI

		fmt.Println(RequestURI)
		// authorization := c.Request.Header.Get("Authorization")

		c.Next()
	}
}
