package math

import "math/rand"

type randomMath struct{}

// Int 在某个区间内返回一个随机整数
func (rm randomMath) Int(min, max int) int {
	return min + rand.Intn(max-min)
}

// Float64 在某个区间内返回一个随机浮点数
func (rm randomMath) Float64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// String 返回指定长度的大小写数子组合随机字符串
func (rm randomMath) String(length int) string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var b = make([]rune, length)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func (rm randomMath) Bool() bool {
	return rand.Intn(2) == 0
}
