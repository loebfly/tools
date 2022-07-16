package stringT

import (
	"encoding/json"
	"strings"
)

type jsonString struct{}

// ToJsonFromObj converts obj to JSON.
func (js jsonString) ToJsonFromObj(obj interface{}) string {
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

// ToObjFromJson converts JSON to obj.
func (js jsonString) ToObjFromJson(src string, obj interface{}) bool {
	err := json.Unmarshal([]byte(src), &obj)
	if err != nil {
		return false
	} else {
		return true
	}
}

// ToObjFromInterface converts map or slice to obj.
func (js jsonString) ToObjFromInterface(src interface{}, obj interface{}) bool {
	jsonStr := js.ToJsonFromObj(src)
	return js.ToObjFromJson(jsonStr, obj)
}
