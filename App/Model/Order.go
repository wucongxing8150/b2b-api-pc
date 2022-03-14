package Model

type Order struct {
	OrderId                   int64     `gorm:"column:order_id;primary_key;AUTO_INCREMENT" json:"order_id"`                                    // 订单ID
	ShopId                    int64     `gorm:"column:shop_id" json:"shop_id"`                                                                 // 店铺id
	ProdName                  string    `gorm:"column:prod_name;NOT NULL" json:"prod_name"`                                                    // 产品名称,多个产品将会以逗号隔开
	UserId                    string    `gorm:"column:user_id;NOT NULL" json:"user_id"`                                                        // 订购用户ID
	AddrOrderId               int64     `gorm:"column:addr_order_id" json:"addr_order_id"`                                                     // 用户订单地址Id
	OrderNumber               string    `gorm:"column:order_number;NOT NULL" json:"order_number"`                                              // 订购流水号
	Total                     string    `gorm:"column:total;default:0.00;NOT NULL" json:"total"`                                               // 总值
	ActualTotal               string    `gorm:"column:actual_total" json:"actual_total"`                                                       // 实际总值
	PayType                   int       `gorm:"column:pay_type" json:"pay_type"`                                                               // 支付方式 请参考枚举PayType 1:微信小程序支付 2:支付宝 3:微信扫码支付 4:微信h5支付 5:微信公众号支付 6:支付宝H5支付 7:支付宝APP支付 8:微信APP支付 9:余额支付 10:全球PayPal支付 11:线下支付
	PayMethod                 int       `gorm:"column:pay_method;default:1" json:"pay_method"`                                                 // 支付类型 1:线上支付 2:线下支付
	Remarks                   string    `gorm:"column:remarks" json:"remarks"`                                                                 // 买家备注
	ShopRemarks               string    `gorm:"column:shop_remarks" json:"shop_remarks"`                                                       // 卖家备注
	Status                    int       `gorm:"column:status;default:0;NOT NULL" json:"status"`                                                // 订单状态 1:待付款 2:待发货 3:待收货 4:待评价 5:成功 6:失败 7:待成团 8:待审核
	DvyType                   int       `gorm:"column:dvy_type" json:"dvy_type"`                                                               // 配送类型 1:快递 2:自提 3：无需快递 4同城配送
	DvyId                     int64     `gorm:"column:dvy_id" json:"dvy_id"`                                                                   // 配送方式ID
	DvyFlowId                 string    `gorm:"column:dvy_flow_id" json:"dvy_flow_id"`                                                         // 物流单号
	FreightAmount             string    `gorm:"column:freight_amount;default:0.00" json:"freight_amount"`                                      // 订单运费
	ProductNums               int       `gorm:"column:product_nums" json:"product_nums"`                                                       // 订单商品总数
	CreateTime                LocalTime `gorm:"column:create_time;NOT NULL" json:"create_time"`                                                // 订购时间
	UpdateTime                LocalTime `gorm:"column:update_time" json:"update_time"`                                                         // 订单更新时间
	PayTime                   LocalTime `gorm:"column:pay_time" json:"pay_time"`                                                               // 付款时间
	DvyTime                   LocalTime `gorm:"column:dvy_time" json:"dvy_time"`                                                               // 发货时间
	FinallyTime               LocalTime `gorm:"column:finally_time" json:"finally_time"`                                                       // 完成时间
	SettledTime               LocalTime `gorm:"column:settled_time" json:"settled_time"`                                                       // 结算时间
	CancelTime                LocalTime `gorm:"column:cancel_time" json:"cancel_time"`                                                         // 取消时间
	PreSaleTime               LocalTime `gorm:"column:pre_sale_time" json:"pre_sale_time"`                                                     // 预售发货时间
	IsPayed                   int       `gorm:"column:is_payed" json:"is_payed"`                                                               // 是否已经支付，1：已经支付过，0：，没有支付过
	DeleteStatus              int       `gorm:"column:delete_status;default:0" json:"delete_status"`                                           // 用户订单删除状态，0：没有删除， 1：回收站， 2：永久删除
	RefundStatus              int       `gorm:"column:refund_status" json:"refund_status"`                                                     // 订单退款状态（1:申请退款 2:退款成功 3:部分退款成功 4:退款失败）
	ReduceAmount              string    `gorm:"column:reduce_amount;default:0.00;NOT NULL" json:"reduce_amount"`                               // 优惠总额
	PlatformAmount            string    `gorm:"column:platform_amount;default:0.00;NOT NULL" json:"platform_amount"`                           // 平台优惠金额
	PlatformCommission        string    `gorm:"column:platform_commission;default:0.00;NOT NULL" json:"platform_commission"`                   // 平台佣金
	ScoreAmount               string    `gorm:"column:score_amount;default:0.00;NOT NULL" json:"score_amount"`                                 // 积分抵扣金额
	MemberAmount              string    `gorm:"column:member_amount;default:0.00;NOT NULL" json:"member_amount"`                               // 会员折扣金额
	DistributionAmount        string    `gorm:"column:distribution_amount;default:0.00;NOT NULL" json:"distribution_amount"`                   // 分销佣金金额
	ShopCouponAmount          string    `gorm:"column:shop_coupon_amount;default:0.00;NOT NULL" json:"shop_coupon_amount"`                     // 店铺优惠券优惠金额
	DiscountAmount            string    `gorm:"column:discount_amount;default:0.00;NOT NULL" json:"discount_amount"`                           // 满减优惠金额
	PlatformCouponAmount      string    `gorm:"column:platform_coupon_amount;default:0.00;NOT NULL" json:"platform_coupon_amount"`             // 平台优惠券优惠金额
	PlatformFreeFreightAmount string    `gorm:"column:platform_free_freight_amount;default:0.00;NOT NULL" json:"platform_free_freight_amount"` // 平台免运费金额
	FreeTransfee              string    `gorm:"column:free_transfee;default:0.00;NOT NULL" json:"free_transfee"`                               // 商家免运费金额
	ShopChangeFreeAmount      string    `gorm:"column:shop_change_free_amount;default:0.00;NOT NULL" json:"shop_change_free_amount"`           // 商家改价金额
	IsSettled                 int       `gorm:"column:is_settled" json:"is_settled"`                                                           // 是否已经进行结算 1.已结算，0未结算
	OrderType                 int       `gorm:"column:order_type" json:"order_type"`                                                           // 订单类型1团购订单 2秒杀订单 3积分订单
	CloseType                 int       `gorm:"column:close_type" json:"close_type"`                                                           // 订单关闭原因 1-超时未支付 2-退款关闭 4-买家取消 15-已通过货到付款交易
	ReceiverName              string    `gorm:"column:receiver_name" json:"receiver_name"`                                                     // 收件人姓名
	ReceiverMobile            string    `gorm:"column:receiver_mobile" json:"receiver_mobile"`                                                 // 收件人电话
	ChangeAmountVersion       int       `gorm:"column:change_amount_version" json:"change_amount_version"`                                     // 支付金额版本号
	OrderMold                 int       `gorm:"column:order_mold" json:"order_mold"`                                                           // 订单类别 1.实物商品订单 2. 虚拟商品订单
	VirtualRemark             string    `gorm:"column:virtual_remark" json:"virtual_remark"`                                                   // 虚拟商城的留言备注
	WriteOffStatus            int       `gorm:"column:write_off_status" json:"write_off_status"`                                               // 订单核销状态 0.待核销 1.核销完成
	WriteOffNum               int       `gorm:"column:write_off_num" json:"write_off_num"`                                                     // 核销次数 -1.多次核销 0.无需核销 1.单次核销
	WriteOffStart             LocalTime `gorm:"column:write_off_start" json:"write_off_start"`                                                 // 核销开始时间
	WriteOffEnd               LocalTime `gorm:"column:write_off_end" json:"write_off_end"`                                                     // 核销结束时间
	IsRefund                  int       `gorm:"column:is_refund" json:"is_refund"`                                                             // 是否可以退款 1.可以 0不可以
	UseLang                   int       `gorm:"column:use_lang;default:1" json:"use_lang"`                                                     // 使用语言 0中文 1中英文
}

func (m *Order) TableName() string {
	return "tz_order"
}
