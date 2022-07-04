package tools

import (
	"github.com/loebfly/tools/math"
	str "github.com/loebfly/tools/string"
)

type Main struct {
	Math   math.Instance
	String str.Instance
}

func Instance() *Main {
	return &Main{}
}
