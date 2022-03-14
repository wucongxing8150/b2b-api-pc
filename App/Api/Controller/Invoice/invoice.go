package Invoice

import (
	"fmt"
	"strings"
	"time"

	OrderModel "b2b-api-pc/App/Logic/Order"

	Response "b2b-api-pc/App/Api/response"
	"b2b-api-pc/App/Cores/mysql"
	InvoiceInfoModel "b2b-api-pc/App/Logic/InvoiceInfo"
	InvoiceItemModel "b2b-api-pc/App/Logic/InvoiceItem"
	"b2b-api-pc/App/Model"
	"b2b-api-pc/App/Validator"
	"github.com/gin-gonic/gin"
)

type Apply struct {
	OrderNumber           string `json:"order_number" validate:"required" label:"订单号"`
	InvoiceType           int    `json:"invoice_type" validate:"required" label:"invoice_type"`
	HeaderType            int    `json:"header_type" validate:"required" label:"header_type"`
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
	if ApplyStruct.InvoiceType == 2 {
		if ApplyStruct.HeaderType == 2 {
			Response.FailWithMessage("错误请求，增值税发票只允许商家抬头", c)
			return
		}

		maps := make(map[string]interface{})
		maps["user_id"] = userId
		maps["info_type"] = 2
		maps["info_invoice_type"] = 2
		maps["header_type"] = 1

		result := InvoiceInfoModel.Get(maps)
		if len(result) <= 0 {
			Response.FailWithMessage("请提前申请开具增值税发票资格", c)
			return
		}
	}

	// 分割订单号
	orderNumber := strings.Split(ApplyStruct.OrderNumber, ",")
	if len(orderNumber) <= 0 {
		Response.FailWithMessage("订单号错误", c)
		return
	}

	// 检测订单号是否存在
	for _, v := range orderNumber {
		// 检测订单是否已申请开具过
		maps := make(map[string]interface{})
		maps["order_number"] = v
		invoiceItem := InvoiceItemModel.Get(maps)
		if len(invoiceItem) > 0 {
			Response.FailWithMessage("存在已开具发票订单，请勿重复申请", c)
			return
		}
	}

	// 检测是否为同一店铺id
	where := []interface{}{
		[]interface{}{"user_id", "=", userId},
		[]interface{}{"order_number", "in", orderNumber},
	}
	orders := OrderModel.GetWhere(where)
	if len(orders) <= 0 {
		Response.FailWithMessage("订单数据错误", c)
		return
	}

	for k, v := range orders {
		if k == 0 {
			continue
		}

		if v.ShopId != orders[k-1].ShopId {
			Response.FailWithMessage("请选择同一店铺订单", c)
			return
		}
	}

	// 检测订单状态是否符合开具要求

	// 开启事务
	tx := mysql.Db.Begin()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		tx.Rollback()
		Response.FailWithMessage(err.Error(), c)
		return
	}

	// 添加申请表
	invoice := &Model.Invoice{
		UserId:                userId.(string),
		ApplyType:             2,
		InvoiceType:           ApplyStruct.InvoiceType,
		HeaderType:            ApplyStruct.HeaderType,
		HeaderName:            ApplyStruct.HeaderName,
		InvoiceCode:           ApplyStruct.InvoiceCode,
		InvoiceBank:           ApplyStruct.InvoiceBank,
		InvoiceBranchBank:     ApplyStruct.InvoiceBranchBank,
		InvoiceCompanyTel:     ApplyStruct.InvoiceCompanyTel,
		InvoiceCompanyAddress: ApplyStruct.InvoiceCompanyAddress,
		InvoiceBankAccount:    ApplyStruct.InvoiceBankAccount,
		InvoiceContent:        ApplyStruct.InvoiceContent,
		ApplyTime:             Model.LocalTime(time.Now()),
	}

	if err := tx.Create(&invoice).Error; err != nil {
		tx.Rollback()
		Response.FailWithMessage(err.Error(), c)
		return
	}

	// 添加申请明细表
	var InvoiceItems []Model.InvoiceItem

	for _, v := range orderNumber {
		var InvoiceItem = Model.InvoiceItem{
			InvoiceId:   invoice.InvoiceId,
			OrderNumber: v,
		}
		InvoiceItems = append(InvoiceItems, InvoiceItem)
	}

	if err := tx.Create(&InvoiceItems).Error; err != nil {
		tx.Rollback()
		Response.FailWithMessage(err.Error(), c)
		return
	}

	// 添加订单发票表（tz_order_invoice）
	for _, v := range orderNumber {
		maps := make(map[string]interface{})
		maps["user_id"] = userId
		maps["order_number"] = v
		order := OrderModel.Get(maps)
		if len(order) <= 0 {
			Response.FailWithMessage("存在非法订单", c)
			return
		}

		orderInvoice := &Model.OrderInvoice{
			OrderNumber:      v,
			ShopId:           order[0].OrderId,
			InvoiceType:      ApplyStruct.InvoiceType,
			HeaderType:       ApplyStruct.HeaderType,
			HeaderName:       ApplyStruct.HeaderName,
			InvoiceTaxNumber: ApplyStruct.InvoiceCode,
			InvoiceContext:   1,
			InvoiceState:     1,
			ApplicationTime:  Model.LocalTime(time.Now()),
			UploadTime:       Model.LocalTime(time.Now()),
		}

		if err := tx.Create(&orderInvoice).Error; err != nil {
			tx.Rollback()
			Response.FailWithMessage(err.Error(), c)
			return
		}
	}
	tx.Commit()
	Response.Ok(c)
}
