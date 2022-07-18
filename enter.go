package tools

import (
	"github.com/loebfly/tools/crypto"
	fileT "github.com/loebfly/tools/file"
	"github.com/loebfly/tools/gin"
	jsonT "github.com/loebfly/tools/json"
	"github.com/loebfly/tools/map"
	"github.com/loebfly/tools/math"
	"github.com/loebfly/tools/string"
)

var Math = new(math.Enter)      // 计算工具
var String = new(stringT.Enter) // 字符串工具
var Map = new(mapT.Enter)       // map工具
var Crypto = new(crypto.Enter)  // 加密工具
var Gin = new(ginT.Enter)       // gin工具
var Json = new(jsonT.Enter)     // json工具
var File = new(fileT.Enter)     // 文件工具
