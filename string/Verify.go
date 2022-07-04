package str

import "regexp"

type Verify struct {
}

func (receiver *Verify) IsCNMobile(s string) bool {
	reg, _ := regexp.Compile(`^1[3456789]\d{9}$`)
	return reg.MatchString(s)
}

func (receiver *Verify) IsEmail(s string) bool {
	reg, _ := regexp.Compile(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`)
	return reg.MatchString(s)
}

func (receiver *Verify) IsIDCard(s string) bool {
	reg, _ := regexp.Compile(`^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`)
	return reg.MatchString(s)
}

func (receiver *Verify) IsIP(s string) bool {
	reg, _ := regexp.Compile(`^((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)$`)
	return reg.MatchString(s)
}

func (receiver *Verify) IsURL(s string) bool {
	reg, _ := regexp.Compile(`^http[s]?://[\w.]+[\w/]$`)
	return reg.MatchString(s)
}

func (receiver *Verify) IsNum(s string) bool {
	reg, _ := regexp.Compile(`^[0-9]+$`)
	return reg.MatchString(s)
}

func (receiver *Verify) IsEnglish(s string) bool {
	reg, _ := regexp.Compile(`^[a-zA-Z]+$`)
	return reg.MatchString(s)
}

func (receiver *Verify) IsChinese(s string) bool {
	reg, _ := regexp.Compile(`^[\u4e00-\u9fa5]+$`)
	return reg.MatchString(s)
}

func (receiver *Verify) IsLowerCase(s string) bool {
	reg, _ := regexp.Compile(`^[a-z]+$`)
	return reg.MatchString(s)
}

func (receiver *Verify) IsUpperCase(s string) bool {
	reg, _ := regexp.Compile(`^[A-Z]+$`)
	return reg.MatchString(s)
}
