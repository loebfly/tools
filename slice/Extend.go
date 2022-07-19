package slice

import "reflect"

type extend struct{}

// FilterEmptyS 去除字符串切片的空值
func (e extend) FilterEmptyS(list []string) []string {
	result := make([]string, 0)
	for _, value := range list {
		if value != "" {
			result = append(result, value)
		}
	}
	return result
}

// FilterEmptyI 去除interface切片的空值
func (e extend) FilterEmptyI(list []interface{}) []interface{} {
	result := make([]interface{}, 0)
	for _, value := range list {
		if str, ok := value.(string); ok && str != "" {
			result = append(result, value)
		} else if value != nil {
			result = append(result, value)
		}
	}
	return result
}

// InS 判断字符串是否在字符串切片中
func (e extend) InS(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

// InI 判断interface是否在interface切片中
func (e extend) InI(list []interface{}, value interface{}) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

// UniqueS 去重字符串数组的重复值
func (e extend) UniqueS(list []string) []string {
	result := make([]string, 0)
	for _, value := range list {
		if !e.InS(result, value) {
			result = append(result, value)
		}
	}
	return result
}

// UniqueI 去重interface数组的重复值
func (e extend) UniqueI(list []interface{}) []interface{} {
	result := make([]interface{}, 0)
	for _, value := range list {
		if !e.InI(result, value) {
			result = append(result, value)
		}
	}
	return result
}

// ReverseS 反转字符串切片
func (e extend) ReverseS(list []string) []string {
	result := make([]string, 0)
	for i := len(list) - 1; i >= 0; i-- {
		result = append(result, list[i])
	}
	return result
}

// ReverseI 反转interface切片
func (e extend) ReverseI(list []interface{}) []interface{} {
	result := make([]interface{}, 0)
	for i := len(list) - 1; i >= 0; i-- {
		result = append(result, list[i])
	}
	return result
}

// RemoveS 删除字符串切片中的某个值
func (e extend) RemoveS(list []string, value string) []string {
	result := make([]string, 0)
	for _, v := range list {
		if v != value {
			result = append(result, v)
		}
	}
	return result
}

// RemoveI 删除interface切片中的某个值
func (e extend) RemoveI(list []interface{}, value interface{}) []interface{} {
	result := make([]interface{}, 0)
	for _, v := range list {
		if v != value {
			result = append(result, v)
		}
	}
	return result
}

// Filter 切片过滤
func (e extend) Filter(array interface{}, filter func(i interface{}) bool) []interface{} {
	result := make([]interface{}, 0)
	v := reflect.ValueOf(array)
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if filter(v.Index(i).Interface()) {
				result = append(result, v.Index(i).Interface())
			}
		}
		return result
	default:
		return result
	}
}

// Map 切片类型转换
func (e extend) Map(array interface{}, transform func(i interface{}) interface{}) []interface{} {
	result := make([]interface{}, 0)
	v := reflect.ValueOf(array)
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			result = append(result, transform(v.Index(i).Interface()))
		}
		return result
	default:
		return result
	}
}

// ForEach 切片遍历
func (e extend) ForEach(array interface{}, each func(element interface{})) {
	v := reflect.ValueOf(array)
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			each(v.Index(i).Interface())
		}
	default:
		return
	}
}
