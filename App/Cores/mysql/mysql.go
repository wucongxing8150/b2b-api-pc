// @Description: mysql初始化连接
// @Author: wucongxing
// @Date:2021/12/24 14:00

package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

type Mysql struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`                           // 服务器地址
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`                           // 端口
	DbName       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`                  // 数据库名
	Username     string `mapstructure:"username" json:"username" yaml:"username"`               // 数据库用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`               // 数据库密码
	MaxIdleConns int    `mapstructure:"max-idle-cons" json:"MaxIdleConns" yaml:"max-idle-cons"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-cons" json:"MaxOpenConns" yaml:"max-open-cons"` // 打开到数据库的最大连接数
	Debug        bool   `mapstructure:"debug" json:"debug" yaml:"debug"`                        // 是否开启Gorm全局日志
}

func Init() *gorm.DB {
	var err error

	m := &Mysql{}

	if err := viper.UnmarshalKey("mysql", m); err != nil {
		panic(fmt.Errorf("解析配置文件失败, err:%s \n", err))
	}

	dns := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", m.Username,
		m.Password, m.Host, m.Port, m.DbName, "10s")

	DB, err = gorm.Open("mysql", dns)

	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	// 连接数配置也可以写入配置，在此读取
	DB.DB().SetMaxIdleConns(m.MaxIdleConns)
	DB.DB().SetMaxOpenConns(m.MaxOpenConns)

	// 调试模式
	DB.LogMode(m.Debug)    // 打印sql
	DB.SingularTable(true) // 全局禁用表名复数
	fmt.Println("初始化数据库成功......")
	return DB
}
