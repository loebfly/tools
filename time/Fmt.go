package timeT

import (
	"fmt"
	"github.com/araddon/dateparse"
	"strings"
	"time"
)

type timeFmt struct{}

// LayoutDateTime 返回日期时间格式
func (tf timeFmt) LayoutDateTime() string {
	return "2006-01-02 15:04:05"
}

// LayoutDate 返回日期格式
func (tf timeFmt) LayoutDate() string {
	return "2006-01-02"
}

// LayoutTime 返回时间格式
func (tf timeFmt) LayoutTime() string {
	return "15:04:05"
}

func (tf timeFmt) CTSZone() *time.Location {
	return time.FixedZone("CST", 8*3600)
}

// GetNowDateTime 获取当前日期时间字符串
func (tf timeFmt) GetNowDateTime() string {
	return time.Now().In(tf.CTSZone()).Format(tf.LayoutDateTime())
}

// GetNowDate 获取当前日期字符串
func (tf timeFmt) GetNowDate() string {
	return time.Now().In(tf.CTSZone()).Format(tf.LayoutDate())
}

// GetNowTime 获取当前时间字符串
func (tf timeFmt) GetNowTime() string {
	return time.Now().In(tf.CTSZone()).Format(tf.LayoutTime())
}

// GetNowWeekday 获取当前星期字符串
func (tf timeFmt) GetNowWeekday() string {
	return time.Now().In(tf.CTSZone()).Weekday().String()
}

// Get 以任意格式获取日期时间字符串
// MMMM - month - January
// MMM - month - Jan
// MM - month - 01
// M - month - 1, single bit
// DD - day - 02
// D - day 2
// YYYY - year - 2006
// YY - year - 06
// HH - 24 hours - 03
// H - 24 hours - 3
// hh - 12 hours - 03
// h - 12 hours - 3
// mm - minute - 04
// m - minute - 4
// ss - second - 05
// s - second = 5
func (tf timeFmt) Get(t time.Time, format string) string {
	fmt := strings.Replace(format, "MMMM", "January", -1)
	fmt = strings.Replace(fmt, "MMM", "Jan", -1)
	fmt = strings.Replace(fmt, "MM", "01", -1)
	fmt = strings.Replace(fmt, "M", "1", -1)
	fmt = strings.Replace(fmt, "dddd", "Monday", -1)
	fmt = strings.Replace(fmt, "ddd", "Mon", -1)
	fmt = strings.Replace(fmt, "dd", "02", -1)
	fmt = strings.Replace(fmt, "d", "2", -1)
	fmt = strings.Replace(fmt, "yyyy", "2006", -1)
	fmt = strings.Replace(fmt, "yy", "06", -1)
	fmt = strings.Replace(fmt, "hh", "15", -1)
	fmt = strings.Replace(fmt, "HH", "03", -1)
	fmt = strings.Replace(fmt, "H", "3", -1)
	fmt = strings.Replace(fmt, "mm", "04", -1)
	fmt = strings.Replace(fmt, "m", "4", -1)
	fmt = strings.Replace(fmt, "ss", "05", -1)
	fmt = strings.Replace(fmt, "s", "5", -1)
	fmt = strings.Replace(fmt, "tt", "PM", -1)
	fmt = strings.Replace(fmt, "ZZZ", "MST", -1)
	fmt = strings.Replace(fmt, "Z", "Z07:00", -1)
	return t.Format(fmt)
}

// Convert 日期时间字符串转换为指定格式的日期时间字符串
func (tf timeFmt) Convert(timeString string, format string) string {
	t, err := dateparse.ParseLocal(timeString)
	if err != nil {
		fmt.Println("ParseLocal error:", err.Error())
		return ""
	}
	return tf.Get(t, format)
}
