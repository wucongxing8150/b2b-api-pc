package Router

import (
	"fmt"
	"net/http"

	"b2b-api-pc/App/Api/Controller/Invoice"
	"b2b-api-pc/App/Api/Middlewares/Auth"
	"b2b-api-pc/App/Api/Middlewares/Corss"
	"b2b-api-pc/App/Api/Middlewares/Recover"
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

	// 处理异常
	r.Use(Recover.Recover())

	// 跨域中间件
	r.Use(Corss.Cors())

	// 鉴权中间件
	r.Use(Auth.VerifyAuth())

	// 路由群组
	invoice := r.Group("/invoice")

	postGroup := invoice.Group("/post")
	{
		// 邮寄地址列表
		postGroup.GET("address", Invoice.ListAddr)

		// 修改邮寄地址
		postGroup.PUT("address", Invoice.UpdateAddr)
	}

	return r
}
