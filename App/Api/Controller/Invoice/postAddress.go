// 邮寄地址
package Invoice

import (
	"fmt"
	"github.com/gin-gonic/gin"

	Response "b2b-api-pc/App/Api/response"
	InvoiceAddrModel "b2b-api-pc/App/Model/InvoiceAddr"
)

type InvoiceAddr struct {
	InvoiceAddrID  int    `json:"invoice_addr_id" form:"invoice_addr_id" binding:"required"`
	ReceiverName   string `json:"receiver_name"`
	ReceiverMobile int    `json:"receiver_mobile"`
	ProvinceID     int    `json:"province_id"`
	Province       string `json:"province"`
	AreaID         int    `json:"area_id"`
	Area           string `json:"area"`
	CityID         int    `json:"city_id"`
	City           string `json:"city"`
	Addr           string `json:"addr"`
	PostCode       string `json:"post_code"`
}

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

// UpdateAddr
func UpdateAddr(c *gin.Context) {
	userId, _ := c.Get("user_id")

	if userId == "" {
		Response.FailWithMessage("缺少参数", c)
		return
	}

	var invoiceAddr InvoiceAddr
	if err := c.ShouldBind(&invoiceAddr); err != nil {
		Response.FailWithMessage(fmt.Sprint(err), c)
		return
	}
	fmt.Println(invoiceAddr)
}
