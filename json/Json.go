package jsonT

import (
	"encoding/json"
	"github.com/tidwall/gjson"
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

// Valid checks if the JSON is valid.
func (js Enter) Valid(src string) bool {
	return gjson.Valid(src)
}

// GetValue searches json for the specified path.
// A path is in dot syntax, such as "name.last" or "age".
// When the value is found it's returned immediately.
//
// A path is a series of keys separated by a dot.
// A key may contain special wildcard characters '*' and '?'.
// To access an array value use the index as the key.
// To get the number of elements in an array or to access a child path, use
// the '#' character.
// The dot and wildcard character can be escaped with '\'.
//
//  {
//    "name": {"first": "Tom", "last": "Anderson"},
//    "age":37,
//    "children": ["Sara","Alex","Jack"],
//    "friends": [
//      {"first": "James", "last": "Murphy"},
//      {"first": "Roger", "last": "Craig"}
//    ]
//  }
//  "name.last"          >> "Anderson"
//  "age"                >> 37
//  "children"           >> ["Sara","Alex","Jack"]
//  "children.#"         >> 3
//  "children.1"         >> "Alex"
//  "child*.2"           >> "Jack"
//  "c?ildren.0"         >> "Sara"
//  "friends.#.first"    >> ["James","Roger"]
//
// This function expects that the json is well-formed, and does not validate.
// Invalid json will not panic, but it may return back unexpected results.
// If you are consuming JSON from an unpredictable source then you may want to
// use the Valid function first.
func (js Enter) GetValue(src string, path string) gjson.Result {
	return gjson.Get(src, path)
}

// GetValues searches json for the multiple paths.
// The return value is a Result array where the number of items
// will be equal to the number of input paths.
func (js Enter) GetValues(src string, path ...string) []gjson.Result {
	return gjson.GetMany(src, path...)
}
