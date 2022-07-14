package crypto

import (
	"bytes"
	"crypto/aes"
	"errors"
)

type aesCrypto struct{}

// PKCS5Padding 填充
func (ac aesCrypto) PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

// PKCS5UnPadding 去除填充
func (ac aesCrypto) PKCS5UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	unPadding := int(origData[length-1])
	if (length - unPadding) < 0 {
		return origData, errors.New("解密错误")
	} else {
		return origData[:(length - unPadding)], nil
	}
}

// EncryptForByte 给字节数组加密
func (ac aesCrypto) EncryptForByte(src, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}
	bs := block.BlockSize()
	src = ac.PKCS5Padding(src, bs)
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

// DecryptForByte 给字节数组解密
func (ac aesCrypto) DecryptForByte(src, key []byte) []byte {
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
	out, err = ac.PKCS5UnPadding(out)
	if err != nil {
		return nil
	}
	return out
}

// EncryptString 给字符串加密
func (ac aesCrypto) EncryptString(src, key string) string {
	return string(ac.EncryptForByte([]byte(src), []byte(key)))
}

// DecryptString 给字符串解密
func (ac aesCrypto) DecryptString(src, key string) string {
	return string(ac.DecryptForByte([]byte(src), []byte(key)))
}

// EncryptECB 加密
func (ac aesCrypto) EncryptECB(src, key string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return ""
	}
	src = string(ac.PKCS5Padding([]byte(src), block.BlockSize()))
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

// DecryptECB 解密
func (ac aesCrypto) DecryptECB(src, key string) string {
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
	out, err = ac.PKCS5UnPadding(out)
	if err != nil {
		return ""
	}
	return string(out)
}
