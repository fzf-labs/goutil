package timeutil

import (
	"database/sql"
	"time"

	"github.com/golang-module/carbon/v2"
)

// Carbon 获取Carbon对象
func Carbon() *carbon.Carbon {
	return carbon.NewCarbon().SetTimezone(carbon.PRC)
}

// TimeToCarbon 时间转Carbon
func TimeToCarbon(t time.Time) *carbon.Carbon {
	return carbon.CreateFromStdTime(t, carbon.PRC)
}

// StrToCarbon 字符串转Carbon
func StrToCarbon(str string) *carbon.Carbon {
	return carbon.Parse(str, carbon.PRC)
}

// StrToTime 字符串转时间
func StrToTime(str string) time.Time {
	return carbon.Parse(str, carbon.PRC).StdTime()
}

// NowCarbon 获取当前时间的Carbon对象
func NowCarbon() *carbon.Carbon {
	return carbon.Now(carbon.PRC)
}

// NowTime 获取当前时间
func NowTime() time.Time {
	return carbon.Now().StdTime()
}

// NowUnix 获取当前时间戳
func NowUnix() int64 {
	return carbon.Now().Timestamp()
}

// NowString 转换为当前时间 2021-06-29 23:53:32
func NowString() string {
	return carbon.Now().ToDateTimeString()
}

// NowMillisecondString 转换为当前时间 2021-06-29 23:53:32.010
func NowMillisecondString() string {
	return carbon.Now().ToDateTimeMilliString()
}

// NowMicrosecondString 转换为当前时间 2021-06-29 23:53:32.100000
func NowMicrosecondString() string {
	return carbon.Now().ToDateTimeMicroString()
}

// NowSQLNullTime 获取当前时间的sql.NullTime
func NowSQLNullTime() sql.NullTime {
	return sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
}

// TimeToDateTimeString 时间转日期 eg:2021-06-29 23:53:32
func TimeToDateTimeString(t time.Time) string {
	return carbon.CreateFromStdTime(t).ToDateTimeString()
}

// TimeToSQLNullTime 将time.Time转换为sql.NullTime
func TimeToSQLNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}

// TimeToDiffForHumans 时间差转换为人类友好时间
func TimeToDiffForHumans(t time.Time) string {
	return carbon.CreateFromStdTime(t, carbon.PRC).SetLocale("zh-CN").DiffForHumans()
}

// RFC3339 格式化时间 RFC3339
func RFC3339(tt time.Time) string {
	if tt.IsZero() {
		return ""
	}
	return tt.Format(time.RFC3339)
}
