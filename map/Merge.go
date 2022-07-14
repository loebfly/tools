package mapT

type Merge struct{}

// Interfaces 多个字典合并
func (m Merge) Interfaces(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, map_ := range maps {
		for key, value := range map_ {
			result[key] = value
		}
	}
	return result
}

// Strings 多个字符串字典合并
func (m Merge) Strings(maps ...map[string]string) map[string]string {
	result := make(map[string]string)
	for _, map_ := range maps {
		for key, value := range map_ {
			result[key] = value
		}
	}
	return result
}
