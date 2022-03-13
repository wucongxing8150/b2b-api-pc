package Model

type Invoice struct {
	InvoiceId             int64     `gorm:"column:invoice_id;primary_key;AUTO_INCREMENT" json:"invoice_id"` // 申请id
	ShopId                int64     `gorm:"column:shop_id" json:"shop_id"`                                  // 申请/开具店铺id
	UserId                string    `gorm:"column:user_id" json:"user_id"`                                  // 申请用户id
	ApplyType             int       `gorm:"column:apply_type;default:1" json:"apply_type"`                  // 申请类型 1.平台-商家 2.商家-个人/商家
	InvoiceType           int       `gorm:"column:invoice_type;default:1" json:"invoice_type"`              // 发票类型 1.电子普通发票 2.增值税专用发票
	HeaderType            int       `gorm:"column:header_type;default:1" json:"header_type"`                // 抬头类型 1.单位 2.个人（增值税只允许单位）
	HeaderName            string    `gorm:"column:header_name" json:"header_name"`                          // 抬头名称
	InvoiceCode           string    `gorm:"column:invoice_code" json:"invoice_code"`                        // 发票税号
	InvoiceBank           string    `gorm:"column:invoice_bank" json:"invoice_bank"`                        // 开户行
	InvoiceBranchBank     string    `gorm:"column:invoice_branch_bank" json:"invoice_branch_bank"`          // 开户分行
	InvoiceCompanyTel     string    `gorm:"column:invoice_company_tel" json:"invoice_company_tel"`          // 注册电话
	InvoiceCompanyAddress string    `gorm:"column:invoice_company_address" json:"invoice_company_address"`  // 注册地址
	InvoiceBankAccount    string    `gorm:"column:invoice_bank_account" json:"invoice_bank_account"`        // 银行账号
	InvoiceNumber         string    `gorm:"column:invoice_number" json:"invoice_number"`                    // 发票号
	InvoiceContent        string    `gorm:"column:invoice_content" json:"invoice_content"`                  // 发票内容
	InvoiceState          int       `gorm:"column:invoice_state;default:0" json:"invoice_state"`            // 发票状态 0.申请中 1.申请成功 2.申请失败
	IsVoid                int       `gorm:"column:is_void;default:0" json:"is_void"`                        // 作废状态 0.正常 1.作废 2.换开
	ApplyTime             LocalTime `gorm:"column:apply_time" json:"apply_time"`                            // 申请时间
	VerifyTime            LocalTime `gorm:"column:verify_time" json:"verify_time"`                          // 审核时间
	FileId                int64     `gorm:"column:file_id" json:"file_id"`                                  // 文件id
	Reason                string    `gorm:"column:reason" json:"reason"`                                    // 驳回原因
}

func (m *Invoice) TableName() string {
	return "tz_invoice"
}
