package timeT

import (
	"time"
)

type timeCalculate struct{}

// AddSeconds 增加秒数
func (tc timeCalculate) AddSeconds(t time.Time, seconds int) time.Time {
	return t.Add(time.Duration(seconds) * time.Second)
}

// AddMinutes 增加分钟数
func (tc timeCalculate) AddMinutes(t time.Time, minutes int) time.Time {
	return t.Add(time.Duration(minutes) * time.Minute)
}

// AddHours 增加小时数
func (tc timeCalculate) AddHours(t time.Time, hours int) time.Time {
	return t.Add(time.Duration(hours) * time.Hour)
}

// AddDays 增加天数
func (tc timeCalculate) AddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

// AddMonths 增加月数
func (tc timeCalculate) AddMonths(t time.Time, months int) time.Time {
	return t.AddDate(0, months, 0)
}

// AddYears 增加年数
func (tc timeCalculate) AddYears(t time.Time, years int) time.Time {
	return t.AddDate(years, 0, 0)
}

// SubtractSeconds 减去秒数
func (tc timeCalculate) SubtractSeconds(t time.Time, seconds int) time.Time {
	return t.Add(time.Duration(-seconds) * time.Second)
}

// SubtractMinutes 减去分钟数
func (tc timeCalculate) SubtractMinutes(t time.Time, minutes int) time.Time {
	return t.Add(time.Duration(-minutes) * time.Minute)
}

// SubtractHours 减去小时数
func (tc timeCalculate) SubtractHours(t time.Time, hours int) time.Time {
	return t.Add(time.Duration(-hours) * time.Hour)
}

// SubtractDays 减去天数
func (tc timeCalculate) SubtractDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, -days)
}

// SubtractMonths 减去月数
func (tc timeCalculate) SubtractMonths(t time.Time, months int) time.Time {
	return t.AddDate(0, -months, 0)
}

// SubtractYears 减去年数
func (tc timeCalculate) SubtractYears(t time.Time, years int) time.Time {
	return t.AddDate(-years, 0, 0)
}

// GetSeconds 获取秒数
func (tc timeCalculate) GetSeconds(t time.Time) int {
	return int(t.Unix())
}

// GetMinutes 获取分钟数
func (tc timeCalculate) GetMinutes(t time.Time) int {
	return int(t.Unix() / 60)
}

// GetHours 获取小时数
func (tc timeCalculate) GetHours(t time.Time) int {
	return int(t.Unix() / 3600)
}

// GetDays 获取天数
func (tc timeCalculate) GetDays(t time.Time) int {
	return int(t.Unix() / 86400)
}

// GetMonths 获取月数
func (tc timeCalculate) GetMonths(t time.Time) int {
	return int(t.Unix() / 2592000)
}

// GetYears 获取年数
func (tc timeCalculate) GetYears(t time.Time) int {
	return int(t.Unix() / 31536000)
}

// GetWeekday 获取星期
func (tc timeCalculate) GetWeekday(t time.Time) int {
	return int(t.Weekday())
}

// GetMonth 获取月份
func (tc timeCalculate) GetMonth(t time.Time) int {
	return int(t.Month())
}

// GetYear 获取年份
func (tc timeCalculate) GetYear(t time.Time) int {
	return t.Year()
}

// GetUnixSecond 获取时间戳秒数
func (tc timeCalculate) GetUnixSecond(unix int64) int {
	return int(unix)
}

// GetUnixMinute 获取时间戳分钟数
func (tc timeCalculate) GetUnixMinute(unix int64) int {
	return int(unix / 60)
}

// GetUnixHour 获取时间戳小时数
func (tc timeCalculate) GetUnixHour(unix int64) int {
	return int(unix / 3600)
}

// GetUnixDay 获取时间戳天数
func (tc timeCalculate) GetUnixDay(unix int64) int {
	return int(unix / 86400)
}

// GetUnixMonth 获取时间戳月数
func (tc timeCalculate) GetUnixMonth(unix int64) int {
	return int(unix / 2592000)
}

// GetUnixYear 获取时间戳年数
func (tc timeCalculate) GetUnixYear(unix int64) int {
	return int(unix / 31536000)
}

// DiffSeconds 计算两个时间差秒数
func (tc timeCalculate) DiffSeconds(start, end time.Time) int {
	return int(end.Unix() - start.Unix())
}

// DiffMinutes 计算两个时间差分钟数
func (tc timeCalculate) DiffMinutes(start, end time.Time) int {
	return int(end.Unix()/60 - start.Unix()/60)
}

// DiffHours 计算两个时间差小时数
func (tc timeCalculate) DiffHours(start, end time.Time) int {
	return int(end.Unix()/3600 - start.Unix()/3600)
}

// DiffDays 计算两个时间差天数
func (tc timeCalculate) DiffDays(start, end time.Time) int {
	return int(end.Unix()/86400 - start.Unix()/86400)
}

// DiffMonths 计算两个时间差月数
func (tc timeCalculate) DiffMonths(start, end time.Time) int {
	return int(end.Unix()/2592000 - start.Unix()/2592000)
}

// DiffYears 计算两个时间差年数
func (tc timeCalculate) DiffYears(start, end time.Time) int {
	return int(end.Unix()/31536000 - start.Unix()/31536000)
}
