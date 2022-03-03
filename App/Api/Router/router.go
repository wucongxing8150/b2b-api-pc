package Router

import (
	"fmt"
	"net/http"

	"b2b-api-pc/App/Api/Middlewares/Auth"
	"b2b-api-pc/App/Api/Middlewares/Corss"
	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var options = []Option{}

// 初始化
func Init() *gin.Engine {

	r := gin.Default()

	// 404处理
	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  fmt.Sprintf("%s %s not found", method, path),
			"code": 404,
			"data": "",
		})
	})

	// 跨域中间件
	r.Use(Corss.Cors())

	// 鉴权中间件
	r.Use(Auth.VerifyAuth())

	return r
}
