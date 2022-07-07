package main

import "github.com/loebfly/tools"

func main() {
	tools.Router.Math.Compare.Min(1, 2)

	tools.Router.String.Verify.IsCNMobile("13888888888")

	tools.Router.Map.StringExt.IsExistKey(map[string]string{"a": "1", "b": "2"}, "a")

	tools.Router.Map.StringExt.GetIntValue(map[string]string{"a": "1", "b": "2"}, "a")
}
