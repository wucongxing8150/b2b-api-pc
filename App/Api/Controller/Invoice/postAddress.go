// 邮寄地址
package Invoice

import (
	"fmt"
	"strconv"
	"time"

	Response "b2b-api-pc/App/Api/response"
	"b2b-api-pc/App/Cores/mysql"
	InvoiceAddrModel "b2b-api-pc/App/Logic/InvoiceAddr"
	"b2b-api-pc/App/Model"
	"b2b-api-pc/App/Validator"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
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

	// 参数验证
	if err := Validator.Validate.Struct(invoiceAddr); err != nil {
		Response.FailWithMessage(Validator.Translate(err), c)
		return
	}

	// 获取邮寄地址是否存在
	maps := make(map[string]interface{})
	maps["invoice_addr_id"] = invoiceAddr.InvoiceAddrId
	maps["user_id"] = userId
	result := InvoiceAddrModel.Get(maps)
	fmt.Println(result)
	if len(result) <= 0 {
		Response.FailWithMessage("用户数据错误", c)
		return
	}

	// 执行修改

	res := InvoiceAddrModel.Edit(invoiceAddr)
	if res == false {
		Response.FailWithMessage("修改失败", c)
		return
	}
	Response.Ok(c)
}

func AddAddr(c *gin.Context) {
	userId, _ := c.Get("user_id")

	if userId == "" {
		Response.FailWithMessage("缺少参数", c)
		return
	}

	// 查询是否超出最大限制
	maps := make(map[string]interface{})
	maps["user_id"] = userId

	result := InvoiceAddrModel.Search(maps)
	if len(result) >= 10 {
		Response.FailWithMessage("已超出最大邮寄地址限制", c)
		return
	}

	var invoiceAddr Model.InvoiceAddr
	if err := c.ShouldBindJSON(&invoiceAddr); err != nil {
		Response.FailWithMessage(fmt.Sprint(err), c)
		return
	}

	// 参数验证
	if err := Validator.Validate.Struct(invoiceAddr); err != nil {
		Response.FailWithMessage(Validator.Translate(err), c)
		return
	}

	if invoiceAddr.IsDefault == 1 {
		for _, r := range result {
			if r.IsDefault == 1 {
				Response.FailWithMessage("已存在默认邮寄地址", c)
				return
			}
		}
	}

	// 查询是否重复
	maps = make(map[string]interface{})
	maps["user_id"] = userId
	maps["receiver_name"] = invoiceAddr.ReceiverName
	maps["receiver_mobile"] = invoiceAddr.ReceiverMobile
	maps["province_id"] = invoiceAddr.ProvinceId
	maps["province"] = invoiceAddr.Province
	maps["area_id"] = invoiceAddr.AreaId
	maps["area"] = invoiceAddr.Area
	maps["city_id"] = invoiceAddr.CityId
	maps["city"] = invoiceAddr.City
	maps["addr"] = invoiceAddr.Addr
	maps["post_code"] = invoiceAddr.PostCode

	result = InvoiceAddrModel.Get(maps)
	if len(result) > 0 {
		Response.FailWithMessage("已存在相同数据，请勿重复提交", c)
		return
	}

	// 新增
	data := make(map[string]interface{})
	data["user_id"] = userId
	data["is_default"] = invoiceAddr.IsDefault
	data["receiver_name"] = invoiceAddr.ReceiverName
	data["receiver_mobile"] = invoiceAddr.ReceiverMobile
	data["province_id"] = invoiceAddr.ProvinceId
	data["province"] = invoiceAddr.Province
	data["area_id"] = invoiceAddr.AreaId
	data["area"] = invoiceAddr.Area
	data["city_id"] = invoiceAddr.CityId
	data["city"] = invoiceAddr.City
	data["addr"] = invoiceAddr.Addr
	data["post_code"] = invoiceAddr.PostCode
	data["create_time"] = time.Now()
	data["update_time"] = time.Now()

	// 转换结构体
	err := mapstructure.Decode(data, &invoiceAddr)
	if err != nil {
		fmt.Println(err.Error())
	}

	invoiceAddr.UserId = userId.(string)

	_, res := InvoiceAddrModel.Add(invoiceAddr)
	if res == false {
		Response.FailWithMessage("添加失败", c)
		return
	}

	Response.Ok(c)
}

// DeleteAddr
// @Description:删除邮寄地址
func DeleteAddr(c *gin.Context) {
	userId, _ := c.Get("user_id")

	if userId == "" {
		Response.FailWithMessage("缺少参数", c)
		return
	}

	InvoiceAddrId, res := c.GetQuery("invoice_addr_id")
	if !res || InvoiceAddrId == "" {
		Response.FailWithMessage("缺少参数", c)
		return
	}

	// 验证数据是否存在
	maps := make(map[string]interface{})
	maps["user_id"] = userId
	maps["invoice_addr_id"] = InvoiceAddrId

	result := InvoiceAddrModel.Get(maps)
	if len(result) <= 0 {
		Response.FailWithMessage("非法数据", c)
		return
	}

	//
	Id, _ := strconv.ParseInt(InvoiceAddrId, 10, 64)

	_ = InvoiceAddrModel.DeleteId(Id)

	Response.Ok(c)
}

// DefaultAddr
// @Description:邮寄地址设为默认
func DefaultAddr(c *gin.Context) {
	userId, _ := c.Get("user_id")

	if userId == "" {
		Response.FailWithMessage("缺少参数", c)
		return
	}

	InvoiceAddrId, res := c.GetQuery("invoice_addr_id")
	if !res || InvoiceAddrId == "" {
		Response.FailWithMessage("缺少参数", c)
		return
	}

	// 验证数据是否存在
	maps := make(map[string]interface{})
	maps["user_id"] = userId
	maps["invoice_addr_id"] = InvoiceAddrId

	result := InvoiceAddrModel.Get(maps)
	if len(result) <= 0 {
		Response.FailWithMessage("非法数据", c)
		return
	}

	// 开启事务
	tx := mysql.Db.Begin()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			tx.Rollback()
		}
	}()

	Id, _ := strconv.ParseInt(InvoiceAddrId, 10, 64)

	_ = InvoiceAddrModel.DeleteId(Id)

	Response.Ok(c)
}
