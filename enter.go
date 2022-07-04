package tools

import (
	m "github.com/loebfly/tools/map"
	"github.com/loebfly/tools/math"
	str "github.com/loebfly/tools/string"
)

type Tools struct {
	Math   math.Instance
	String str.Instance
	Map    m.Instance
}

func Get() *Tools {
	return &Tools{}
}
