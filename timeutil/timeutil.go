package timeutil

import (
	"database/sql"
	"time"
)

// NowSQLNullTime 获取当前时间的sql.NullTime
func NowSQLNullTime() sql.NullTime {
	return sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
}

// TimeToSQLNullTime 将time.Time转换为sql.NullTime
func TimeToSQLNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}

// RFC3339 格式化时间 RFC3339
func RFC3339(tt time.Time) string {
	if tt.IsZero() {
		return ""
	}
	return tt.Format(time.RFC3339)
}
