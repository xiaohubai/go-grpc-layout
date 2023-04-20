// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCasbinRule = "casbin_rule"

// CasbinRule mapped from table <casbin_rule>
type CasbinRule struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Ptype      string    `gorm:"column:ptype" json:"ptype"`
	V0         string    `gorm:"column:v0" json:"v0"`
	V1         string    `gorm:"column:v1" json:"v1"`
	V2         string    `gorm:"column:v2" json:"v2"`
	V3         string    `gorm:"column:v3" json:"v3"`
	V4         string    `gorm:"column:v4" json:"v4"`
	V5         string    `gorm:"column:v5" json:"v5"`
	Desc       string    `gorm:"column:desc;not null" json:"desc"`                                         // 描述
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 记录创建时间
	UpdateTime time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP" json:"update_time"` // 记录修改时间
	DeleteTime time.Time `gorm:"column:delete_time" json:"delete_time"`                                    // 删除时间
	CreateUser string    `gorm:"column:create_user;not null" json:"create_user"`                           // 创建人
	UpdateUser string    `gorm:"column:update_user;not null" json:"update_user"`                           // 修改人
}

// TableName CasbinRule's table name
func (*CasbinRule) TableName() string {
	return TableNameCasbinRule
}
