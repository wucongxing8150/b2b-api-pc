package Model

import (
	"time"

	"b2b-api-pc/App/Cores/mysql"
)

var db = mysql.DB

type InvoiceAddr struct {
	InvoiceAddrId  int64     `gorm:"column:invoice_addr_id;primary_key;AUTO_INCREMENT" json:"invoice_addr_id"` // 主键id
	ShopId         int64     `gorm:"column:shop_id" json:"shop_id"`                                            // 店铺id
	UserId         string    `gorm:"column:user_id" json:"user_id"`                                            // 用户id
	IsDefault      int       `gorm:"column:is_default;default:0" json:"is_default"`                            // 是否默认使用 0.否 1.是
	ReceiverName   string    `gorm:"column:receiver_name" json:"receiver_name"`                                // 收件人姓名
	ReceiverMobile string    `gorm:"column:receiver_mobile" json:"receiver_mobile"`                            // 收件人电话
	ProvinceId     int64     `gorm:"column:province_id" json:"province_id"`                                    // 省ID
	Province       string    `gorm:"column:province" json:"province"`                                          // 省
	AreaId         int64     `gorm:"column:area_id" json:"area_id"`                                            // 区域ID
	Area           string    `gorm:"column:area" json:"area"`                                                  // 区
	CityId         int64     `gorm:"column:city_id" json:"city_id"`                                            // 城市ID
	City           string    `gorm:"column:city" json:"city"`                                                  // 城市
	Addr           string    `gorm:"column:addr" json:"addr"`                                                  // 地址
	PostCode       string    `gorm:"column:post_code" json:"post_code"`                                        // 邮编
	CreateTime     time.Time `gorm:"column:create_time" json:"create_time"`                                    // 创建时间
	UpdateTime     time.Time `gorm:"column:update_time" json:"update_time"`                                    // 修改时间
}

func (i InvoiceAddr) Get(maps interface{}) (invoiceAddr []InvoiceAddr) {

	db.First(&invoiceAddr)
	return
}
func Search(maps interface{}) (table []InvoiceAddr) {
	db.Model(&InvoiceAddr{}).Where(maps).Find(&table)
	return
}
func SearchPage(pageNum int, pageSize int, maps interface{}) (table []InvoiceAddr) {
	db.Model(&InvoiceAddr{}).Where(maps).Offset(pageNum).Limit(pageSize).Find(&table)
	return
}
func GetTotal(maps interface{}) (count int) {
	db.Model(&InvoiceAddr{}).Where(maps).Count(&count)
	return
}
func (i *InvoiceAddr) Add(Data map[string]interface{}) bool {
	db.Model(&InvoiceAddr{}).Create(&i)
	return !db.NewRecord(&i)
}
func EditId(id int, data interface{}) bool {
	db.Model(&InvoiceAddr{}).Where("id = ?", id).Updates(data)
	return true
}
func EditMap(maps interface{}, data interface{}) bool {
	db.Model(&InvoiceAddr{}).Where(maps).Updates(data)
	return true
}

// func (i *InvoiceAddr) BeforeCreate(scope *gorm.Scope) error {
// 	ID, err := Tool.NewWorker(1)
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = scope.SetColumn("InvoiceAddrId", ID.GetId())
// 	if err != nil {
// 		panic(err)
// 	}
// 	return nil
// }
