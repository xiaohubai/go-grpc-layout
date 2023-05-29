package data

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/xiaohubai/go-grpc-layout/configs/conf"
	"github.com/xiaohubai/go-grpc-layout/internal/biz"
	"github.com/xiaohubai/go-grpc-layout/internal/consts"
	"github.com/xiaohubai/go-grpc-layout/internal/data/gen"
	"github.com/xiaohubai/go-grpc-layout/internal/data/model"
	pelasticsearch "github.com/xiaohubai/go-grpc-layout/pkg/elasticsearch"
	pgorm "github.com/xiaohubai/go-grpc-layout/pkg/gorm"
	predis "github.com/xiaohubai/go-grpc-layout/pkg/redis"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDataRepo)

type dataRepo struct {
	data *Data
	log  *log.Helper
}

func NewDataRepo(d *Data, lg log.Logger) biz.Repo {
	return &dataRepo{
		data: d,
		log:  log.NewHelper(lg),
	}
}

// Data .
type Data struct {
	db  *gen.Query
	rdb *redis.Client
	es  *elasticsearch.Client
}

// NewData .
func NewData(c *conf.Data, logg log.Logger) (*Data, error) {
	db, err := pgorm.NewClient(c.Mysql)
	if err != nil {
		panic(fmt.Errorf("MySQL启动异常: %s", err))
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	if err := AutoMigrate(db); err != nil {
		panic(fmt.Errorf("autoMigrate failed: %s", err))
	}
	rdb, err := predis.NewClient(c.Redis)
	if err != nil {
		panic(fmt.Errorf("redis connect ping failed: %s", err))
	}
	es, err := pelasticsearch.NewClient(c.Es)
	if err != nil {
		panic(fmt.Errorf("elasticsearch connect failed: %s", err))
	}

	consts.DB = db
	consts.RDB = rdb.Client
	consts.ES = es
	return &Data{db: gen.Use(db), rdb: rdb.Client, es: es}, nil
}

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.User{}); err != nil {
		return err
	}
	return nil
}
