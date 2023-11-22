package utils

import (
	"os"
	"strconv"
	"time"
)

func init() {
	_ = SetTimeZone("Asia/Shanghai")
}

// SetTimeZone 设置地区时间
// eg: Asia/Shanghai
func SetTimeZone(zone string) error {
	location, err := time.LoadLocation(zone)
	if err != nil {
		return err
	}
	return os.Setenv("TZ", location.String())
}

// Timestamp returns the timestamp in seconds(秒).
func Timestamp() int64 {
	return time.Now().UnixNano() / 1e9
}

// TimestampMilli returns the timestamp in milliseconds(毫秒).
func TimestampMilli() int64 {
	return time.Now().UnixNano() / 1e6
}

// TimestampMicro returns the timestamp in microseconds(微秒).
func TimestampMicro() int64 {
	return time.Now().UnixNano() / 1e3
}

// TimestampNano returns the timestamp in nanoseconds(纳秒).
func TimestampNano() int64 {
	return time.Now().UnixNano()
}

// the timestamp in seconds as string.
func TimestampStr() string {
	return strconv.FormatInt(Timestamp(), 10)
}

// the timestamp in milliseconds as string.
func TimestampMilliStr() string {
	return strconv.FormatInt(TimestampMilli(), 10)
}

// Date returns current date in string like "2006-01-02".
func Date() string {
	return time.Now().Format("2006-01-02")
}

// Datetime returns current datetime in string like "2006-01-02 15:04:05".
func Datetime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// ISO8601 returns current datetime in ISO8601 format like "2006-01-02T15:04:05-07:00".
func ISO8601() string {
	return time.Now().Format("2006-01-02T15:04:05-07:00")
}

// RFC822 returns current datetime in RFC822 format like "Mon, 02 Jan 06 15:04 MST".
func RFC822() string {
	return time.Now().Format("Mon, 02 Jan 06 15:04 MST")
}

// the time in string as time like "2022-01-02 18:04:05".
func StrToTime(str string, layout string) time.Time {
	if t, err := time.ParseInLocation(layout, str, time.Local); err == nil {
		return t
	} else {
		return time.Time{}
	}
}

// the t in time as string like "2022-01-02 18:04:05".
func TimeToStr(t time.Time, layout string) string {
	return time.Now().Format(layout)
}
