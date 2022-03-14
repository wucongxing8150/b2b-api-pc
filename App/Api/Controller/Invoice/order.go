// @Description: 文件描述
// @Author: wucongxing
// @Date:2022/3/14 18:38

package Invoice

import (
	"fmt"

	Response "b2b-api-pc/App/Api/response"
	"b2b-api-pc/App/Validator"
	"github.com/gin-gonic/gin"
)

type Query struct {
	OrderNumber  string `json:"order_number"`
	Status       string `json:"status"`
	InvoiceState string `json:"invoice_state"`
	PayTimeStart string `json:"pay_time_start"`
	PayTimeEnd   string `json:"pay_time_end"`
}

// ListOrder
// @Description:可开发票列表
// @param c
func ListOrder(c *gin.Context) {
	userId, _ := c.Get("user_id")

	if userId == "" {
		Response.FailWithMessage("缺少参数", c)
		return
	}

	// 搜索条件
	// 订单号、订单状态、开票状态、下单时间
	var QueryStruct Query
	if err := c.ShouldBindJSON(&QueryStruct); err != nil {
		Response.FailWithMessage(fmt.Sprint(err), c)
		return
	}

	// 参数验证
	if err := Validator.Validate.Struct(QueryStruct); err != nil {
		Response.FailWithMessage(Validator.Translate(err), c)
		return
	}

	maps := make(map[string]interface{})

	if QueryStruct.OrderNumber
}
