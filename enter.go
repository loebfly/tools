package tools

import (
	"github.com/loebfly/tools/crypto"
	ginext "github.com/loebfly/tools/gin"
	m "github.com/loebfly/tools/map"
	"github.com/loebfly/tools/math"
	str "github.com/loebfly/tools/string"
)

var Math = new(math.Enter)     // 算法工具
var String = new(str.Enter)    // 字符串工具
var Map = new(m.Enter)         // map工具
var Crypto = new(crypto.Enter) // 加密工具
var Gin = new(ginext.Enter)    // gin工具
