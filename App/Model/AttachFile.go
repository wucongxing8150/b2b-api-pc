// @Description: 文件描述
// @Author: wucongxing
// @Date:2022/3/17 10:55

package Model

type AttachFile struct {
	FileId            int64     `gorm:"column:file_id;primary_key;AUTO_INCREMENT" json:"file_id"`
	FilePath          string    `gorm:"column:file_path" json:"file_path"`                                 // 文件路径
	FileType          string    `gorm:"column:file_type" json:"file_type"`                                 // 文件类型
	FileName          string    `gorm:"column:file_name" json:"file_name"`                                 // 文件名
	FileSize          int       `gorm:"column:file_size" json:"file_size"`                                 // 文件大小
	UploadTime        LocalTime `gorm:"column:upload_time" json:"upload_time"`                             // 上传时间
	ShopId            int64     `gorm:"column:shop_id" json:"shop_id"`                                     // 店铺id
	Type              int       `gorm:"column:type" json:"type"`                                           // 文件 1:图片 2:视频 3:文件
	AttachFileGroupId int64     `gorm:"column:attach_file_group_id;default:0" json:"attach_file_group_id"` // 文件分组id
}

func (m *AttachFile) TableName() string {
	return "tz_attach_file"
}
