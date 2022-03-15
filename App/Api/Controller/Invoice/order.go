// @Description: 文件描述
// @Author: wucongxing
// @Date:2022/3/14 18:38

package Invoice

import (
	"fmt"
	"strconv"

	Response "b2b-api-pc/App/Api/response"
	"b2b-api-pc/App/Cores/mysql"
	"b2b-api-pc/App/Model"
	"b2b-api-pc/App/Tool"
	"b2b-api-pc/App/Validator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Query struct {
	OrderNumber  string `json:"order_number"`
	Status       string `json:"status"`
	InvoiceState string `json:"invoice_state"`
	PayTime      string `json:"pay_time"`
	Page         string `json:"page"`
	PageSize     string `json:"pageSize"`
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
	// 订单号、订单状态、开票状态、下单时间、页码、每页个数
	var queryStruct Query
	var wheres []interface{}

	// 订单号
	orderNumber, res := c.GetQuery("order_number")
	if res == false || orderNumber != "" {
		where := []interface{}{"tz_order.order_number", "=", orderNumber}
		wheres = append(wheres, where)
	}

	// 订单状态
	status, res := c.GetQuery("status")
	if res == false || status != "" {
		where := []interface{}{"tz_order.status", "=", status}
		wheres = append(wheres, where)
	}

	// 开票状态
	// invoiceState, res := c.GetQuery("invoice_state")
	// if res == false || invoiceState != "" {
	// 	where := []interface{}{"OrderInvoice.invoice_state", "=", invoiceState}
	// 	wheres = append(wheres, where)
	// }

	// 下单时间开始
	PayTime, res := c.GetQuery("pay_time")
	if res == false || PayTime != "" {
		// 全部
		// 最近7天
		if PayTime == "2" {
			where := []interface{}{"tz_order.pay_time", ">=", Tool.BeforeDate(-1000)}
			wheres = append(wheres, where)

			where = []interface{}{"tz_order.pay_time", "<=", Tool.BeforeDate(0)}
			wheres = append(wheres, where)
		}
		// 最近三个月
		if PayTime == "3" {
			where := []interface{}{"tz_order.pay_time", ">=", Tool.BeforeDate(90)}
			wheres = append(wheres, where)

			where = []interface{}{"tz_order.pay_time", "<=", Tool.BeforeDate(0)}
			wheres = append(wheres, where)
		}
		// 三个月以前
		if PayTime == "4" {
			where := []interface{}{"tz_order.pay_time", ">=", Tool.BeforeDate(90)}
			wheres = append(wheres, where)
		}
	}

	// 页码
	page, res := c.GetQuery("page")
	if res == false || page != "" {
		page = "1"
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return
	}

	// 每页个数
	pageSize, res := c.GetQuery("pageSize")
	if res == false || pageSize != "" {
		pageSize = "20"
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		return
	}

	// 参数验证
	if err := Validator.Validate.Struct(queryStruct); err != nil {
		Response.FailWithMessage(Validator.Translate(err), c)
		return
	}

	var orderModel []*Model.Order
	var order Model.Order

	db, err := Model.BuildQueryList(mysql.Db, wheres, []string{"order_id,shop_id,order_number"}, "pay_time desc", pageInt, pageSizeInt)
	if err != nil {
		fmt.Println("查询失败", err)
		Response.FailWithMessage("查询失败", c)
		return
	}

	if err := db.Preload("OrderInvoice", func(db *gorm.DB) *gorm.DB {
		return db.Select("order_invoice_id,order_number,invoice_state")
	}).Preload("ShopDetail", func(db *gorm.DB) *gorm.DB {
		return db.Select("shop_id,shop_name")
	}).Preload("OrderItem").Find(&orderModel).Error; err != nil {
		Response.FailWithMessage(err.Error(), c)
		return
	}

	var dataArray map[string]interface{}

	// 处理数据
	if len(orderModel) >= 0 {
		for k, v := range orderModel {
			dataArray["order_id"] = v.OrderId
			dataArray["shop_id"] = v.ShopId
			dataArray["shop_name"] = ""
			if v.ShopDetail.ShopName != "" {
				dataArray["shop_name"] = v.ShopDetail.ShopName
			}
			dataArray["shop_id"] = v.ShopId
		}
	}

	total := int64(0)
	db.Model(&order).Count(&total)

	data := map[string]interface{}{
		"dataArray": &orderModel,
		"page":      pageInt,
		"pageSize":  pageSizeInt,
		"total":     total,
	}
	Response.OkWithData(data, c)

}
