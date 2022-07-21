package stringT

import "strings"

// Sub 以char的方式进行字符串截取
// start:开始位置，strLength:截取长度
// 如果 start < 0, 则从字符串末尾开始计算
// 如果 strLength <= 0, 则截取到字符串末尾
func (e Enter) Sub(src string, start int, length ...int) string {
	charList := []rune(src)
	l := len(charList)
	step := 0
	end := 0

	if len(length) == 0 {
		step = l
	} else {
		step = length[0]
	}

	if start < 0 {
		start = l + start
	}
	end = start + step

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}

	if start > l {
		start = l
	}

	if end < 0 {
		end = 0
	}

	if end > l {
		end = l
	}
	return string(charList[start:end])
}

// Len 返回char字符串长度
func (e Enter) Len(src string) int {
	return len([]rune(src))
}

// Before 获取某个字符串第一个出现位置之前的字符串，如果不存在返回源字符串
func (e Enter) Before(src string, target string) string {
	if target == "" {
		return src
	}
	i := strings.Index(src, target)
	if i != -1 {
		return src[:i]
	}
	return src
}

// BeforeLast 获取某个字符串最后出现位置之前的字符串，如果不存在返回源字符串
func (e Enter) BeforeLast(src string, target string) string {
	if target == "" {
		return src
	}
	i := strings.LastIndex(src, target)
	if i != -1 {
		return src[:i]
	}
	return src
}

// After 获取某个字符串第一个出现位置之后的字符串，如果不存在返回源字符串
func (e Enter) After(src string, target string) string {
	if target == "" {
		return src
	}
	i := strings.Index(src, target)
	if i != -1 {
		return src[i+len(target):]
	}
	return src
}

// AfterLast 获取某个字符串最后出现位置之后的字符串，如果不存在返回源字符串
func (e Enter) AfterLast(src string, target string) string {
	if target == "" {
		return src
	}
	i := strings.LastIndex(src, target)
	if i != -1 {
		return src[i+len(target):]
	}
	return src
}
