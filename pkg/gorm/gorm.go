package gorm

import (
	"github.com/xiaohubai/go-grpc-layout/configs/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewClient(c *conf.Data_Mysql) (*gorm.DB, error) {
	mysqlConfig := mysql.Config{
		DSN:                       c.Source, // DSN data source name
		DefaultStringSize:         191,      // string 类型字段的默认长度
		DisableDatetimePrecision:  true,     // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,     // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,     // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,    // 根据版本自动配置
	}
	return gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		//Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true, //禁用外键约束
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //禁用表复数形式
			TablePrefix:   "",   //表前缀
		},
	})
}
