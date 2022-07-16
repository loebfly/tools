package math

import (
	"math/rand"
)

type randomMath struct{}

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

// SrcEnNumber  英文数字源
func (rm randomMath) SrcEnNumber() string {
	return "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz0123456789"
}

// SrcNumber 数字源
func (rm randomMath) SrcNumber() string {
	return "0123456789"
}

// SrcEn 英文源
func (rm randomMath) SrcEn() string {
	return "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
}

// SrcEnUpper 字母大写英文源
func (rm randomMath) SrcEnUpper() string {
	return "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
}

// SrcEnLower 字母小写英文源
func (rm randomMath) SrcEnLower() string {
	return "abcdefghijklmnopqrstuvwxyz"
}

// Generate 创建一个随机字符串
func (rm randomMath) Generate(src string, length int) string {
	var chars = []rune(src)
	var b = make([]rune, length)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}
