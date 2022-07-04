package math

type Compare struct{}

func (receiver *Compare) MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (receiver *Compare) MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (receiver *Compare) MinFloat64(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func (receiver *Compare) MaxFloat64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func (receiver *Compare) Min(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		switch b.(type) {
		case int:
			return receiver.MinInt(a.(int), b.(int))
		case float64:
			return receiver.MinFloat64(float64(a.(int)), b.(float64))
		}
	case float64:
		switch b.(type) {
		case int:
			return receiver.MinFloat64(a.(float64), float64(b.(int)))
		case float64:
			return receiver.MinFloat64(a.(float64), b.(float64))
		}
	}
	return nil
}

func (receiver *Compare) Max(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		switch b.(type) {
		case int:
			return receiver.MaxInt(a.(int), b.(int))
		case float64:
			return receiver.MaxFloat64(float64(a.(int)), b.(float64))
		}
	case float64:
		switch b.(type) {
		case int:
			return receiver.MaxFloat64(a.(float64), float64(b.(int)))
		case float64:
			return receiver.MaxFloat64(a.(float64), b.(float64))
		}
	}
	return nil
}
