package string

import "regexp"

type Verify struct {
}

func (receiver *Verify) isCNMobile(s string) bool {
	reg, _ := regexp.Compile(`^1[3456789]\d{9}$`)
	return reg.MatchString(s)
}

func (receiver *Verify) isEmail(s string) bool {
	reg, _ := regexp.Compile(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`)
	return reg.MatchString(s)
}

func (receiver *Verify) isIDCard(s string) bool {
	reg, _ := regexp.Compile(`^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`)
	return reg.MatchString(s)
}

func (receiver *Verify) isIP(s string) bool {
	reg, _ := regexp.Compile(`^((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)$`)
	return reg.MatchString(s)
}

func (receiver *Verify) isURL(s string) bool {
	reg, _ := regexp.Compile(`^http[s]?://[\w.]+[\w/]$`)
	return reg.MatchString(s)
}

func (receiver *Verify) isNum(s string) bool {
	reg, _ := regexp.Compile(`^[0-9]+$`)
	return reg.MatchString(s)
}

func (receiver *Verify) isEnglish(s string) bool {
	reg, _ := regexp.Compile(`^[a-zA-Z]+$`)
	return reg.MatchString(s)
}

func (receiver *Verify) isChinese(s string) bool {
	reg, _ := regexp.Compile(`^[\u4e00-\u9fa5]+$`)
	return reg.MatchString(s)
}

func (receiver *Verify) isLowerCase(s string) bool {
	reg, _ := regexp.Compile(`^[a-z]+$`)
	return reg.MatchString(s)
}

func (receiver *Verify) isUpperCase(s string) bool {
	reg, _ := regexp.Compile(`^[A-Z]+$`)
	return reg.MatchString(s)
}