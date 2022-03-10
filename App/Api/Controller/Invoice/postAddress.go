// 邮寄地址
package Invoice

import (
	Response "b2b-api-pc/App/Api/response"
	InvoiceAddrModel "b2b-api-pc/App/Logic/InvoiceAddr"
	"b2b-api-pc/App/Model"
	"b2b-api-pc/App/Validator"
	"fmt"
	"github.com/gin-gonic/gin"
)

// ListAddr List
// @Description: 邮寄地址列表
// @param c
func ListAddr(c *gin.Context) {
	userId, _ := c.Get("user_id")

	if userId == "" {
		Response.FailWithMessage("缺少参数", c)
		return
	}

	maps := make(map[string]interface{})
	maps["user_id"] = userId

	result := InvoiceAddrModel.Search(maps)
	fmt.Println(result)
	if len(result) <= 0 {
		Response.OkWithMessage("成功，数据为空", c)
		return
	}
	Response.OkWithData(result, c)
}

// UpdateAddr 邮寄地址修改
func UpdateAddr(c *gin.Context) {
	userId, _ := c.Get("user_id")

	if userId == "" {
		Response.FailWithMessage("缺少参数", c)
		return
	}

	var invoiceAddr Model.InvoiceAddr
	if err := c.ShouldBindJSON(&invoiceAddr); err != nil {
		Response.FailWithMessage(fmt.Sprint(err), c)
		return
	}

	if err := Validator.Validate.Struct(invoiceAddr); err != nil {
		Response.FailWithMessage(Validator.Translate(err), c)
		return
	}

	//获取邮寄地址是否存在
	maps := make(map[string]interface{})
	maps["invoice_addr_id"] = invoiceAddr.InvoiceAddrId
	maps["user_id"] = userId
	result := InvoiceAddrModel.Get(maps)
	fmt.Println(result)
	if len(result) <= 0 {
		Response.FailWithMessage("用户数据错误", c)
		return
	}

	//执行修改

	res := InvoiceAddrModel.Edit(invoiceAddr)
	if res == false {
		Response.FailWithMessage("修改失败", c)
		return
	}
	Response.Ok(c)
}
