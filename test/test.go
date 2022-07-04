package main

import "github.com/loebfly/tools"

func main() {
	compare := tools.Instance().Math.Compare
	compare.Min(1, 2)

	tools.Instance().String.Verify.IsCNMobile("13888888888")
}
