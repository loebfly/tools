package slice

type Extend struct{}

// FilterEmptyS 去除字符串切片的空值
func (e Extend) FilterEmptyS(list []string) []string {
	result := make([]string, 0)
	for _, value := range list {
		if value != "" {
			result = append(result, value)
		}
	}
	return result
}

// FilterEmptyI 去除interface切片的空值
func (e Extend) FilterEmptyI(list []interface{}) []interface{} {
	result := make([]interface{}, 0)
	for _, value := range list {
		if str, ok := value.(string); ok && str != "" {
			result = append(result, value)
		} else if value != nil {
			result = append(result, value)
		}
	}
	return result
}

// InS 判断字符串是否在字符串切片中
func (e Extend) InS(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

// InI 判断interface是否在interface切片中
func (e Extend) InI(list []interface{}, value interface{}) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

// UniqueS 去重字符串数组的重复值
func (e Extend) UniqueS(list []string) []string {
	result := make([]string, 0)
	for _, value := range list {
		if !e.InS(result, value) {
			result = append(result, value)
		}
	}
	return result
}

// UniqueI 去重interface数组的重复值
func (e Extend) UniqueI(list []interface{}) []interface{} {
	result := make([]interface{}, 0)
	for _, value := range list {
		if !e.InI(result, value) {
			result = append(result, value)
		}
	}
	return result
}
