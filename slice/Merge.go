package slice

type merge struct{}

// Interfaces 合并泛型数组
func (m merge) Interfaces(list ...[]interface{}) (result []interface{}) {
	for _, value := range list {
		for _, v := range value {
			result = append(result, v)
		}
	}
	return
}

// Strings 合并字符串数组
func (m merge) Strings(list ...[]string) (result []string) {
	for _, value := range list {
		for _, v := range value {
			result = append(result, v)
		}
	}
	return
}

// Ints 合并整数数组
func (m merge) Ints(list ...[]int) (result []int) {
	for _, value := range list {
		for _, v := range value {
			result = append(result, v)
		}
	}
	return
}
