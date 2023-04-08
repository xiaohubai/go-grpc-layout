// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UID        string    `gorm:"column:uid;not null" json:"uid"`
	Username   string    `gorm:"column:username;not null" json:"username"`
	Nickname   string    `gorm:"column:nickname;not null" json:"nickname"`
	Motto      string    `gorm:"column:motto;not null" json:"motto"`
	Password   string    `gorm:"column:password;not null" json:"password"`
	Salt       string    `gorm:"column:salt;not null" json:"salt"`
	Birth      time.Time `gorm:"column:birth;not null;default:2006-01-02 00:00:00" json:"birth"`
	Avatar     string    `gorm:"column:avatar;not null;default:avatar.jpg" json:"avatar"`
	RoleID     string    `gorm:"column:role_id;not null" json:"role_id"`
	RoleName   string    `gorm:"column:role_name;not null" json:"role_name"`
	Phone      string    `gorm:"column:phone;not null" json:"phone"`
	Wechat     string    `gorm:"column:wechat;not null" json:"wechat"`
	Email      string    `gorm:"column:email;not null" json:"email"`
	State      int32     `gorm:"column:state;not null" json:"state"`
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP" json:"update_time"`
	DeleteTime time.Time `gorm:"column:delete_time" json:"delete_time"`
	CreateUser string    `gorm:"column:create_user;not null" json:"create_user"`
	UpdateUser string    `gorm:"column:update_user;not null" json:"update_user"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}