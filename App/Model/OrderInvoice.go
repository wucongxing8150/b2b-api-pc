// @Description: 文件描述
// @Author: wucongxing
// @Date:2022/3/14 16:50

package Model

type OrderInvoice struct {
	OrderInvoiceId   int64     `gorm:"column:order_invoice_id;primary_key;AUTO_INCREMENT" json:"order_invoice_id"`         // 订单发票ID
	OrderNumber      string    `gorm:"column:order_number;NOT NULL" json:"order_number"`                                   // 订单编号
	ShopId           int64     `gorm:"column:shop_id;NOT NULL" json:"shop_id"`                                             // 店铺id
	InvoiceType      int       `gorm:"column:invoice_type;default:1;NOT NULL" json:"invoice_type"`                         // 发票类型 1.电子普通发票
	HeaderType       int       `gorm:"column:header_type;NOT NULL" json:"header_type"`                                     // 抬头类型 1.单位 2.个人
	HeaderName       string    `gorm:"column:header_name" json:"header_name"`                                              // 抬头名称
	InvoiceTaxNumber string    `gorm:"column:invoice_tax_number" json:"invoice_tax_number"`                                // 发票税号
	InvoiceContext   int       `gorm:"column:invoice_context;default:1;NOT NULL" json:"invoice_context"`                   // 发票内容 1.商品明细
	InvoiceState     int       `gorm:"column:invoice_state;NOT NULL" json:"invoice_state"`                                 // 发票状态 1.申请中 2.已开票
	FileId           int64     `gorm:"column:file_id" json:"file_id"`                                                      // 文件id
	ApplicationTime  LocalTime `gorm:"column:application_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"application_time"` // 申请时间
	UploadTime       LocalTime `gorm:"column:upload_time" json:"upload_time"`                                              // 上传时间
}

func (m *OrderInvoice) TableName() string {
	return "tz_order_invoice"
}
