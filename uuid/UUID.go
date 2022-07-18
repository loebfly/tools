package uuidT

import (
	"github.com/gofrs/uuid"
	"strings"
)

// Get 获取uuid
func (ue Enter) Get() string {
	u, _ := uuid.NewV4()
	return u.String()
}

// GetNoMinus 获取没有-的uuid
func (ue Enter) GetNoMinus() string {
	return strings.ReplaceAll(ue.Get(), "-", "")
}

// GetNoMinusWithLen 获取没有-的uuid，长度为len
func (ue Enter) GetNoMinusWithLen(l int) string {
	u := ue.GetNoMinus()
	if len(u) > l {
		u = u[:l]
	}
	return u
}
