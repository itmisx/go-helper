package helper

import (
	"strings"
	"time"
)

type Datetime struct {
	timestamp int64
	timezone  *int
}

// WithTimestamp 待转换的时间戳
func (dt *Datetime) WithTimestamp(timestamp int64) *Datetime {
	dt.timestamp = timestamp
	return dt
}

// WithTimezone 相对时区
func (dt *Datetime) WithTimezone(timezone int) *Datetime {
	if timezone > 12 {
		timezone = 12
	} else if timezone < -12 {
		timezone = -12
	}
	dt.timezone = &timezone
	return dt
}

// ToDate 时间日期格式化显示
// format 格式化  Y-m-d H:i:s
// time 时间戳，0为当前时间戳
// timezone 时区，nil为系统默认时区
func (dt *Datetime) ToDate(format string) string {
	if dt.timestamp == 0 {
		dt.timestamp = time.Now().Unix()
	}
	// 格式化替换为go内置的格式化字符串
	format = strings.ReplaceAll(format, "Y", "2006")
	format = strings.ReplaceAll(format, "m", "1")
	format = strings.ReplaceAll(format, "d", "02")
	format = strings.ReplaceAll(format, "H", "15")
	format = strings.ReplaceAll(format, "i", "04")
	format = strings.ReplaceAll(format, "s", "05")
	// 时间戳转time.Time
	tm := time.Unix(dt.timestamp, 0)
	// 根据时区进行转换
	if dt.timezone != nil {
		cstZone := time.FixedZone("UTC", *dt.timezone*3600)
		return tm.In(cstZone).Format(format)
	}
	return tm.Format(format)
}

// ToTime 时间转为时间戳
func (dt *Datetime) ToTime(value string) (timestamp int64, err error) {
	layout := "2006-01-02 15:04:05"
	var lc *time.Location
	if dt.timezone != nil {
		lc = time.FixedZone("UTC", *dt.timezone*3600)
	} else {
		lc = time.Local
	}
	tm, err := time.ParseInLocation(layout, value, lc)
	if err != nil {
		return 0, err
	}
	return tm.Unix(), nil
}
