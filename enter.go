package tools

import (
	"github.com/loebfly/tools/crypto"
	m "github.com/loebfly/tools/map"
	"github.com/loebfly/tools/math"
	str "github.com/loebfly/tools/string"
)

type Tools struct {
	Math   math.Instance
	String str.Instance
	Map    m.Instance
	Crypto crypto.Instance
}

var Router = new(Tools)
