// @Description: 文件描述
// @Author: wucongxing
// @Date:2022/3/15 10:58

package Tool

import (
	"time"
)

// BeforeDate
// @Description: 获取N天前时间
// @param day 天数
// @return string Y-m-d H:i:s形式字符串
func BeforeDate(day int) string {

	currentTime := time.Now()
	oldTime := currentTime.AddDate(0, 0, day)           // 若要获取3天前的时间，则应将-2改为-3
	timeString := oldTime.Format("2006-01-02 15:04:05") // 2020-10-17 16:20:20

	return timeString

}
