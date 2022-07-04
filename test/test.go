package main

import "github.com/loebfly/tools"

func main() {
	tools.Get().Math.Compare.Min(1, 2)

	tools.Get().String.Verify.IsCNMobile("13888888888")

	tools.Get().Map.StringExt.IsExistKey(map[string]string{"a": "1", "b": "2"}, "a")

	tools.Get().Map.StringExt.GetIntValue(map[string]string{"a": "1", "b": "2"}, "a")
}
