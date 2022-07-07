package main

import (
	"fmt"
	"github.com/loebfly/tools"
)

func main() {
	tools.Math.Compare.Min(1, 2)

	tools.String.Verify.IsCNMobile("13888888888")

	tools.Map.StringExt.IsExistKey(map[string]string{"a": "1", "b": "2"}, "a")

	tools.Map.StringExt.GetIntValue(map[string]string{"a": "1", "b": "2"}, "a")

	type T struct {
		A int `json:"a"`
	}

	var t T
	tools.String.Json.StoI("{\"a\":1}", &t)
	fmt.Println(t)

	tools.Crypto.MD5.EncodeString("123456")
}
