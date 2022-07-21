package stringT

import (
	"regexp"
	"unicode"
)

// IsCNMobile 判断是否是中国大陆手机号
func (e Enter) IsCNMobile(s string) bool {
	reg, _ := regexp.Compile(`^1[3456789]\d{9}$`)
	return reg.MatchString(s)
}

// IsEmail 判断是否是邮箱
func (e Enter) IsEmail(s string) bool {
	reg, _ := regexp.Compile(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`)
	return reg.MatchString(s)
}

// IsIDCard 判断是否是身份证号
func (e Enter) IsIDCard(s string) bool {
	reg, _ := regexp.Compile(`^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`)
	return reg.MatchString(s)
}

// IsURL 判断是否是URL
func (e Enter) IsURL(s string) bool {
	reg, _ := regexp.Compile(`^http[s]?://[\w.]+[\w/]$`)
	return reg.MatchString(s)
}

// IsNumber 判断是否是数字
func (e Enter) IsNumber(s string) bool {
	reg, _ := regexp.Compile(`^[0-9]+$`)
	return reg.MatchString(s)
}

// IsEnglish 判断是否是英文
func (e Enter) IsEnglish(s string) bool {
	reg, _ := regexp.Compile(`^[a-zA-Z]+$`)
	return reg.MatchString(s)
}

// IsChinese 判断是否是中文
func (e Enter) IsChinese(s string) bool {
	//reg, _ := regexp.Compile(`^[\u4e00-\u9fa5]+$`)
	//return reg.MatchString(s)
	for _, r := range s {
		if !unicode.Is(unicode.Han, r) {
			return false
		}
	}
	return true
}

// IsLowerCase 判断是否是小写
func (e Enter) IsLowerCase(s string) bool {
	reg, _ := regexp.Compile(`^[a-z]+$`)
	return reg.MatchString(s)
}

// IsUpperCase 判断是否是大写
func (e Enter) IsUpperCase(s string) bool {
	reg, _ := regexp.Compile(`^[A-Z]+$`)
	return reg.MatchString(s)
}
