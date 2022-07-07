package m

import "strconv"

type StringExt struct{}

func (receiver *StringExt) IsExistKey(p map[string]string, k string) bool {
	_, ok := p[k]
	return ok
}

func (receiver *StringExt) isEmpty(p map[string]string, k string) bool {
	if receiver.IsExistKey(p, k) {
		return p[k] == ""
	}
	return true
}

func (receiver *StringExt) GetKeys(p map[string]string) []string {
	var keys = make([]string, 0)
	for k := range p {
		keys = append(keys, k)
	}
	return keys
}

func (receiver *StringExt) GetValues(p map[string]string) []interface{} {
	var values = make([]interface{}, 0)
	for _, v := range p {
		values = append(values, v)
	}
	return values
}

func (receiver *StringExt) GetFloatValue(p map[string]string, k string) float64 {
	if receiver.IsExistKey(p, k) {
		v, e := strconv.ParseFloat(p[k], 64)
		if e == nil {
			return 0
		}
		return v
	}
	return 0
}

func (receiver *StringExt) GetBoolValue(p map[string]string, k string) bool {
	if receiver.IsExistKey(p, k) {
		v, e := strconv.ParseBool(p[k])
		if e == nil {
			return v
		}
		return false
	}
	return false
}

func (receiver *StringExt) GetIntValue(p map[string]string, k string) int64 {
	if receiver.IsExistKey(p, k) {
		v, e := strconv.ParseInt(p[k], 10, 64)
		if e == nil {
			return v
		}
		return 0
	}
	return 0
}

func (receiver *StringExt) GetInt32Value(p map[string]string, k string) int32 {
	if receiver.IsExistKey(p, k) {
		v, e := strconv.ParseInt(p[k], 10, 32)
		if e == nil {
			return int32(v)
		}
		return 0
	}
	return 0
}

func (receiver *StringExt) GetInt16Value(p map[string]string, k string) int16 {
	if receiver.IsExistKey(p, k) {
		v, e := strconv.ParseInt(p[k], 10, 16)
		if e == nil {
			return int16(v)
		}
		return 0
	}
	return 0
}

func (receiver *StringExt) GetInt8Value(p map[string]string, k string) int8 {
	if receiver.IsExistKey(p, k) {
		v, e := strconv.ParseInt(p[k], 10, 8)
		if e == nil {
			return int8(v)
		}
		return 0
	}
	return 0
}

func (receiver *StringExt) GetUintValue(p map[string]string, k string) uint {
	if receiver.IsExistKey(p, k) {
		v, e := strconv.ParseUint(p[k], 10, 64)
		if e == nil {
			return uint(v)
		}
		return 0
	}
	return 0
}

func (receiver *StringExt) GetUint32Value(p map[string]string, k string) uint32 {
	if receiver.IsExistKey(p, k) {
		v, e := strconv.ParseUint(p[k], 10, 32)
		if e == nil {
			return uint32(v)
		}
		return 0
	}
	return 0
}

func (receiver *StringExt) GetUint16Value(p map[string]string, k string) uint16 {
	if receiver.IsExistKey(p, k) {
		v, e := strconv.ParseUint(p[k], 10, 16)
		if e == nil {
			return uint16(v)
		}
		return 0
	}
	return 0
}

func (receiver *StringExt) GetUint8Value(p map[string]string, k string) uint8 {
	if receiver.IsExistKey(p, k) {
		v, e := strconv.ParseUint(p[k], 10, 8)
		if e == nil {
			return uint8(v)
		}
		return 0
	}
	return 0
}
