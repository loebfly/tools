package crypto

import "encoding/base64"

type base64Crypto struct{}

// EncodeByte 对字节数组进行 base64 编码
func (bc base64Crypto) EncodeByte(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

// DecodeByte 对 base64 编码的字节数组进行解码
func (bc base64Crypto) DecodeByte(src []byte) []byte {
	out, err := base64.StdEncoding.DecodeString(string(src))
	if err != nil {
		return nil
	} else {
		return out
	}
}

// EncodeString 对字符串进行 base64 编码
func (bc base64Crypto) EncodeString(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

// DecodeString 对 base64 编码的字符串进行解码
func (bc base64Crypto) DecodeString(src string) string {
	out, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return ""
	} else {
		return string(out)
	}
}
