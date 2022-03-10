package Validator

import (
	"b2b-api-pc/App/Tool"
	cnzh "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

var Validate *validator.Validate

// Validate/v10 全局验证器
var trans ut.Translator

// 初始化Validate/v10国际化
func init() {
	zh := cnzh.New()
	uni := ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")

	Validate = validator.New()

	//通过label标签返回自定义错误内容
	Validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			return field.Name
		}
		return label
	})
	_ = zhTranslations.RegisterDefaultTranslations(Validate, trans)

	//注册自定义函数和标签
	//手机号验证
	_ = Validate.RegisterValidation("mobile", mobile) //注册自定义函数，前一个参数是struct里tag自定义，后一个参数是自定义的函数

	//自定义required错误内容
	_ = Validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0}为必填字段!", false) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	//自定义max错误内容
	_ = Validate.RegisterTranslation("max", trans, func(ut ut.Translator) error {
		return ut.Add("max", "{0}超出最大长度", false) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("max", fe.Field())
		return t
	})

	//自定义min错误内容
	_ = Validate.RegisterTranslation("min", trans, func(ut ut.Translator) error {
		return ut.Add("min", "{0}超出最小长度", false) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("min", fe.Field())
		return t
	})

	//自定义lt错误内容
	_ = Validate.RegisterTranslation("lt", trans, func(ut ut.Translator) error {
		return ut.Add("lt", "{0}超出最大值", false) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("lt", fe.Field())
		return t
	})

	//自定义min错误内容
	_ = Validate.RegisterTranslation("gt", trans, func(ut ut.Translator) error {
		return ut.Add("gt", "{0}不满足最小值", false) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("gt", fe.Field())
		return t
	})

	//自定义email错误内容
	_ = Validate.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0}邮件格式错误", false) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})

	//自定义mobile错误内容
	_ = Validate.RegisterTranslation("mobile", trans, func(ut ut.Translator) error {
		return ut.Add("mobile", "手机号格式错误", false) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("mobile", fe.Field())
		return t
	})

}

// Translate 检验并返回检验错误信息
func Translate(err error) (errMsg string) {
	errs := err.(validator.ValidationErrors)
	for _, err := range errs {
		errMsg = err.Translate(trans)
	}
	return
}

//自定义手机号验证
func mobile(fl validator.FieldLevel) bool {
	return Tool.RegexpMobile(fl.Field().String())
}
