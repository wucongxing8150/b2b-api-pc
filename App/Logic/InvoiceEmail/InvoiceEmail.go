package InvoiceEmail

import (
	"b2b-api-pc/App/Cores/mysql"
	"b2b-api-pc/App/Model"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type TableStruct struct {
	*Model.InvoiceEmail // 结构体继承
}

// var Db = mysql.Init()

func Get(maps interface{}) (table []TableStruct) {
	Db := mysql.Init()
	Db.Model(&TableStruct{}).Where(maps).Find(&table)
	return
}

func GetWhere(maps interface{}) (table []TableStruct) {
	db, err := Model.BuildWhere(mysql.Init(), maps)
	if err != nil {
		fmt.Println("插入失败", err)
		return nil
	}

	db.Model(&TableStruct{}).Find(&table)
	return table
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
func GetTotal(maps interface{}) (count int) {
	mysql.Db.Model(&TableStruct{}).Where(maps).Count(&count)
	return
}

func Add(data Model.InvoiceEmail) (string, bool) {
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

func Edit(data Model.InvoiceEmail) bool {
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

func (i *TableStruct) BeforeCreate(scope *gorm.Scope) error {
	// ID, err := Tool.NewWorker(1)
	// if err != nil {
	// 	panic(err)
	// }
	// err = scope.SetColumn("InvoiceAddrId", ID.GetId())
	// if err != nil {
	// 	panic(err)
	// }

	err := scope.SetColumn("CreateTime", time.Now())
	if err != nil {
		panic(err)
	}

	err = scope.SetColumn("UpdateTime", time.Now())
	if err != nil {
		panic(err)
	}
	return nil
}

func (i *TableStruct) BeforeUpdate(scope *gorm.Scope) (err error) {
	err = scope.SetColumn("UpdateTime", time.Now())
	if err != nil {
		panic(err)
	}
	return nil
}
