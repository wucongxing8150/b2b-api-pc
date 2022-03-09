// 邮寄地址
package Invoice

import (
	"fmt"

	Response "b2b-api-pc/App/Api/response"
	"b2b-api-pc/App/Model/InvoiceAddr"
	"github.com/gin-gonic/gin"
)

// List
// @Description: 邮寄地址列表
// @param c
func List(c *gin.Context) {
	userId, _ := c.Get("user_id")

	if userId == "" {
		Response.FailWithMessage("缺少参数", c)
		return
	}

	maps := make(map[string]interface{})
	maps["user_id"] = userId

	result := InvoiceAddr.Search(maps)
	fmt.Println(result)
	if len(result) <= 0 {
		Response.OkWithMessage("成功，数据为空", c)
		return
	}
	Response.OkWithData(result, c)
}
