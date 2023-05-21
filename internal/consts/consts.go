package consts

import (
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"github.com/xiaohubai/go-grpc-layout/configs/conf"
)

const (
	D  = 24 * time.Hour
	H  = time.Hour
	M  = time.Minute
	S  = time.Second
	MS = time.Millisecond
	US = time.Microsecond
	NS = time.Nanosecond

	Jan  = 1
	Feb  = 2
	Mar  = 3
	Apr  = 4
	May  = 5
	Jun  = 6
	Jul  = 7
	Aug  = 8
	Sept = 9
	Oct  = 10
	Nov  = 11
	Dec  = 12

	January   = 1
	February  = 2
	March     = 3
	April     = 4
	June      = 6
	July      = 7
	August    = 8
	September = 9
	October   = 10
	November  = 11
	December  = 12
)

var (
	DB    *gorm.DB
	RDB   *redis.Client
	Conf  *conf.Conf
	Viper *viper.Viper
)

var (
	KafkaTopicOperationRecord = "operationRecord"
)

var (
	EmailTitleViperRemoteWatch = "viper remote watch"
	EmailTitleViperLocalWatch  = "viper local watch"
	EmailTitlePanic            = "panic"
	EmailTitlePprof            = "pprof"
	EmailTitleKafkaProducer    = "kafka producer"
	EmailTitleKafkaConsumer    = "kafka consumer"
	PwdPath, _                 = os.Getwd()
)
