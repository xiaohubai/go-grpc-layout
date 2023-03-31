package dao

import (
	"fmt"

	"github.com/xiaohubai/go-grpc-layout/configs"
	"github.com/xiaohubai/go-grpc-layout/internal/biz"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDataRepo)

type dataRepo struct {
	dao *Dao
	log *log.Helper
}

func NewDataRepo(dao *Dao, lg log.Logger) biz.Repo {
	return &dataRepo{
		dao: dao,
		log: log.NewHelper(lg),
	}
}

// Dao .
type Dao struct {
	db  *gorm.DB
	rdb *redis.Client
}

// NewData .
func NewData(c *configs.Dao, logger log.Logger) (*Dao, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	mysqlConfig := mysql.Config{
		DSN:                       c.Mysql.Source, // DSN data source name
		DefaultStringSize:         191,            // string 类型字段的默认长度
		DisableDatetimePrecision:  true,           // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,           // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,           // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,          // 根据版本自动配置
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, //禁用外键约束
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //禁用表复数形式
			TablePrefix:   "",   //表前缀
		},
	})
	if err != nil {
		panic(fmt.Errorf("MySQL启动异常: %s", err))
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: c.Redis.Password,
		DB:       int(c.Redis.Db),
	})
	_, err = rdb.Ping(rdb.Context()).Result()
	if err != nil {
		panic(fmt.Errorf("redis connect ping failed: %s", err))
	}

	return &Dao{db: db, rdb: rdb}, cleanup, nil
}