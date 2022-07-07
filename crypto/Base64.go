package crypto

import "encoding/base64"

type Base64 struct{}

func (b *Base64) EncodeByte(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func (b *Base64) DecodeByte(src []byte) []byte {
	out, err := base64.StdEncoding.DecodeString(string(src))
	if err != nil {
		return nil
	} else {
		return out
	}
}

func (b *Base64) EncodeString(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

func (b *Base64) DecodeString(src string) string {
	out, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return ""
	} else {
		return string(out)
	}
}
