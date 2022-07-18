package jsonT

import (
	"encoding/json"
	"strings"
)

// ToJsonFromObj converts obj to JSON.
func (js Enter) ToJsonFromObj(obj interface{}) string {
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
func (js Enter) ToObjFromJson(src string, obj interface{}) bool {
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

// ToObjFromSets converts map or slice to obj.
func (js Enter) ToObjFromSets(src interface{}, obj interface{}) bool {
	if src == nil {
		return false
	}
	jsonStr := js.ToJsonFromObj(src)
	if jsonStr == "" {
		return false
	}

	return js.ToObjFromJson(jsonStr, obj)
}
