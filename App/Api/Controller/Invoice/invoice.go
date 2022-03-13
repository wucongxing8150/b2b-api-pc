package Invoice

import (
	Response "b2b-api-pc/App/Api/response"
	InvoiceInfoModel "b2b-api-pc/App/Logic/InvoiceInfo"
	"b2b-api-pc/App/Validator"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Apply struct {
	OrderNumber           string `json:"order_number" validate:"required" label:"订单号"`
	InvoiceType           string `json:"invoice_type" validate:"required" label:"invoice_type"`
	HeaderType            string `json:"header_type" validate:"required" label:"header_type"`
	HeaderName            string `json:"header_name" validate:"required" label:"header_name"`
	InvoiceCode           string `json:"invoice_code" validate:"required" label:"invoice_code"`
	InvoiceBank           string `json:"invoice_bank" validate:"required" label:"invoice_bank"`
	InvoiceBranchBank     string `json:"invoice_branch_bank" validate:"required" label:"invoice_branch_bank"`
	InvoiceCompanyTel     string `json:"invoice_company_tel" validate:"required" label:"invoice_company_tel"`
	InvoiceCompanyAddress string `json:"invoice_company_address" validate:"required" label:"invoice_company_address"`
	InvoiceBankAccount    string `json:"invoice_bank_account" validate:"required" label:"invoice_bank_account"`
	InvoiceContent        string `json:"invoice_content"`
}

// ApplyInvoice 申请开具发票
func ApplyInvoice(c *gin.Context) {
	userId, _ := c.Get("user_id")

	if userId == "" {
		Response.FailWithMessage("缺少参数", c)
		return
	}

	var ApplyStruct Apply
	if err := c.ShouldBindJSON(&ApplyStruct); err != nil {
		Response.FailWithMessage(fmt.Sprint(err), c)
		return
	}

	// 参数验证
	if err := Validator.Validate.Struct(ApplyStruct); err != nil {
		Response.FailWithMessage(Validator.Translate(err), c)
		return
	}

	// 检测用户是否可申请此种类型发票
	maps := make(map[string]interface{})
	maps["user_id"] = userId

	result := InvoiceInfoModel.Get(maps)
	if len(result) <= 0 {
		Response.FailWithMessage("非法请求", c)
		return
	}

	// 检测订单号是否存在
	// 检测订单状态是否符合开具要求
	// 检测订单是否已申请开具过
	// 添加申请表
	// 添加申请明细表
	// 添加订单发票表（tz_order_invoice）

}
