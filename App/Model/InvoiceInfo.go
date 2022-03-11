// @Description: 发票配置信息表
// @Author: wucongxing
// @Date:2022/3/11 13:12

package Model

type InvoiceInfo struct {
	InvoiceInfoId         int64     `gorm:"column:invoice_info_id;primary_key;AUTO_INCREMENT" json:"invoice_info_id"`                                          // 主键id
	ShopId                int64     `gorm:"column:shop_id" json:"shop_id"`                                                                                     // 店铺id
	UserId                string    `gorm:"column:user_id" json:"user_id"`                                                                                     // 用户id
	InfoType              int       `gorm:"column:info_type;default:1" json:"info_type"`                                                                       // 信息类型 1.商家 2.个人
	InfoInvoiceType       int       `gorm:"column:info_invoice_type;default:1" json:"info_invoice_type"`                                                       // 发票类型 1.电子普通发票 2.增值税专用发票
	HeaderType            int       `gorm:"column:header_type;default:1" json:"header_type"`                                                                   // 抬头类型 1.单位 2.个人（增值税只允许单位）
	HeaderName            string    `gorm:"column:header_name" json:"header_name" validate:"required" label:"header_name"`                                     // 抬头名称
	InvoiceCode           string    `gorm:"column:invoice_code" json:"invoice_code" validate:"required" label:"invoice_code"`                                  // 发票税号
	InvoiceBank           string    `gorm:"column:invoice_bank" json:"invoice_bank" validate:"required" label:"invoice_bank"`                                  // 开户行
	InvoiceBranchBank     string    `gorm:"column:invoice_branch_bank" json:"invoice_branch_bank" validate:"required" label:"invoice_branch_bank"`             // 开户分行
	InvoiceCompanyTel     string    `gorm:"column:invoice_company_tel" json:"invoice_company_tel" validate:"required" label:"invoice_company_tel"`             // 注册电话
	InvoiceCompanyAddress string    `gorm:"column:invoice_company_address" json:"invoice_company_address" validate:"required" label:"invoice_company_address"` // 注册地址
	InvoiceBankAccount    string    `gorm:"column:invoice_bank_account" json:"invoice_bank_account" validate:"required" label:"invoice_bank_account"`          // 银行账号
	ApplyStatus           int       `gorm:"column:apply_status;default:0" json:"apply_status"`                                                                 // 申请状态 0.申请中 1.申请成功 2.申请失败
	Status                int       `gorm:"column:status;default:1" json:"status"`                                                                             // 状态 1.正常 2.禁用
	ApplyTime             LocalTime `gorm:"column:apply_time" json:"apply_time"`                                                                               // 申请时间
	VerifyTime            LocalTime `gorm:"column:verify_time" json:"verify_time"`                                                                             // 审核时间
	Reason                string    `gorm:"column:reason" json:"reason"`                                                                                       // 驳回原因
}

func (m *InvoiceInfo) TableName() string {
	return "tz_invoice_info"
}
