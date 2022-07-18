package tools

import (
	"github.com/loebfly/tools/crypto"
	fileT "github.com/loebfly/tools/file"
	ginT "github.com/loebfly/tools/gin"
	ipT "github.com/loebfly/tools/ip"
	jsonT "github.com/loebfly/tools/json"
	mapT "github.com/loebfly/tools/map"
	"github.com/loebfly/tools/math"
	sqlT "github.com/loebfly/tools/sql"
	stringT "github.com/loebfly/tools/string"
)

var (
	Math   = new(math.Enter)    // 计算工具
	String = new(stringT.Enter) // 字符串工具
	Map    = new(mapT.Enter)    // map工具
	Crypto = new(crypto.Enter)  // 加密工具
	Gin    = new(ginT.Enter)    // gin工具
	Json   = new(jsonT.Enter)   // json工具
	File   = new(fileT.Enter)   // 文件工具
	SQL    = new(sqlT.Enter)    // sql工具
	IP     = new(ipT.Enter)     // ip工具
)
