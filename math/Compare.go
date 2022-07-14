package math

type compareMath struct{}

// MinInt 返回两个整数中的小值
func (cm compareMath) MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MaxInt 返回两个整数中的大值
func (cm compareMath) MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MinFloat64 返回两个浮点数中的小值
func (cm compareMath) MinFloat64(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// MaxFloat64 返回两个浮点数中的大值
func (cm compareMath) MaxFloat64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// Min 返回两个数中的小值
func (cm compareMath) Min(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		switch b.(type) {
		case int:
			return cm.MinInt(a.(int), b.(int))
		case float64:
			return cm.MinFloat64(float64(a.(int)), b.(float64))
		}
	case float64:
		switch b.(type) {
		case int:
			return cm.MinFloat64(a.(float64), float64(b.(int)))
		case float64:
			return cm.MinFloat64(a.(float64), b.(float64))
		}
	}
	return nil
}

// Max 返回两个数中的大值
func (cm compareMath) Max(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		switch b.(type) {
		case int:
			return cm.MaxInt(a.(int), b.(int))
		case float64:
			return cm.MaxFloat64(float64(a.(int)), b.(float64))
		}
	case float64:
		switch b.(type) {
		case int:
			return cm.MaxFloat64(a.(float64), float64(b.(int)))
		case float64:
			return cm.MaxFloat64(a.(float64), b.(float64))
		}
	}
	return nil
}
