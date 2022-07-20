package slice

import "reflect"

type extend struct{}

// IsExist 判断切片是否存在某个值
func (e extend) IsExist(array interface{}, value interface{}) bool {
	v := reflect.ValueOf(array)
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if v.Index(i).Interface() == value {
				return true
			}
		}
		return false
	default:
		return false
	}
}

// Unique 去重切片
func (e extend) Unique(array interface{}) []interface{} {
	result := make([]interface{}, 0)
	v := reflect.ValueOf(array)
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if !e.IsExist(result, v.Index(i).Interface()) {
				result = append(result, v.Index(i).Interface())
			}
		}
		return result
	default:
		return result
	}
}

// Reverse 反转切片
func (e extend) Reverse(array interface{}) []interface{} {
	result := make([]interface{}, 0)
	v := reflect.ValueOf(array)
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		for i := v.Len() - 1; i >= 0; i-- {
			result = append(result, v.Index(i).Interface())
		}
		return result
	default:
		return result
	}
}

// Remove 删除切片中的某个值
func (e extend) Remove(array interface{}, value interface{}) []interface{} {
	result := make([]interface{}, 0)
	v := reflect.ValueOf(array)
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if v.Index(i).Interface() != value {
				result = append(result, v.Index(i).Interface())
			}
		}
	default:
		return result
	}
	return result
}

// Filter 切片过滤
func (e extend) Filter(array interface{}, filter func(element interface{}) bool) []interface{} {
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
func (e extend) Map(array interface{}, transform func(element interface{}) interface{}) []interface{} {
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
