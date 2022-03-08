package main

import (
	"fmt"
	"strconv"

	"b2b-api-pc/App/Api/Router"
	"b2b-api-pc/App/Cores"
	"github.com/spf13/viper"
)

func main() {
	// 初始化路由
	r := Router.Init()

	// 加载核心
	Cores.Init()

	if err := r.Run(":" + strconv.Itoa(viper.GetInt("port"))); err != nil {
		fmt.Println("启动失败:%v\n", err)
	}

}
