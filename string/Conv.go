package stringT

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type convString struct{}

// ToInt 数字字符串转换为int, 当转换失败时有默认值返回第一个默认值否则返回0
func (conv convString) ToInt(src string, defaultInt ...int) int {
	num, err := strconv.Atoi(src)
	if err != nil {
		if len(defaultInt) > 0 {
			return defaultInt[0]
		} else {
			return 0
		}
	}
	return num
}

// ToInt64 数字字符串转换为int64, 当转换失败时有默认值返回第一个默认值否则返回0
func (conv convString) ToInt64(src string, defaultInt64 ...int64) int64 {
	num, err := strconv.ParseInt(src, 10, 64)
	if err != nil {
		if len(defaultInt64) > 0 {
			return defaultInt64[0]
		} else {
			return 0
		}
	}
	return num
}

// ToInt32 数字字符串转换为int32, 当转换失败时有默认值返回第一个默认值否则返回0
func (conv convString) ToInt32(src string, defaultInt32 ...int32) int32 {
	num, err := strconv.ParseInt(src, 10, 32)
	if err != nil {
		if len(defaultInt32) > 0 {
			return defaultInt32[0]
		} else {
			return 0
		}
	}
	return int32(num)
}

// ToFloat64 数字字符串转换为float64, 当转换失败时有默认值返回第一个默认值否则返回0
func (conv convString) ToFloat64(src string, defaultFloat64 ...float64) float64 {
	num, err := strconv.ParseFloat(src, 64)
	if err != nil {
		if len(defaultFloat64) > 0 {
			return defaultFloat64[0]
		} else {
			return 0
		}
	}
	return num
}

// ToFloat32 数字字符串转换为float32, 当转换失败时有默认值返回第一个默认值否则返回0
func (conv convString) ToFloat32(src string, defaultFloat32 ...float32) float32 {
	num, err := strconv.ParseFloat(src, 32)
	if err != nil {
		if len(defaultFloat32) > 0 {
			return defaultFloat32[0]
		} else {
			return 0
		}
	}
	return float32(num)
}

// String 任意类型转换为字符串
func (conv convString) String(iFace interface{}) string {
	switch val := iFace.(type) {
	case []byte:
		return string(val)
	case string:
		return val
	}
	v := reflect.ValueOf(iFace)
	switch v.Kind() {
	case reflect.Invalid:
		return ""
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'f', -1, 32)
	case reflect.Ptr, reflect.Struct, reflect.Map, reflect.Array:
		b, err := json.Marshal(v.Interface())
		if err != nil {
			return ""
		}
		return string(b)
	}
	return fmt.Sprintf("%v", iFace)
}
