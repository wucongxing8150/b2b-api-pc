// @Description: 文件描述
// @Author: wucongxing
// @Date:2022/3/14 18:38

package Invoice

import (
	Response "b2b-api-pc/App/Api/response"
	"b2b-api-pc/App/Cores/mysql"
	"b2b-api-pc/App/Model"
	"b2b-api-pc/App/Validator"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
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
	var queryStruct Query
	var wheres []interface{}
	var orderInvoice Model.OrderInvoice

	// 订单号
	orderNumber, res := c.GetQuery("order_number")
	if res == false || orderNumber != "" {
		where := []interface{}{"order_number", "=", orderNumber}
		wheres = append(wheres, where)
	}

	// 订单状态
	status, res := c.GetQuery("status")
	if res == false || status != "" {
		where := []interface{}{"status", "=", status}
		wheres = append(wheres, where)
	}

	// 开票状态
	invoiceState, res := c.GetQuery("invoice_state")
	if res == false || invoiceState != "" {
		invoiceState, _ := strconv.ParseInt(invoiceState, 10, 0)
		orderInvoice.InvoiceState = int(invoiceState)
	}

	// 下单时间开始
	payTimeStart, res := c.GetQuery("pay_time_start")
	if res == false || payTimeStart != "" {
		where := []interface{}{"pay_time_start", ">=", payTimeStart}
		wheres = append(wheres, where)
	}

	// 下单时间结束
	payTimeEnd, res := c.GetQuery("pay_time_end")
	if res == false || payTimeEnd != "" {
		where := []interface{}{"pay_time_end", "<=", payTimeEnd}
		wheres = append(wheres, where)
	}

	// 参数验证
	if err := Validator.Validate.Struct(queryStruct); err != nil {
		Response.FailWithMessage(Validator.Translate(err), c)
		return
	}

	fmt.Println(wheres)
	db, err := Model.BuildWhere(mysql.Init(), wheres)
	if err != nil {
		fmt.Println("查询失败", err)
		Response.FailWithMessage("查询失败", c)
		return
	}

	result := db.Model(&Model.Order{}).Order("pay_time desc").Association("OrderInvoice").Find(&orderInvoice)

	fmt.Println(result)
	// Response.OkWithData(result,c)

}
