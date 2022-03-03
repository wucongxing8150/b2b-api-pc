// @Description: mysql初始化连接
// @Author: wucongxing
// @Date:2021/12/24 14:00

package mysql

import (
	"fmt"

	"b2b-api-pc/App/Config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {
	var err error

	DB, err = gorm.Open("mysql", Config.C.Mysql.Dns())
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	// 连接数配置也可以写入配置，在此读取
	DB.DB().SetMaxIdleConns(Config.C.Mysql.MaxIdleConns)
	DB.DB().SetMaxOpenConns(Config.C.Mysql.MaxOpenConns)

	// 调试模式
	DB.LogMode(Config.C.Mysql.Debug) // 打印sql
	DB.SingularTable(true)           // 全局禁用表名复数
	fmt.Println("初始化数据库成功......")
	return DB
}
