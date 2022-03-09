package InvoiceAddr

import (
	"b2b-api-pc/App/Cores/mysql"
	"b2b-api-pc/App/Model"
	"b2b-api-pc/App/Tool"
	"github.com/jinzhu/gorm"
	"time"
)

type TableStruct struct {
	*Model.InvoiceAddr //结构体继承
}

func Get(maps interface{}) (table []TableStruct) {
	mysql.Db.Model(&TableStruct{}).Where(maps).Find(&table)
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
	mysql.Db.Model(&Model.InvoiceAddr{}).Where(maps).Count(&count)
	return
}

func Add(data map[string]interface{}) bool {
	mysql.Db.Model(&TableStruct{}).Create(&data)
	return !mysql.Db.NewRecord(&data)
}
func EditId(id int, data interface{}) bool {
	mysql.Db.Model(&TableStruct{}).Where("id = ?", id).Updates(data)
	return true
}

func EditMap(maps interface{}, data interface{}) bool {
	mysql.Db.Model(&TableStruct{}).Where(maps).Updates(&data)
	return true
}

func Edit(data Model.InvoiceAddr) bool {
	mysql.Db.Model(&TableStruct{}).Updates(&data)
	return true
}

func (i *TableStruct) BeforeCreate(scope *gorm.Scope) error {
	ID, err := Tool.NewWorker(1)
	if err != nil {
		panic(err)
	}
	err = scope.SetColumn("InvoiceAddrId", ID.GetId())
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
