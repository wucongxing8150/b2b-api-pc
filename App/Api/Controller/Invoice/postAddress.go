// 邮寄地址
package Invoice

import (
	"fmt"

	Response "b2b-api-pc/App/Api/response"
	"b2b-api-pc/App/Model"
	"github.com/gin-gonic/gin"
)

// List
// @Description: 邮寄地址列表
// @param c
func List(c *gin.Context) {
	userId, _ := c.Get("user_id")

	if userId == "" {
		Response.OkWithMessage("缺少参数", c)
		return
	}

	maps := make(map[string]interface{})
	maps["ShopId"] = 1

	p2 := new(Model.InvoiceAddr)
	result := p2.Get(maps)
	fmt.Println(result)
}
