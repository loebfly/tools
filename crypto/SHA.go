package crypto

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)

type shaCrypto struct{}

// SHA1 加密
func (se shaCrypto) SHA1(src string) string {
	h := sha1.New()
	h.Write([]byte(src))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

// SHA256 加密
func (se shaCrypto) SHA256(src string) string {
	sha256Hash := sha256.New()
	sData := []byte(src)
	sha256Hash.Write(sData)
	hashed := sha256Hash.Sum(nil)
	return fmt.Sprintf("%x", hashed)
}
