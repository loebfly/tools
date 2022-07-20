package mapT

import (
	"reflect"
)

type extend struct{}

// IsExist 判断key是否存在
func (e extend) IsExist(p interface{}, k string) bool {
	v := reflect.ValueOf(p)
	switch v.Kind() {
	case reflect.Map:
		_, ok := v.MapIndex(reflect.ValueOf(k)).Interface().(string)
		return ok
	default:
		return false
	}
}

// IsEmptyV 判断值是否为空
func (e extend) IsEmptyV(p interface{}, k string) bool {
	v := reflect.ValueOf(p)
	switch v.Kind() {
	case reflect.Map:
		val := v.MapIndex(reflect.ValueOf(k)).Interface()
		return val == "" || val == nil
	default:
		return false
	}
}

// GetKeys 获取所有key
func (e extend) GetKeys(p interface{}) []string {
	var keys = make([]string, 0)
	v := reflect.ValueOf(p)
	switch v.Kind() {
	case reflect.Map:
		for _, k := range v.MapKeys() {
			keys = append(keys, k.String())
		}
		return keys
	default:
		return keys
	}
}

// GetValues 获取所有value
func (e extend) GetValues(p interface{}) []interface{} {
	var values = make([]interface{}, 0)
	v := reflect.ValueOf(p)
	switch v.Kind() {
	case reflect.Map:
		for _, k := range v.MapKeys() {
			values = append(values, v.MapIndex(k).Interface())
		}
		return values
	default:
		return values
	}
}

// Filter 字典过滤
func (e extend) Filter(p interface{}, filter func(key string, value interface{}) bool) map[string]interface{} {
	var result = make(map[string]interface{})
	v := reflect.ValueOf(p)
	switch v.Kind() {
	case reflect.Map:
		for _, k := range v.MapKeys() {
			if filter(k.String(), v.MapIndex(k).Interface()) {
				result[k.String()] = v.MapIndex(k).Interface()
			}
		}
		return result
	default:
		return result
	}
}

// MapValue 字典值转换
func (e extend) MapValue(p interface{}, each func(key string, value interface{}) interface{}) map[string]interface{} {
	var result = make(map[string]interface{})
	v := reflect.ValueOf(p)
	switch v.Kind() {
	case reflect.Map:
		for _, k := range v.MapKeys() {
			result[k.String()] = each(k.String(), v.MapIndex(k).Interface())
		}
		return result
	default:
		return result
	}
}

// ForEach 字典遍历
func (e extend) ForEach(p interface{}, each func(key string, value interface{})) {
	v := reflect.ValueOf(p)
	switch v.Kind() {
	case reflect.Map:
		for _, k := range v.MapKeys() {
			each(k.String(), v.MapIndex(k).Interface())
		}
	default:
		return
	}
}
