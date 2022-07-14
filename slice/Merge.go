package slice

type Merge struct{}

// Interfaces 合并泛型数组
func (m Merge) Interfaces(list ...[]interface{}) (result []interface{}) {
	for _, value := range list {
		for _, v := range value {
			result = append(result, v)
		}
	}
	return
}

// Strings 合并字符串数组
func (m Merge) Strings(list ...[]string) (result []string) {
	for _, value := range list {
		for _, v := range value {
			result = append(result, v)
		}
	}
	return
}
