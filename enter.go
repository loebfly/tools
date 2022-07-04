package tools

import "github.com/loebfly/tools/math"

type Main struct {
	Math math.Instance
}

func Instance() *Main {
	return &Main{}
}
