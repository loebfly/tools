package math

import (
	"math/rand"
	"time"
)

type randomMath struct{}

// Generate 创建一个随机字符串
func (rm randomMath) Generate(src string, length int) string {
	bytes := []byte(src)
	result := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// IntFromRange 在某个区间内返回一个随机整数
func (rm randomMath) IntFromRange(min, max int) int {
	return min + rand.Intn(max-min)
}

// Float64FromRange 在某个区间内返回一个随机浮点数
func (rm randomMath) Float64FromRange(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// Bool 返回随机布尔值
func (rm randomMath) Bool() bool {
	return rand.Intn(2) == 0
}

// EnNumber 创建一个随机英文数字字符串
func (rm randomMath) EnNumber(length int) string {
	src := "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz0123456789"
	return rm.Generate(src, length)
}

// Number 创建一个随机数字字符串
func (rm randomMath) Number(length int) string {
	src := "0123456789"
	return rm.Generate(src, length)
}

// En 创建一个随机英文字符串
func (rm randomMath) En(length int) string {
	src := "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	return rm.Generate(src, length)
}

// EnUpper 创建一个随机大写英文字符串
func (rm randomMath) EnUpper(length int) string {
	src := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return rm.Generate(src, length)
}

// EnLower 创建一个随机小写英文字符串
func (rm randomMath) EnLower(length int) string {
	src := "abcdefghijklmnopqrstuvwxyz"
	return rm.Generate(src, length)
}
