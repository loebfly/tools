package Dictionary

import "strconv"

type StringExt struct{}

// IsExistKey 判断key是否存在
func (se StringExt) IsExistKey(p map[string]string, k string) bool {
	_, ok := p[k]
	return ok
}

// IsEmpty 判断是否为空
func (se StringExt) IsEmpty(p map[string]string, k string) bool {
	if se.IsExistKey(p, k) {
		return p[k] == ""
	}
	return true
}

// GetKeys 获取所有key
func (se StringExt) GetKeys(p map[string]string) []string {
	var keys = make([]string, 0)
	for k := range p {
		keys = append(keys, k)
	}
	return keys
}

// GetValues 获取所有value
func (se StringExt) GetValues(p map[string]string) []interface{} {
	var values = make([]interface{}, 0)
	for _, v := range p {
		values = append(values, v)
	}
	return values
}

// GetFloatValue 获取某个Key的float64值
func (se StringExt) GetFloatValue(p map[string]string, k string) float64 {
	if se.IsExistKey(p, k) {
		v, e := strconv.ParseFloat(p[k], 64)
		if e == nil {
			return 0
		}
		return v
	}
	return 0
}

// GetBoolValue 获取某个Key的bool值
func (se StringExt) GetBoolValue(p map[string]string, k string) bool {
	if se.IsExistKey(p, k) {
		v, e := strconv.ParseBool(p[k])
		if e == nil {
			return v
		}
		return false
	}
	return false
}

// GetIntValue 获取某个Key的int值
func (se StringExt) GetIntValue(p map[string]string, k string) int64 {
	if se.IsExistKey(p, k) {
		v, e := strconv.ParseInt(p[k], 10, 64)
		if e == nil {
			return v
		}
		return 0
	}
	return 0
}

// GetInt32Value 获取某个Key的int32值
func (se StringExt) GetInt32Value(p map[string]string, k string) int32 {
	if se.IsExistKey(p, k) {
		v, e := strconv.ParseInt(p[k], 10, 32)
		if e == nil {
			return int32(v)
		}
		return 0
	}
	return 0
}

// GetInt16Value 获取某个Key的int16值
func (se StringExt) GetInt16Value(p map[string]string, k string) int16 {
	if se.IsExistKey(p, k) {
		v, e := strconv.ParseInt(p[k], 10, 16)
		if e == nil {
			return int16(v)
		}
		return 0
	}
	return 0
}

// GetInt8Value 获取某个Key的int8值
func (se StringExt) GetInt8Value(p map[string]string, k string) int8 {
	if se.IsExistKey(p, k) {
		v, e := strconv.ParseInt(p[k], 10, 8)
		if e == nil {
			return int8(v)
		}
		return 0
	}
	return 0
}

// GetUintValue 获取某个Key的uint值
func (se StringExt) GetUintValue(p map[string]string, k string) uint {
	if se.IsExistKey(p, k) {
		v, e := strconv.ParseUint(p[k], 10, 64)
		if e == nil {
			return uint(v)
		}
		return 0
	}
	return 0
}

// GetUint32Value 获取某个Key的uint32值
func (se StringExt) GetUint32Value(p map[string]string, k string) uint32 {
	if se.IsExistKey(p, k) {
		v, e := strconv.ParseUint(p[k], 10, 32)
		if e == nil {
			return uint32(v)
		}
		return 0
	}
	return 0
}

// GetUint16Value 获取某个Key的uint16值
func (se StringExt) GetUint16Value(p map[string]string, k string) uint16 {
	if se.IsExistKey(p, k) {
		v, e := strconv.ParseUint(p[k], 10, 16)
		if e == nil {
			return uint16(v)
		}
		return 0
	}
	return 0
}

// GetUint8Value 获取某个Key的uint8值
func (se StringExt) GetUint8Value(p map[string]string, k string) uint8 {
	if se.IsExistKey(p, k) {
		v, e := strconv.ParseUint(p[k], 10, 8)
		if e == nil {
			return uint8(v)
		}
		return 0
	}
	return 0
}
