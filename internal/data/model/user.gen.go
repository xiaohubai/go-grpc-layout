// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID         int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UID        string         `gorm:"column:uid;not null;comment:uid" json:"uid"`
	Username   string         `gorm:"column:username;not null;comment:用户名" json:"username"`
	Nickname   string         `gorm:"column:nickname;not null;comment:昵称" json:"nickname"`
	Motto      string         `gorm:"column:motto;not null;comment:座右铭" json:"motto"`
	Password   string         `gorm:"column:password;not null;comment:密码" json:"password"`
	Salt       string         `gorm:"column:salt;not null;comment:加盐" json:"salt"`
	Birth      time.Time      `gorm:"column:birth;not null;default:2006-01-02;comment:出生日期" json:"birth"`
	Avatar     string         `gorm:"column:avatar;not null;default:avatar.jpg;comment:头像" json:"avatar"`
	RoleID     string         `gorm:"column:role_id;not null;comment:角色Id" json:"role_id"`
	RoleName   string         `gorm:"column:role_name;not null;comment:角色名称" json:"role_name"`
	Phone      string         `gorm:"column:phone;not null;comment:手机号" json:"phone"`
	Wechat     string         `gorm:"column:wechat;not null;comment:微信号" json:"wechat"`
	Email      string         `gorm:"column:email;not null;comment:邮箱" json:"email"`
	State      int32          `gorm:"column:state;not null;comment:用户状态:(0:初始,1:使用,2:停用,3:删除)" json:"state"`
	CreateAt   time.Time      `gorm:"column:create_at;not null;default:CURRENT_TIMESTAMP;comment:记录创建时间" json:"create_at"`
	UpdateAt   time.Time      `gorm:"column:update_at;not null;default:CURRENT_TIMESTAMP;comment:记录修改时间" json:"update_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	CreateUser string         `gorm:"column:create_user;not null;comment:创建人" json:"create_user"`
	UpdateUser string         `gorm:"column:update_user;not null;comment:修改人" json:"update_user"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
