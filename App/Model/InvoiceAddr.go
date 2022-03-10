package Model

type InvoiceAddr struct {
	// gorm.Model
	InvoiceAddrId  int64     `gorm:"column:invoice_addr_id;primary_key;AUTO_INCREMENT" json:"invoice_addr_id"`                // 主键id
	ShopId         int64     `gorm:"column:shop_id" json:"shop_id" label:"shop_id"`                                           // 店铺id
	UserId         string    `gorm:"column:user_id" json:"user_id"`                                                           // 用户id
	IsDefault      int       `gorm:"column:is_default;default:0" json:"is_default"`                                           // 是否默认使用 0.否 1.是
	ReceiverName   string    `gorm:"column:receiver_name" json:"receiver_name" label:"receiver_name"`                         // 收件人姓名
	ReceiverMobile string    `gorm:"column:receiver_mobile" json:"receiver_mobile" validate:"Mobile" label:"receiver_mobile"` // 收件人电话
	ProvinceId     int64     `gorm:"column:province_id" json:"province_id"`                                                   // 省ID
	Province       string    `gorm:"column:province" json:"province"`                                                         // 省
	AreaId         int64     `gorm:"column:area_id" json:"area_id"`                                                           // 区域ID
	Area           string    `gorm:"column:area" json:"area"`                                                                 // 区
	CityId         int64     `gorm:"column:city_id" json:"city_id"`                                                           // 城市ID
	City           string    `gorm:"column:city" json:"city"`                                                                 // 城市
	Addr           string    `gorm:"column:addr" json:"addr"`                                                                 // 地址
	PostCode       string    `gorm:"column:post_code" json:"post_code"`                                                       // 邮编
	CreateTime     LocalTime `gorm:"column:create_time" json:"create_time"`                                                   // 创建时间
	UpdateTime     LocalTime `gorm:"column:update_time" json:"update_time"`                                                   // 修改时间
}

func (i *InvoiceAddr) TableName() string {
	return "tz_invoice_addr"
}
