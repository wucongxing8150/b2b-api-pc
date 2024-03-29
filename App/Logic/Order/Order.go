package Order

import (
	"b2b-api-pc/App/Cores/mysql"
	"b2b-api-pc/App/Model"
	"fmt"
)

type TableStruct struct {
	*Model.Order // 结构体继承
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
func GetWhere(maps interface{}) (table []TableStruct) {
	db, err := Model.BuildWhere(mysql.Init(), maps)
	if err != nil {
		fmt.Println("查询失败", err)
		return nil
	}

	db.Model(&TableStruct{}).Find(&table)
	return table
}

func Add(data Model.Order) (string, bool) {
	if err := mysql.Db.Model(&TableStruct{}).Create(&data).Error; err != nil {
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

func Edit(data Model.Order) bool {
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
