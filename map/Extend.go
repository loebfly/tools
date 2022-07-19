package mapT

import (
	"encoding/json"
	"reflect"
	"strings"
)

type extend struct{}

// IsExistS 判断key是否存在
func (e extend) IsExistS(p map[string]string, k string) bool {
	_, ok := p[k]
	return ok
}

// IsExistI 判断key是否存在
func (e extend) IsExistI(p map[string]interface{}, k string) bool {
	_, ok := p[k]
	return ok
}

// IsEmptyS 判断是否为空
func (e extend) IsEmptyS(p map[string]string, k string) bool {
	if e.IsExistS(p, k) {
		return p[k] == ""
	}
	return true
}

// IsEmptyI 判断是否为空
func (e extend) IsEmptyI(p map[string]interface{}, k string) bool {
	if e.IsExistI(p, k) {
		return p[k] == ""
	}
	return true
}

// GetKeysS 获取所有key
func (e extend) GetKeysS(p map[string]string) []string {
	var keys = make([]string, 0)
	for k := range p {
		keys = append(keys, k)
	}
	return keys
}

// GetKeysI 获取所有key
func (e extend) GetKeysI(p map[string]interface{}) []string {
	var keys = make([]string, 0)
	for k := range p {
		keys = append(keys, k)
	}
	return keys
}

// GetValuesS 获取所有value
func (e extend) GetValuesS(p map[string]string) []string {
	var values = make([]string, 0)
	for _, v := range p {
		values = append(values, v)
	}
	return values
}

// GetValuesI 获取所有value
func (e extend) GetValuesI(p map[string]interface{}) []interface{} {
	var values = make([]interface{}, 0)
	for _, v := range p {
		values = append(values, v)
	}
	return values
}

// ToJsonFromMapI 字典转Json
func (e extend) ToJsonFromMapI(p map[string]interface{}) string {
	jsonStr, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(jsonStr)
}

// ToJsonFromMapS 字典转Json
func (e extend) ToJsonFromMapS(p map[string]string) string {
	jsonStr, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(jsonStr)
}

// ToMapIFromObj 结构体转字典
func (e extend) ToMapIFromObj(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[strings.ToLower(obj1.Field(i).Name)] = obj2.Field(i).Interface()
	}
	return data
}
