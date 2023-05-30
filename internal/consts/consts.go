package consts

import (
	"os"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
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
	DB  *gorm.DB
	RDB *redis.Client
	ES  *elasticsearch.Client
)

var (
	KafkaTopicOperationRecord = "operationRecord"
)

var (
	PwdPath, _ = os.Getwd()
)

var (
	EmailTitleViperRemoteWatch = "viper remote watch"
	EmailTitleViperLocalWatch  = "viper local watch"
	EmailTitlePanic            = "panic"
	EmailTitlePprof            = "pprof"
	EmailTitleKafkaProducer    = "kafka producer"
	EmailTitleKafkaConsumer    = "kafka consumer"
)

var (
	ESIndexOperationRecord = "operation_record"
)
