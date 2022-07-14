package math

type conditionMath struct{}

// IfThen 返回条件为真时的结果
func (cm conditionMath) IfThen(condition bool, then interface{}) interface{} {
	if condition {
		return then
	}
	return nil
}

// IfThenElse 返回条件为真时的结果，否则返回另一个结果
func (cm conditionMath) IfThenElse(condition bool, t1, t2 interface{}) interface{} {
	if condition {
		return t1
	}
	return t2
}

// DefaultIfNil 返回值为空时的默认值，否则返回原值
func (cm conditionMath) DefaultIfNil(value interface{}, defaultValue interface{}) interface{} {
	if value != nil {
		return value
	}
	return defaultValue
}

// FirstNonNil 返回第一个非空值
func (cm conditionMath) FirstNonNil(values ...interface{}) interface{} {
	for _, value := range values {
		if value != nil {
			return value
		}
	}
	return nil
}
