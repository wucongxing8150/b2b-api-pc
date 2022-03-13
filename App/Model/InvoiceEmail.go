package Model

type InvoiceEmail struct {
	InvoiceEmailId int64     `gorm:"column:invoice_email_id;primary_key;AUTO_INCREMENT" json:"invoice_email_id"` // 主键id
	ShopId         int64     `gorm:"column:shop_id" json:"shop_id"`                                              // 店铺id
	UserId         string    `gorm:"column:user_id" json:"user_id"`                                              // 用户id
	IsDefault      int       `gorm:"column:is_default;default:0" json:"is_default"`                              // 是否默认使用 0.否 1.是
	Email          string    `gorm:"column:email" json:"email" validate:"required,email" label:"email"`          // 邮件地址
	Tel            string    `gorm:"column:tel" json:"tel" validate:"Mobile,required" label:"tel"`               // 电话
	CreateTime     LocalTime `gorm:"column:create_time" json:"create_time"`                                      // 创建时间
	UpdateTime     LocalTime `gorm:"column:update_time" json:"update_time"`                                      // 修改时间
}

func (m *InvoiceEmail) TableName() string {
	return "tz_invoice_email"
}
