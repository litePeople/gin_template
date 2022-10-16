package models

import (
	"gin_template/modules/gormx"
)

// Enclosure 存储桶的文件
type Enclosure struct {
	Id int64 `json:"id" gorm:"column:id;AUTO_INCREMENT;primary_key"`
	// 文件的md5
	Md5 string `json:"md5" gorm:"column:md5;type:varchar(255)"`
	// 文件名
	Name string `json:"name" gorm:"column:name;type:varchar(255)"`
	// 扩展名
	ExtName string `json:"extName" gorm:"column:ext_name;type:varchar(255)"`
	// oss的路径
	OssPath string `json:"ossPath" gorm:"column:oss_path;type:varchar(255)"`
	// 违规检测的单据id
	TraceId string `json:"-" gorm:"column:trace_id;type:varchar(255)"`
	// 违规检测的状态
	// 0-未提交违规检测
	// 1-不合规
	// 2-合规
	TraceState int8 `json:"traceState" gorm:"column:trace_state;type:tinyint"`
	// 违规检测的信息
	TraceInfo string `json:"traceLabel" gorm:"column:trace_label;type:varchar(300)"`
	// 是否已经删除
	IsDeleted bool `json:"isDeleted" gorm:"column:is_deleted"`
	// 注册时间
	CreateTime int64 `json:"createTime" gorm:"column:create_time;autoCreateTime:milli"`
	// 更新时间
	UpdateTime int64 `json:"updateTime" gorm:"column:update_time;autoUpdateTime:milli"`
	// 删除时间
	DeleteTime int64 `json:"deleteTime" gorm:"column:delete_time"`
}

func (e *Enclosure) TableName() string {
	return EnclosureColStr()
}

func EnclosureColStr() string {
	return gormx.TablePrefix() + "enclosure"
}
