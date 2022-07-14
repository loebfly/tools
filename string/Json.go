package str

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
		json := string(b)
		json = strings.Replace(json, "\\u003c", "<", -1)
		json = strings.Replace(json, "\\u003e", ">", -1)
		json = strings.Replace(json, "\\u0026", "&", -1)
		return json
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
