package mapT

type merge struct{}

// Interfaces 多个字典合并
func (m merge) Interfaces(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, map_ := range maps {
		for key, value := range map_ {
			result[key] = value
		}
	}
	return result
}

// Strings 多个字符串字典合并
func (m merge) Strings(maps ...map[string]string) map[string]string {
	result := make(map[string]string)
	for _, map_ := range maps {
		for key, value := range map_ {
			result[key] = value
		}
	}
	return result
}

// Ints 多个整数字典合并
func (m merge) Ints(maps ...map[string]int) map[string]int {
	result := make(map[string]int)
	for _, map_ := range maps {
		for key, value := range map_ {
			result[key] = value
		}
	}
	return result
}
