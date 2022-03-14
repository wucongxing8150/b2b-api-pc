package Model

type InvoiceItem struct {
	InvoiceItemId int64  `gorm:"column:invoice_item_id;primary_key;AUTO_INCREMENT" json:"invoice_item_id"` // 主键id
	InvoiceId     int64  `gorm:"column:invoice_id" json:"invoice_id"`                                      // 发票申请表id
	OrderNumber   string `gorm:"column:order_number" json:"order_number"`                                  // 订单编号
}

func (m *InvoiceItem) TableName() string {
	return "tz_invoice_item"
}
