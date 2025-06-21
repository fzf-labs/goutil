package timeutil

import (
	"fmt"
	"strconv"
	"time"
)

// TimeToHuman 格式化人类友好时间
func TimeToHuman(t time.Time) string {
	var res = ""
	if t.IsZero() {
		return res
	}
	tt := time.Now().Unix() - t.Unix()
	data := []map[string]any{
		{"key": 31536000, "value": "年"},
		{"key": 2592000, "value": "个月"},
		{"key": 604800, "value": "星期"},
		{"key": 86400, "value": "天"},
		{"key": 3600, "value": "小时"},
		{"key": 60, "value": "分钟"},
		{"key": 1, "value": "秒"},
	}
	for _, v := range data {
		var c = tt / int64(v["key"].(int))
		if c != 0 {
			suffix := "前"
			if c < 0 {
				suffix = "后"
				c = -c
			}
			res = strconv.Itoa(int(c)) + v["value"].(string) + suffix
			break
		}
	}
	return res
}

// TimeToHumanShow 格式化人类友好展示时间
func TimeToHumanShow(t time.Time) string {
	duration := time.Now().Unix() - t.Unix()
	var timeStr string
	if duration < 60 {
		timeStr = "刚刚发布"
	} else if duration < 3600 {
		timeStr = fmt.Sprintf("%d分钟前更新", duration/60)
	} else if duration < 86400 {
		timeStr = fmt.Sprintf("%d小时前更新", duration/3600)
	} else if duration < 86400*2 {
		timeStr = "昨天更新"
	} else {
		timeStr = time.Unix(t.Unix(), 00).Format("2006.01.02") + "日更新"
	}
	return timeStr
}
