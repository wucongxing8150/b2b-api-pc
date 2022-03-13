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

	// 发票
	invoice := r.Group("/invoice")
	{
		// 邮寄地址
		postGroup := invoice.Group("/post")
		{
			// 邮寄地址列表
			postGroup.GET("address", Invoice.ListAddr)

			// 修改邮寄地址
			postGroup.PUT("address", Invoice.UpdateAddr)

			// 新增邮寄地址
			postGroup.POST("address", Invoice.AddAddr)

			// 删除邮寄地址
			postGroup.DELETE("address", Invoice.DeleteAddr)

			// 邮寄地址设为默认
			postGroup.PUT("defaultAddress", Invoice.DefaultAddr)
		}

		// 发票配置
		sysGroup := invoice.Group("/sys")
		{
			// 发票配置列表
			sysGroup.GET("list", Invoice.ListSys)

			// 发票配置详情
			sysGroup.GET("detail", Invoice.DetailSys)

			// 发票配置删除
			sysGroup.DELETE("/", Invoice.DeleteSys)

			// 发票配置新增
			sysGroup.POST("/", Invoice.AddSys)

			// 发票配置修改
			sysGroup.PUT("/", Invoice.UpdateSys)
		}

		// 发票邮箱
		emailGroup := invoice.Group("/email")
		{
			// 发票邮箱列表
			emailGroup.GET("list", Invoice.ListEmail)

			// 发票邮箱删除
			emailGroup.DELETE("/", Invoice.DeleteEmail)

			// 发票邮箱新增
			emailGroup.POST("/", Invoice.AddEmail)

			// 发票邮箱修改
			emailGroup.PUT("/", Invoice.UpdateEmail)
		}

		// 申请开具发票
		applyGroup := invoice.Group("/invoice")
		{
			// 申请开具发票
			applyGroup.POST("apply", Invoice.ApplyInvoice)
		}

	}

	return r
}
