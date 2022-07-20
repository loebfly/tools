package uuidT

import (
	"github.com/gofrs/uuid"
	"strings"
)

// NewV1 基于时间和MAC地址生成的UUID, offMinus 是否去除"-"
func (ue Enter) NewV1(offMinus bool) string {
	u, _ := uuid.NewV1()
	if !offMinus {
		return strings.ReplaceAll(u.String(), "-", "")
	}
	return u.String()
}

// NewV4 基于随机数生成的UUID, offMinus 是否去除"-"
func (ue Enter) NewV4(offMinus bool) string {
	u, _ := uuid.NewV4()
	if !offMinus {
		return strings.ReplaceAll(u.String(), "-", "")
	}
	return u.String()
}

// New == NewV4 获取去除"-"基于随机数生成的UUID
func (ue Enter) New() string {
	return ue.NewV4(true)
}
