package slice

import "strconv"

type conv struct{}

// StringToI 将字符串切片转换为interface切片
func (c conv) StringToI(list []string) []interface{} {
	result := make([]interface{}, 0)
	for _, v := range list {
		result = append(result, v)
	}
	return result
}

// IntToI 将整数切片转换为interface切片
func (c conv) IntToI(list []int) []interface{} {
	result := make([]interface{}, 0)
	for _, v := range list {
		result = append(result, v)
	}
	return result
}

// IntToS 将interface切片转换为字符串切片
func (c conv) IntToS(list []int) []string {
	result := make([]string, 0)
	for _, v := range list {
		result = append(result, strconv.Itoa(v))
	}
	return result
}

// Int8ToI 将int8切片转换为interface切片
func (c conv) Int8ToI(list []int8) []interface{} {
	result := make([]interface{}, 0)
	for _, v := range list {
		result = append(result, v)
	}
	return result
}

// Int8ToS 将int8切片转换为字符串切片
func (c conv) Int8ToS(list []int8) []string {
	result := make([]string, 0)
	for _, v := range list {
		result = append(result, strconv.FormatInt(int64(v), 10))
	}
	return result
}

// Int32ToI 将int32切片转换为interface切片
func (c conv) Int32ToI(list []int32) []interface{} {
	result := make([]interface{}, 0)
	for _, v := range list {
		result = append(result, v)
	}
	return result
}

// Int32ToS 将int32切片转换为字符串切片
func (c conv) Int32ToS(list []int32) []string {
	result := make([]string, 0)
	for _, v := range list {
		result = append(result, strconv.FormatInt(int64(v), 10))
	}
	return result
}

// Int64ToI 将int64切片转换为interface切片
func (c conv) Int64ToI(list []int64) []interface{} {
	result := make([]interface{}, 0)
	for _, v := range list {
		result = append(result, v)
	}
	return result
}

// Int64ToS 将int64切片转换为字符串切片
func (c conv) Int64ToS(list []int64) []string {
	result := make([]string, 0)
	for _, v := range list {
		result = append(result, strconv.FormatInt(v, 10))
	}
	return result
}

// BoolToI 将bool切片转换为interface切片
func (c conv) BoolToI(list []bool) []interface{} {
	result := make([]interface{}, 0)
	for _, v := range list {
		result = append(result, v)
	}
	return result
}

// Float32ToI 将float32切片转换为interface切片
func (c conv) Float32ToI(list []float32) []interface{} {
	result := make([]interface{}, 0)
	for _, v := range list {
		result = append(result, v)
	}
	return result
}

// Float32ToS 将float32切片转换为字符串切片
func (c conv) Float32ToS(list []float32) []string {
	result := make([]string, 0)
	for _, v := range list {
		result = append(result, strconv.FormatFloat(float64(v), 'f', -1, 32))
	}
	return result
}

// Float64ToI 将float64切片转换为interface切片
func (c conv) Float64ToI(list []float64) []interface{} {
	result := make([]interface{}, 0)
	for _, v := range list {
		result = append(result, v)
	}
	return result
}

// Float64ToS 将float64切片转换为字符串切片
func (c conv) Float64ToS(list []float64) []string {
	result := make([]string, 0)
	for _, v := range list {
		result = append(result, strconv.FormatFloat(v, 'f', -1, 64))
	}
	return result
}
