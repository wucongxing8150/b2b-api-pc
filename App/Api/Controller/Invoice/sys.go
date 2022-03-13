// @Description: 文件描述
// @Author: wucongxing
// @Date:2022/3/11 13:09

package Invoice

import (
	"fmt"
	"strconv"
	"time"

	Response "b2b-api-pc/App/Api/response"
	InvoiceInfoModel "b2b-api-pc/App/Logic/InvoiceInfo"
	Model "b2b-api-pc/App/Model"
	"b2b-api-pc/App/Validator"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

// ListSys List
// @Description: 发票配置列表
// @param c
func ListSys(c *gin.Context) {
	userId, _ := c.Get("user_id")

	if userId == "" {
		Response.FailWithMessage("缺少参数", c)
		return
	}

	infoInvoiceType, res := c.GetQuery("info_invoice_type")
	if !res || infoInvoiceType == "" {
		Response.FailWithMessage("缺少参数", c)
		return
	}

	maps := make(map[string]interface{})
	maps["user_id"] = userId
	maps["info_type"] = infoInvoiceType

	result := InvoiceInfoModel.Search(maps)
	if len(result) <= 0 {
		Response.OkWithMessage("成功，数据为空", c)
		return
	}
	Response.OkWithData(result, c)
}

// DetailSys
// @Description:发票配置详情
// @param c
func DetailSys(c *gin.Context) {
	userId, _ := c.Get("user_id")

	if userId == "" {
		Response.FailWithMessage("缺少参数", c)
		return
	}

	invoiceInfoId, res := c.GetQuery("invoice_info_id")
	if !res || invoiceInfoId == "" {
		Response.FailWithMessage("缺少参数", c)
		return
	}

	maps := make(map[string]interface{})
	maps["user_id"] = userId
	maps["invoice_info_id"] = invoiceInfoId

	result := InvoiceInfoModel.Get(maps)
	if len(result) <= 0 {
		Response.FailWithMessage("非法请求", c)
		return
	}
	Response.OkWithData(result[0], c)
}

// DeleteSys
// @Description:发票配置删除
func DeleteSys(c *gin.Context) {
	userId, _ := c.Get("user_id")

	if userId == "" {
		Response.FailWithMessage("缺少参数", c)
		return
	}

	invoiceInfoId, res := c.GetQuery("invoice_info_id")

	Id, _ := strconv.ParseInt(invoiceInfoId, 10, 64)

	if !res || Id == 0 {
		Response.FailWithMessage("缺少参数", c)
		return
	}

	// 验证数据是否存在
	maps := make(map[string]interface{})
	maps["user_id"] = userId
	maps["invoice_info_id"] = Id

	result := InvoiceInfoModel.Get(maps)
	if len(result) <= 0 {
		Response.FailWithMessage("非法数据", c)
		return
	}

	_ = InvoiceInfoModel.DeleteId(Id)

	Response.Ok(c)
}

// AddSys
// @Description: 发票配置新增
// @param c
func AddSys(c *gin.Context) {
	userId, _ := c.Get("user_id")

	if userId == "" {
		Response.FailWithMessage("缺少参数", c)
		return
	}

	// 入参解析
	var InvoiceInfo Model.InvoiceInfo
	if err := c.ShouldBindJSON(&InvoiceInfo); err != nil {
		Response.FailWithMessage(fmt.Sprint(err), c)
		return
	}

	// 参数验证
	if err := Validator.Validate.Struct(InvoiceInfo); err != nil {
		Response.FailWithMessage(Validator.Translate(err), c)
		return
	}

	// 查询是否重复
	maps := make(map[string]interface{})
	maps["user_id"] = userId
	maps["info_type"] = 2
	maps["info_invoice_type"] = InvoiceInfo.InfoInvoiceType
	maps["header_name"] = InvoiceInfo.HeaderName
	maps["invoice_code"] = InvoiceInfo.InvoiceCode

	repeat := InvoiceInfoModel.Get(maps)
	if len(repeat) > 0 {
		Response.FailWithMessage("已存在相同数据，请勿重复提交", c)
		return
	}

	var result []InvoiceInfoModel.TableStruct

	// 判断增值税发票
	if InvoiceInfo.InfoInvoiceType == 2 {
		// 增值税类型发票，抬头必须是单位
		if InvoiceInfo.HeaderType != 1 {
			Response.FailWithMessage("增值税类型发票抬头类型只允许选择单位", c)
			return
		}

		maps = make(map[string]interface{})
		maps["user_id"] = userId
		maps["info_type"] = 2
		maps["info_invoice_type"] = 2

		result = InvoiceInfoModel.Get(maps)
		fmt.Println(result)

		if len(result) >= 1 {
			Response.FailWithMessage("已存在增值税发票配置", c)
			return
		}
	} else if InvoiceInfo.InfoInvoiceType == 1 {
		// 查询普通发票是否超出最大限制
		maps = make(map[string]interface{})
		maps["user_id"] = userId
		maps["info_type"] = 2
		maps["info_invoice_type"] = 2

		result = InvoiceInfoModel.Search(maps)

		if len(result) >= 10 {
			Response.FailWithMessage("只允许添加10个常用发票配置", c)
			return
		}
	} else {
		Response.FailWithMessage("非法请求", c)
		return
	}

	// 新增
	data := make(map[string]interface{})
	data["info_invoice_type"] = InvoiceInfo.InfoInvoiceType
	data["header_type"] = InvoiceInfo.HeaderType
	data["header_name"] = InvoiceInfo.HeaderName
	data["invoice_code"] = InvoiceInfo.InvoiceCode
	data["invoice_bank"] = InvoiceInfo.InvoiceBank
	data["invoice_branch_bank"] = InvoiceInfo.InvoiceBranchBank
	data["invoice_company_tel"] = InvoiceInfo.InvoiceCompanyTel
	data["invoice_company_address"] = InvoiceInfo.InvoiceCompanyAddress

	// 转换结构体
	err := mapstructure.Decode(data, &InvoiceInfo)
	if err != nil {
		fmt.Println(err.Error())
	}

	InvoiceInfo.UserId = userId.(string)
	InvoiceInfo.InfoType = 2
	InvoiceInfo.Status = 1
	InvoiceInfo.ApplyStatus = 1
	InvoiceInfo.ApplyTime = Model.LocalTime(time.Now())
	InvoiceInfo.VerifyTime = Model.LocalTime(time.Now())

	_, res := InvoiceInfoModel.Add(InvoiceInfo)
	if res == false {
		Response.FailWithMessage("添加失败", c)
		return
	}

	Response.Ok(c)
}

// UpdateSys
// @Description: 修改发票配置
// @param c
func UpdateSys(c *gin.Context) {
	userId, _ := c.Get("user_id")

	if userId == "" {
		Response.FailWithMessage("缺少参数", c)
		return
	}

	var InvoiceInfo Model.InvoiceInfo
	if err := c.ShouldBindJSON(&InvoiceInfo); err != nil {
		Response.FailWithMessage(fmt.Sprint(err), c)
		return
	}

	// 参数验证
	if err := Validator.Validate.Struct(InvoiceInfo); err != nil {
		Response.FailWithMessage(Validator.Translate(err), c)
		return
	}

	// 查询是否存在
	maps := make(map[string]interface{})
	maps["invoice_info_id"] = InvoiceInfo.InvoiceInfoId
	maps["user_id"] = userId
	result := InvoiceInfoModel.Get(maps)
	if len(result) <= 0 {
		Response.FailWithMessage("非法请求", c)
		return
	}

	// 不允许修改信息类型
	if result[0].InfoType != InvoiceInfo.InfoType {
		Response.FailWithMessage("非法修改", c)
		return
	}

	// 不允许修改发票类型
	if result[0].InfoInvoiceType != InvoiceInfo.InfoInvoiceType {
		Response.FailWithMessage("非法修改", c)
		return
	}

	// 不允许修改发票抬头类型
	if result[0].HeaderType != InvoiceInfo.HeaderType {
		Response.FailWithMessage("非法修改", c)
		return
	}

	// 查询是否重复
	maps = make(map[string]interface{})
	maps["user_id"] = userId
	maps["info_type"] = 2
	maps["info_invoice_type"] = InvoiceInfo.InfoInvoiceType
	maps["header_name"] = InvoiceInfo.HeaderName
	maps["invoice_code"] = InvoiceInfo.InvoiceCode

	repeat := InvoiceInfoModel.Get(maps)
	if len(repeat) > 0 {
		Response.FailWithMessage("已存在相同数据，请勿重复提交", c)
		return
	}

	InvoiceInfo.UserId = userId.(string)
	InvoiceInfo.InfoType = 2

	// 修改
	res := InvoiceInfoModel.Edit(InvoiceInfo)
	if res == false {
		Response.FailWithMessage("修改失败", c)
		return
	}
	Response.Ok(c)
}
