package uuidT

import (
	"github.com/gofrs/uuid"
	"strings"
)

// GetV1 基于时间和MAC地址生成的UUID, offMinus 是否去除"-"
func (ue Enter) GetV1(offMinus bool) string {
	u, _ := uuid.NewV1()
	if !offMinus {
		return strings.ReplaceAll(u.String(), "-", "")
	}
	return u.String()
}

// GetV4 基于随机数生成的UUID, offMinus 是否去除"-"
func (ue Enter) GetV4(offMinus bool) string {
	u, _ := uuid.NewV4()
	if !offMinus {
		return strings.ReplaceAll(u.String(), "-", "")
	}
	return u.String()
}

// Get == GetV1 基于时间和MAC地址生成的UUID, offMinus 是否去除"-"
func (ue Enter) Get(offMinus bool) string {
	return ue.GetV1(offMinus)
}
