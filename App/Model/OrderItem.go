// @Description: 文件描述
// @Author: wucongxing
// @Date:2022/3/15 16:46

package Model

type OrderItem struct {
	OrderItemId               int64     `gorm:"column:order_item_id;primary_key;AUTO_INCREMENT" json:"order_item_id"`                          // 订单项ID
	ShopId                    int64     `gorm:"column:shop_id;NOT NULL" json:"shop_id"`                                                        // 店铺id
	OrderNumber               string    `gorm:"column:order_number;NOT NULL" json:"order_number"`                                              // 订单order_number
	ProdId                    int64     `gorm:"column:prod_id;NOT NULL" json:"prod_id"`                                                        // 产品ID
	SkuId                     int64     `gorm:"column:sku_id;NOT NULL" json:"sku_id"`                                                          // 产品SkuID
	CategoryId                int64     `gorm:"column:category_id" json:"category_id"`                                                         // 分类id
	ProdCount                 int       `gorm:"column:prod_count;default:0;NOT NULL" json:"prod_count"`                                        // 购物车产品个数
	ProdName                  string    `gorm:"column:prod_name;NOT NULL" json:"prod_name"`                                                    // 产品名称
	SkuName                   string    `gorm:"column:sku_name" json:"sku_name"`                                                               // sku名称
	Pic                       string    `gorm:"column:pic;NOT NULL" json:"pic"`                                                                // 产品主图片路径
	Price                     string    `gorm:"column:price;NOT NULL" json:"price"`                                                            // 产品价格
	UserId                    string    `gorm:"column:user_id;NOT NULL" json:"user_id"`                                                        // 用户Id
	ProductTotalAmount        string    `gorm:"column:product_total_amount;NOT NULL" json:"product_total_amount"`                              // 商品总金额
	RecTime                   LocalTime `gorm:"column:rec_time;NOT NULL" json:"rec_time"`                                                      // 购物时间
	CommSts                   int       `gorm:"column:comm_sts;default:0;NOT NULL" json:"comm_sts"`                                            // 评论状态： 0 未评价  1 已评价
	DistributionCardNo        string    `gorm:"column:distribution_card_no" json:"distribution_card_no"`                                       // 推广员使用的推销卡号
	BasketDate                LocalTime `gorm:"column:basket_date" json:"basket_date"`                                                         // 加入购物车时间
	ActualTotal               string    `gorm:"column:actual_total" json:"actual_total"`                                                       // 商品实际金额 = 商品总金额 - 分摊的优惠金额
	ShareReduce               string    `gorm:"column:share_reduce;NOT NULL" json:"share_reduce"`                                              // 分摊的优惠金额
	DistributionAmount        string    `gorm:"column:distribution_amount;default:0.00;NOT NULL" json:"distribution_amount"`                   // 推广员佣金
	DistributionParentAmount  string    `gorm:"column:distribution_parent_amount;default:0.00;NOT NULL" json:"distribution_parent_amount"`     // 上级推广员佣金
	PlatformShareReduce       string    `gorm:"column:platform_share_reduce;default:0.00;NOT NULL" json:"platform_share_reduce"`               // 平台优惠金额
	Rate                      string    `gorm:"column:rate;default:0.000000;NOT NULL" json:"rate"`                                             // 分账比例
	PlatformCommission        string    `gorm:"column:platform_commission;default:0.00;NOT NULL" json:"platform_commission"`                   // 平台佣金
	ScoreAmount               string    `gorm:"column:score_amount;default:0.00;NOT NULL" json:"score_amount"`                                 // 积分优惠金额
	MemberAmount              string    `gorm:"column:member_amount;default:0.00;NOT NULL" json:"member_amount"`                               // 等级优惠金额
	PlatformCouponAmount      string    `gorm:"column:platform_coupon_amount;default:0.00;NOT NULL" json:"platform_coupon_amount"`             // 平台优惠券优惠金额
	ShopCouponAmount          string    `gorm:"column:shop_coupon_amount;default:0.00;NOT NULL" json:"shop_coupon_amount"`                     // 店铺优惠券优惠金额
	DiscountAmount            string    `gorm:"column:discount_amount;default:0.00;NOT NULL" json:"discount_amount"`                           // 满减优惠金额
	PlatformFreeFreightAmount string    `gorm:"column:platform_free_freight_amount;default:0.00;NOT NULL" json:"platform_free_freight_amount"` // 平台运费减免金额
	FreeFreightAmount         string    `gorm:"column:free_freight_amount;default:0.00;NOT NULL" json:"free_freight_amount"`                   // 店铺运费减免金额
	ShopChangeFreeAmount      string    `gorm:"column:shop_change_free_amount;default:0.00;NOT NULL" json:"shop_change_free_amount"`           // 店铺改价优惠金额
	UseScore                  int       `gorm:"column:use_score;default:0" json:"use_score"`                                                   // 使用积分
	GainScore                 int       `gorm:"column:gain_score;default:0" json:"gain_score"`                                                 // 获得积分
	Status                    int       `gorm:"column:status;default:-1" json:"status"`                                                        // -1待发货 0全部发货 其他数量为剩余待发货数量
	DvyType                   int       `gorm:"column:dvy_type" json:"dvy_type"`                                                               // 单个orderItem的配送类型 1:快递 2:自提 3：无需快递 4:同城配送
	ChangeAmountVersion       int       `gorm:"column:change_amount_version" json:"change_amount_version"`                                     // 支付金额版本号
}

func (m *OrderItem) TableName() string {
	return "tz_order_item"
}
