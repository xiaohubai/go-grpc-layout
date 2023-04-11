// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMenu = "menu"

// Menu mapped from table <menu>
type Menu struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Path       string    `gorm:"column:path;not null" json:"path"`                                         // 路由
	Name       string    `gorm:"column:name;not null" json:"name"`                                         // 名称
	Redirect   string    `gorm:"column:redirect;not null" json:"redirect"`                                 // 重定向
	Component  string    `gorm:"column:component;not null" json:"component"`                               // 文件地址
	ParentID   int32     `gorm:"column:parentId;not null" json:"parentId"`                                 // 父id
	RoleIDs    string    `gorm:"column:roleIDs;not null" json:"roleIDs"`                                   // 角色组
	Title      string    `gorm:"column:title;not null" json:"title"`                                       // 标题
	Icon       string    `gorm:"column:icon;not null" json:"icon"`                                         // 图标
	Hidden     bool      `gorm:"column:hidden;not null" json:"hidden"`                                     // 是否隐藏
	KeepAlive  bool      `gorm:"column:keepAlive;not null;default:1" json:"keepAlive"`                     // keepAlive
	Sort       int32     `gorm:"column:sort;not null" json:"sort"`                                         // 排序
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间
	UpdateTime time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP" json:"update_time"` // 记录修改时间
	DeleteTime time.Time `gorm:"column:delete_time" json:"delete_time"`                                    // 删除时间
	CreateUser string    `gorm:"column:create_user;not null" json:"create_user"`                           // 创建人
	UpdateUser string    `gorm:"column:update_user;not null" json:"update_user"`                           // 修改人
}

// TableName Menu's table name
func (*Menu) TableName() string {
	return TableNameMenu
}
