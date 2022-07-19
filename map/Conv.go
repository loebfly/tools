package mapT

type conv struct{}

// StringToI 将字符串字典转换为interface字典
func (c conv) StringToI(src map[string]string) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range src {
		result[k] = v
	}
	return result
}

// IntToI 将整数字典转换为interface字典
func (c conv) IntToI(src map[string]int) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range src {
		result[k] = v
	}
	return result
}

// Int8ToI 将int8字典转换为interface字典
func (c conv) Int8ToI(src map[string]int8) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range src {
		result[k] = v
	}
	return result
}

// Int16ToI 将int16字典转换为interface字典
func (c conv) Int16ToI(src map[string]int16) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range src {
		result[k] = v
	}
	return result
}

// Int32ToI 将int32字典转换为interface字典
func (c conv) Int32ToI(src map[string]int32) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range src {
		result[k] = v
	}
	return result
}

// Int64ToI 将int64字典转换为interface字典
func (c conv) Int64ToI(src map[string]int64) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range src {
		result[k] = v
	}
	return result
}

// UintToI 将uint字典转换为interface字典
func (c conv) UintToI(src map[string]uint) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range src {
		result[k] = v
	}
	return result
}

// Uint8ToI 将uint8字典转换为interface字典
func (c conv) Uint8ToI(src map[string]uint8) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range src {
		result[k] = v
	}
	return result
}

// Uint16ToI 将uint16字典转换为interface字典
func (c conv) Uint16ToI(src map[string]uint16) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range src {
		result[k] = v
	}
	return result
}

// Uint32ToI 将uint32字典转换为interface字典
func (c conv) Uint32ToI(src map[string]uint32) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range src {
		result[k] = v
	}
	return result
}

// Uint64ToI 将uint64字典转换为interface字典
func (c conv) Uint64ToI(src map[string]uint64) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range src {
		result[k] = v
	}
	return result
}

// Float32ToI 将float32字典转换为interface字典
func (c conv) Float32ToI(src map[string]float32) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range src {
		result[k] = v
	}
	return result
}

// Float64ToI 将float64字典转换为interface字典
func (c conv) Float64ToI(src map[string]float64) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range src {
		result[k] = v
	}
	return result
}

// BoolToI 将bool字典转换为interface字典
func (c conv) BoolToI(src map[string]bool) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range src {
		result[k] = v
	}
	return result
}

// Complex64ToI 将complex64字典转换为interface字典
func (c conv) Complex64ToI(src map[string]complex64) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range src {
		result[k] = v
	}
	return result
}

// Complex128ToI 将complex128字典转换为interface字典
func (c conv) Complex128ToI(src map[string]complex128) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range src {
		result[k] = v
	}
	return result
}
