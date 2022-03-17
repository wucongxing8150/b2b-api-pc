// @Description: 文件描述
// @Author: wucongxing
// @Date:2022/3/17 10:30

package order

// StatusToZh
// @Description: 订单状态转变为中文
func StatusToZh(status int) string {
	var zhStatus string
	switch status {
	case 1:
		zhStatus = "待付款"
	case 2:
		zhStatus = "待发货"
	case 3:
		zhStatus = "待收货"
	case 4:
		zhStatus = "待评价"
	case 5:
		zhStatus = "成功"
	case 6:
		zhStatus = "失败"
	case 7:
		zhStatus = "待成团"
	case 8:
		zhStatus = "待审核"
	default:
		zhStatus = "未知"
	}
	return zhStatus
}
