package Invoice

import (
	Response "b2b-api-pc/App/Api/response"
	"b2b-api-pc/App/Cores/mysql"
	InvoiceEmailModel "b2b-api-pc/App/Logic/InvoiceEmail"
	"b2b-api-pc/App/Model"
	"b2b-api-pc/App/Validator"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"strconv"
	"time"
)

// ListEmail List
// @Description: 发票配置列表
// @param c
func ListEmail(c *gin.Context) {
	userId, _ := c.Get("user_id")

	if userId == "" {
		Response.FailWithMessage("缺少参数", c)
		return
	}

	maps := make(map[string]interface{})
	maps["user_id"] = userId

	result := InvoiceEmailModel.Search(maps)
	if len(result) <= 0 {
		Response.OkWithMessage("成功，数据为空", c)
		return
	}
	Response.OkWithData(result, c)
}

// DeleteEmail
// @Description: 发票邮箱删除
// @param c
func DeleteEmail(c *gin.Context) {
	userId, _ := c.Get("user_id")

	if userId == "" {
		Response.FailWithMessage("缺少参数", c)
		return
	}

	InvoiceEmailId, res := c.GetQuery("invoice_email_id")
	if !res || InvoiceEmailId == "" {
		Response.FailWithMessage("缺少参数", c)
		return
	}

	// 验证数据是否存在
	maps := make(map[string]interface{})
	maps["user_id"] = userId
	maps["invoice_email_id"] = InvoiceEmailId

	result := InvoiceEmailModel.Get(maps)
	if len(result) <= 0 {
		Response.FailWithMessage("非法数据", c)
		return
	}

	//
	Id, _ := strconv.ParseInt(InvoiceEmailId, 10, 64)

	_ = InvoiceEmailModel.DeleteId(Id)

	Response.Ok(c)
}

// AddEmail
// @Description:新增发票邮箱配置
// @param c
func AddEmail(c *gin.Context) {
	userId, _ := c.Get("user_id")

	if userId == "" {
		Response.FailWithMessage("缺少参数", c)
		return
	}

	// 查询是否超出最大限制
	maps := make(map[string]interface{})
	maps["user_id"] = userId

	result := InvoiceEmailModel.Search(maps)
	if len(result) >= 10 {
		Response.FailWithMessage("已超出最大邮寄地址限制", c)
		return
	}

	var InvoiceEmail Model.InvoiceEmail
	if err := c.ShouldBindJSON(&InvoiceEmail); err != nil {
		Response.FailWithMessage(fmt.Sprint(err), c)
		return
	}

	// 参数验证
	if err := Validator.Validate.Struct(InvoiceEmail); err != nil {
		Response.FailWithMessage(Validator.Translate(err), c)
		return
	}

	if InvoiceEmail.IsDefault == 1 {
		for _, r := range result {
			if r.IsDefault == 1 {
				Response.FailWithMessage("已存在默认邮箱", c)
				return
			}
		}
	}

	// 查询是否重复
	maps = make(map[string]interface{})
	maps["user_id"] = userId
	maps["email"] = InvoiceEmail.Email

	result = InvoiceEmailModel.Get(maps)
	if len(result) > 0 {
		Response.FailWithMessage("已存在相同数据，请勿重复提交", c)
		return
	}

	// 新增
	data := make(map[string]interface{})
	data["email"] = InvoiceEmail.Email
	data["tel"] = InvoiceEmail.Tel
	data["create_time"] = time.Now()
	data["update_time"] = time.Now()

	// 转换结构体
	err := mapstructure.Decode(data, &InvoiceEmail)
	if err != nil {
		fmt.Println(err.Error())
	}

	InvoiceEmail.UserId = userId.(string)

	_, res := InvoiceEmailModel.Add(InvoiceEmail)
	if res == false {
		Response.FailWithMessage("添加失败", c)
		return
	}

	Response.Ok(c)
}

// UpdateEmail
// @Description: 发票邮箱修改
// @param c
func UpdateEmail(c *gin.Context) {
	userId, _ := c.Get("user_id")

	if userId == "" {
		Response.FailWithMessage("缺少参数", c)
		return
	}

	var InvoiceEmail Model.InvoiceEmail
	if err := c.ShouldBindJSON(&InvoiceEmail); err != nil {
		Response.FailWithMessage(fmt.Sprint(err), c)
		return
	}

	// 参数验证
	if err := Validator.Validate.Struct(InvoiceEmail); err != nil {
		Response.FailWithMessage(Validator.Translate(err), c)
		return
	}

	// 查询是否存在
	maps := make(map[string]interface{})
	maps["invoice_email_id"] = InvoiceEmail.InvoiceEmailId
	maps["user_id"] = userId
	result := InvoiceEmailModel.Get(maps)
	if len(result) <= 0 {
		Response.FailWithMessage("非法请求", c)
		return
	}

	// 判断已是默认情况
	if InvoiceEmail.IsDefault == 1 && result[0].IsDefault == 1 {
		if InvoiceEmail.Email == result[0].Email && InvoiceEmail.Tel == result[0].Tel {
			Response.OkWithMessage("成功，已是默认", c)
			return
		}
	}

	// 开启事务
	tx := mysql.Db.Begin()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		Response.FailWithMessage(err.Error(), c)
		return
	}

	// 查询是否存在默认邮箱

	where := []interface{}{
		[]interface{}{"user_id", "=", userId},
		[]interface{}{"invoice_email_id", "!=", InvoiceEmail.InvoiceEmailId},
		[]interface{}{"is_default", "1"},
	}
	repeat := InvoiceEmailModel.GetWhere(where)
	if len(repeat) > 0 {
		// 存在默认邮箱，修改为非默认
		data := make(map[string]interface{})
		data["is_default"] = "0"
		data["invoice_email_id"] = repeat[0].InvoiceEmailId

		if err := tx.Model(&InvoiceEmailModel.TableStruct{}).Update(data).Error; err != nil {
			tx.Rollback()
			Response.FailWithMessage("修改失败", c)
			return
		}
	}

	// 查询是否重复
	where = []interface{}{
		[]interface{}{"user_id", "=", userId},
		[]interface{}{"invoice_email_id", "!=", InvoiceEmail.InvoiceEmailId},
		[]interface{}{"email", InvoiceEmail.Email},
	}

	repeat = InvoiceEmailModel.GetWhere(where)
	if len(repeat) > 0 {
		tx.Rollback()
		Response.FailWithMessage("已存在相同数据，请勿重复提交", c)
		return
	}

	InvoiceEmail.UserId = userId.(string)

	// 修改
	data := make(map[string]interface{})
	data["invoice_email_id"] = InvoiceEmail.InvoiceEmailId
	data["is_default"] = InvoiceEmail.IsDefault
	data["email"] = InvoiceEmail.Email
	data["tel"] = InvoiceEmail.Tel

	if err := tx.Model(&InvoiceEmailModel.TableStruct{}).Update(data).Error; err != nil {
		tx.Rollback()
		Response.FailWithMessage("修改失败", c)
		return
	}

	tx.Commit()

	Response.Ok(c)
}
