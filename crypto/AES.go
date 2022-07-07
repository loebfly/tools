package crypto

import (
	"bytes"
	"crypto/aes"
	"errors"
)

type AES struct{}

func (a *AES) PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

func (a *AES) PKCS5UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	unPadding := int(origData[length-1])
	if (length - unPadding) < 0 {
		return origData, errors.New("解密错误")
	} else {
		return origData[:(length - unPadding)], nil
	}
}

func (a *AES) EncryptForByte(src, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}
	bs := block.BlockSize()
	src = a.PKCS5Padding(src, bs)
	if len(src)%bs != 0 {
		return nil
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return out
}

func (a *AES) DecryptForByte(src, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return nil
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out, err = a.PKCS5UnPadding(out)
	if err != nil {
		return nil
	}
	return out
}

func (a *AES) EncryptString(src, key string) string {
	return string(a.EncryptForByte([]byte(src), []byte(key)))
}

func (a *AES) DecryptString(src, key string) string {
	return string(a.DecryptForByte([]byte(src), []byte(key)))
}

func (a *AES) EncryptECB(src, key string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return ""
	}
	src = string(a.PKCS5Padding([]byte(src), block.BlockSize()))
	if len(src)%block.BlockSize() != 0 {
		return ""
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, []byte(src[:block.BlockSize()]))
		src = src[block.BlockSize():]
		dst = dst[block.BlockSize():]
	}
	return string(out)
}

func (a *AES) DecryptECB(src, key string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return ""
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Decrypt(dst, []byte(src[:block.BlockSize()]))
		src = src[block.BlockSize():]
		dst = dst[block.BlockSize():]
	}
	out, err = a.PKCS5UnPadding(out)
	if err != nil {
		return ""
	}
	return string(out)
}
