// @Description: 文件描述
// @Author: wucongxing
// @Date:2022/3/15 15:41

package Model

type ShopDetail struct {
	ShopId              int64     `gorm:"column:shop_id;primary_key;AUTO_INCREMENT" json:"shop_id"`  // 店铺id
	ShopName            string    `gorm:"column:shop_name" json:"shop_name"`                         // 店铺名称(数字、中文，英文(可混合，不可有特殊字符)，可修改)、不唯一
	UserId              string    `gorm:"column:user_id" json:"user_id"`                             // 店长用户id
	Intro               string    `gorm:"column:intro" json:"intro"`                                 // 店铺简介(可修改)
	ShopOwner           string    `gorm:"column:shop_owner" json:"shop_owner"`                       // 店长
	Mobile              string    `gorm:"column:mobile" json:"mobile"`                               // 店铺绑定的手机(登录账号：唯一)
	Password            string    `gorm:"column:password" json:"password"`                           // 登录密码
	ReceiveMobile       string    `gorm:"column:receive_mobile" json:"receive_mobile"`               // 接收短信号码
	Tel                 string    `gorm:"column:tel" json:"tel"`                                     // 店铺联系电话
	ShopLat             string    `gorm:"column:shop_lat" json:"shop_lat"`                           // 店铺所在纬度(可修改)
	ShopLng             string    `gorm:"column:shop_lng" json:"shop_lng"`                           // 店铺所在经度(可修改)
	ShopAddress         string    `gorm:"column:shop_address" json:"shop_address"`                   // 店铺详细地址
	Province            string    `gorm:"column:province" json:"province"`                           // 店铺所在省份（描述）
	ProvinceId          int64     `gorm:"column:province_id" json:"province_id"`                     // 店铺所在省份Id
	City                string    `gorm:"column:city" json:"city"`                                   // 店铺所在城市（描述）
	CityId              int64     `gorm:"column:city_id" json:"city_id"`                             // 店铺所在城市id
	Area                string    `gorm:"column:area" json:"area"`                                   // 店铺所在区域（描述）
	AreaId              int64     `gorm:"column:area_id" json:"area_id"`                             // 店铺所在区域Id
	ShopLogo            string    `gorm:"column:shop_logo" json:"shop_logo"`                         // 店铺logo(可修改)
	ShopStatus          int       `gorm:"column:shop_status" json:"shop_status"`                     // 店铺状态(-1:已删除 0: 停业中 1:营业中 2:平台下线 3:平台下线待审核 4:开店申请中 5:开店申请待审核)
	CreateTime          LocalTime `gorm:"column:create_time" json:"create_time"`                     // 创建时间
	UpdateTime          LocalTime `gorm:"column:update_time" json:"update_time"`                     // 更新时间
	IsDistribution      int       `gorm:"column:is_distribution" json:"is_distribution"`             // 分销开关(0:开启 1:关闭)
	BusinessLicense     string    `gorm:"column:business_license" json:"business_license"`           // 营业执照
	IdentityCardFront   string    `gorm:"column:identity_card_front" json:"identity_card_front"`     // 身份证正面
	IdentityCardLater   string    `gorm:"column:identity_card_later" json:"identity_card_later"`     // 身份证反面
	Type                int       `gorm:"column:type;default:0" json:"type"`                         // 0普通店铺 1优选好店
	MobileBackgroundPic string    `gorm:"column:mobile_background_pic" json:"mobile_background_pic"` // 店铺移动端背景图
	PcBackgroundPic     string    `gorm:"column:pc_background_pic" json:"pc_background_pic"`         // 店铺pc背景图
	MerchantName        string    `gorm:"column:merchant_name" json:"merchant_name"`                 // 商家名称
	Email               string    `gorm:"column:email" json:"email"`                                 // 邮箱
	ContractStartTime   LocalTime `gorm:"column:contract_start_time" json:"contract_start_time"`     // 签约起始时间
	ContractEndTime     LocalTime `gorm:"column:contract_end_time" json:"contract_end_time"`         // 签约终止时间
}

func (m *ShopDetail) TableName() string {
	return "tz_shop_detail"
}
