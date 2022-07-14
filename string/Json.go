package stringT

import (
	"encoding/json"
	"strings"
)

type jsonString struct{}

// ItoS converts interface to string.
func (js jsonString) ItoS(src interface{}) string {
	b, err := json.Marshal(src)
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

// StoI converts string to interface.
func (js jsonString) StoI(src string, receiver interface{}) bool {
	err := json.Unmarshal([]byte(src), &receiver)
	if err != nil {
		return false
	} else {
		return true
	}
}
