package jsonT

import (
	"encoding/json"
	"strings"
)

// ToJson converts obj to JSON.
func (js Enter) ToJson(obj interface{}) string {
	b, err := json.Marshal(obj)
	if err != nil {
		return "{}"
	} else {
		result := string(b)
		result = strings.Replace(result, "\\u003c", "<", -1)
		result = strings.Replace(result, "\\u003e", ">", -1)
		result = strings.Replace(result, "\\u0026", "&", -1)
		return result
	}
}

// ToObjS converts JSON to obj.
func (js Enter) ToObjS(src string, obj interface{}) bool {
	if src == "" {
		return false
	}
	err := json.Unmarshal([]byte(src), &obj)
	if err != nil {
		return false
	} else {
		return true
	}
}

// ToObjB converts []byte to obj.
func (js Enter) ToObjB(src []byte, obj interface{}) bool {
	if src == nil {
		return false
	}
	jsonStr := js.ToJson(src)
	return js.ToObjS(jsonStr, obj)
}

// ToObjI converts map or slice to JSON.
func (js Enter) ToObjI(src interface{}, obj interface{}) bool {
	if src == nil {
		return false
	}
	jsonStr := js.ToJson(src)
	return js.ToObjS(jsonStr, obj)
}
