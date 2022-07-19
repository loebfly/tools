package stringT

type extendString struct{}

// Substring 以char的方式进行字符串截取
// start:开始位置，strLength:截取长度
// 如果 start < 0, 则从字符串末尾开始计算
// 如果 strLength <= 0, 则截取到字符串末尾
func (es extendString) Substring(src string, start int, length ...int) string {
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
func (es extendString) Len(src string) int {
	return len([]rune(src))
}
