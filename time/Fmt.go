package timeT

import (
	"fmt"
	"github.com/araddon/dateparse"
	"strings"
	"time"
)

// LayoutDateTime 返回日期时间格式
func (tFmt Enter) LayoutDateTime() string {
	return "2006-01-02 15:04:05"
}

// LayoutDate 返回日期格式
func (tFmt Enter) LayoutDate() string {
	return "2006-01-02"
}

// LayoutTime 返回时间格式
func (tFmt Enter) LayoutTime() string {
	return "15:04:05"
}

func (tFmt Enter) CTSZone() *time.Location {
	return time.FixedZone("CST", 8*3600)
}

// GetNowDateTime 获取当前日期时间字符串
func (tFmt Enter) GetNowDateTime() string {
	return time.Now().In(tFmt.CTSZone()).Format(tFmt.LayoutDateTime())
}

// GetNowDate 获取当前日期字符串
func (tFmt Enter) GetNowDate() string {
	return time.Now().In(tFmt.CTSZone()).Format(tFmt.LayoutDate())
}

// GetNowTime 获取当前时间字符串
func (tFmt Enter) GetNowTime() string {
	return time.Now().In(tFmt.CTSZone()).Format(tFmt.LayoutTime())
}

// GetNowWeekday 获取当前星期字符串
func (tFmt Enter) GetNowWeekday() string {
	return time.Now().In(tFmt.CTSZone()).Weekday().String()
}

// Format 以任意格式获取日期时间字符串
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
func (tFmt Enter) Format(t time.Time, format string) string {
	newFmt := strings.Replace(format, "MMMM", "January", -1)
	newFmt = strings.Replace(newFmt, "MMM", "Jan", -1)
	newFmt = strings.Replace(newFmt, "MM", "01", -1)
	newFmt = strings.Replace(newFmt, "M", "1", -1)
	newFmt = strings.Replace(newFmt, "dddd", "Monday", -1)
	newFmt = strings.Replace(newFmt, "ddd", "Mon", -1)
	newFmt = strings.Replace(newFmt, "dd", "02", -1)
	newFmt = strings.Replace(newFmt, "d", "2", -1)
	newFmt = strings.Replace(newFmt, "yyyy", "2006", -1)
	newFmt = strings.Replace(newFmt, "yy", "06", -1)
	newFmt = strings.Replace(newFmt, "hh", "15", -1)
	newFmt = strings.Replace(newFmt, "HH", "03", -1)
	newFmt = strings.Replace(newFmt, "H", "3", -1)
	newFmt = strings.Replace(newFmt, "mm", "04", -1)
	newFmt = strings.Replace(newFmt, "m", "4", -1)
	newFmt = strings.Replace(newFmt, "ss", "05", -1)
	newFmt = strings.Replace(newFmt, "s", "5", -1)
	newFmt = strings.Replace(newFmt, "tt", "PM", -1)
	newFmt = strings.Replace(newFmt, "ZZZ", "MST", -1)
	newFmt = strings.Replace(newFmt, "Z", "Z07:00", -1)
	return t.Format(newFmt)
}

// Convert 日期时间字符串转换为指定格式的日期时间字符串
func (tFmt Enter) Convert(timeString string, format string) string {
	t, err := dateparse.ParseLocal(timeString)
	if err != nil {
		fmt.Println("ParseLocal error:", err.Error())
		return ""
	}
	return tFmt.Format(t, format)
}
