package Cores

import (
	"b2b-api-pc/App/Cores/mysql"
	"b2b-api-pc/App/Cores/viper"
)

func Init() {
	// 初始化Viper 加载配置 读取如缓存中
	viper.Init()

	// 初始化数据库连接
	mysql.Init()
}
