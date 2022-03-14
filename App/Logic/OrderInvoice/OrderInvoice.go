package OrderInvoice

import (
	"fmt"

	"b2b-api-pc/App/Cores/mysql"
	"b2b-api-pc/App/Model"
	"gorm.io/gorm"
)

type TableStruct struct {
	*Model.OrderInvoice // 结构体继承
}

// var Db = mysql.Init()

func Get(maps interface{}) (table []TableStruct) {
	Db := mysql.Init()
	Db.Model(&TableStruct{}).Where(maps).Find(&table)
	return
}

func GetId(id int64) (table []TableStruct) {
	Db := mysql.Init()
	Db.Model(&TableStruct{}).Find(&table, id)
	return
}

func Search(maps interface{}) (table []TableStruct) {
	mysql.Db.Model(&TableStruct{}).Where(maps).Find(&table)
	return
}
func SearchPage(pageNum int, pageSize int, maps interface{}) (table []TableStruct) {
	mysql.Db.Model(&TableStruct{}).Where(maps).Offset(pageNum).Limit(pageSize).Find(&table)
	return
}
func GetTotal(maps interface{}) (count int64) {
	mysql.Db.Model(&TableStruct{}).Where(maps).Count(&count)
	return
}

func Add(data Model.OrderInvoice) (string, bool) {
	if err := mysql.Db.Model(&TableStruct{}).Create(&data).Error; err != nil {
		fmt.Println("插入失败", err)
		return err.Error(), false
	}
	return "", true
}

func AddMap(tx *gorm.DB, data interface{}) (string, bool) {
	if err := tx.Model(&TableStruct{}).Create(&data).Error; err != nil {
		fmt.Println("插入失败", err)
		return err.Error(), false
	}
	return "", true
}

func EditId(id int, data interface{}) bool {
	mysql.Db.Model(&TableStruct{}).Where("id = ?", id).Updates(data)
	return true
}

func EditMap(maps interface{}, data interface{}) bool {
	mysql.Db.Model(&TableStruct{}).Where(maps).Updates(&data)
	return true
}

func Edit(data Model.OrderInvoice) bool {
	mysql.Db.Model(&TableStruct{}).Updates(&data)
	return true
}

// DeleteMap 条件删除
func DeleteMap(maps interface{}) bool {
	mysql.Db.Model(&TableStruct{}).Where(maps).Delete(&TableStruct{})
	return true
}

// DeleteId 主键删除
func DeleteId(id int64) bool {
	mysql.Db.Model(&TableStruct{}).Delete(&TableStruct{}, id)
	return true
}
