package mapT

type Extend struct{}

// IsExistS 判断key是否存在
func (e Extend) IsExistS(p map[string]string, k string) bool {
	_, ok := p[k]
	return ok
}

// IsExistI 判断key是否存在
func (e Extend) IsExistI(p map[string]interface{}, k string) bool {
	_, ok := p[k]
	return ok
}

// IsEmptyS 判断是否为空
func (e Extend) IsEmptyS(p map[string]string, k string) bool {
	if e.IsExistS(p, k) {
		return p[k] == ""
	}
	return true
}

// IsEmptyI 判断是否为空
func (e Extend) IsEmptyI(p map[string]interface{}, k string) bool {
	if e.IsExistI(p, k) {
		return p[k] == ""
	}
	return true
}

// GetKeysS 获取所有key
func (e Extend) GetKeysS(p map[string]string) []string {
	var keys = make([]string, 0)
	for k := range p {
		keys = append(keys, k)
	}
	return keys
}

// GetKeysI 获取所有key
func (e Extend) GetKeysI(p map[string]interface{}) []string {
	var keys = make([]string, 0)
	for k := range p {
		keys = append(keys, k)
	}
	return keys
}

// GetValuesS 获取所有value
func (e Extend) GetValuesS(p map[string]string) []string {
	var values = make([]string, 0)
	for _, v := range p {
		values = append(values, v)
	}
	return values
}

// GetValuesI 获取所有value
func (e Extend) GetValuesI(p map[string]interface{}) []interface{} {
	var values = make([]interface{}, 0)
	for _, v := range p {
		values = append(values, v)
	}
	return values
}
