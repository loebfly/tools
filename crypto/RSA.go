package crypto

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"runtime"
)

type rsaCrypto struct{}

type RsaKey struct {
	PrivateKey string
	PublicKey  string
}

// GenerateKeyForHex 以十六进制方式生成密钥
func (rc rsaCrypto) GenerateKeyForHex(bits int) (RsaKey, error) {
	if bits != 1024 && bits != 2048 {
		return RsaKey{}, fmt.Errorf("bits must be 1024 or 2048")
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return RsaKey{}, err
	}
	return RsaKey{
		PrivateKey: hex.EncodeToString(x509.MarshalPKCS1PrivateKey(privateKey)),
		PublicKey:  hex.EncodeToString(x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)),
	}, nil
}

// GenerateKeyForBase64 以base64方式生成密钥
func (rc rsaCrypto) GenerateKeyForBase64(bits int) (RsaKey, error) {
	if bits != 1024 && bits != 2048 {
		return RsaKey{}, fmt.Errorf("bits must be 1024 or 2048")
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return RsaKey{}, err
	}
	return RsaKey{
		PrivateKey: base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PrivateKey(privateKey)),
		PublicKey:  base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)),
	}, nil
}

// SignForBytes 以bytes方式签名
func (rc rsaCrypto) SignForBytes(data, priKey []byte) ([]byte, error) {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				_ = fmt.Errorf("runtime err=%v,Check that the key or text is correct", err)
			default:
				_ = fmt.Errorf("error=%v,check the cipherText ", err)
			}
		}
	}()
	privateKey, err := x509.ParsePKCS1PrivateKey(priKey)
	hashed := shaCrypto{}.SHA256(string(data))
	sign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, []byte(hashed))
	if err != nil {
		return nil, err
	}
	return sign, nil
}

// VerifyForBytes 以bytes方式验证签名
func (rc rsaCrypto) VerifyForBytes(data, sign, pubKey []byte) error {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				_ = fmt.Errorf("runtime err=%v,Check that the key or text is correct", err)
			default:
				_ = fmt.Errorf("error=%v,check the cipherText ", err)
			}
		}
	}()
	publicKey, err := x509.ParsePKCS1PublicKey(pubKey)
	hashed := shaCrypto{}.SHA256(string(data))
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, []byte(hashed), sign)
	if err != nil {
		return err
	}
	return nil
}

// SignForBase64 以base64方式签名
func (rc rsaCrypto) SignForBase64(data []byte, base64PriKey string) (string, error) {
	priBytes, err := base64.StdEncoding.DecodeString(base64PriKey)
	if err != nil {
		return "", err
	}
	sign, err := rc.SignForBytes(data, priBytes)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sign), nil
}

// VerifyForBase64 以base64方式验证签名
func (rc rsaCrypto) VerifyForBase64(data []byte, base64Sign, base64PubKey string) error {
	sign, err := base64.StdEncoding.DecodeString(base64Sign)
	if err != nil {
		return err
	}
	pubBytes, err := base64.StdEncoding.DecodeString(base64PubKey)
	if err != nil {
		return err
	}
	return rc.VerifyForBytes(data, sign, pubBytes)
}

// SignForHex 以十六进制方式签名
func (rc rsaCrypto) SignForHex(data []byte, hexPriKey string) (string, error) {
	priBytes, err := hex.DecodeString(hexPriKey)
	if err != nil {
		return "", err
	}
	sign, err := rc.SignForBytes(data, priBytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(sign), nil
}

// VerifyForHex 以十六进制方式验证签名
func (rc rsaCrypto) VerifyForHex(data []byte, hexSign, hexPubKey string) error {
	sign, err := hex.DecodeString(hexSign)
	if err != nil {
		return err
	}
	pubBytes, err := hex.DecodeString(hexPubKey)
	if err != nil {
		return err
	}
	return rc.VerifyForBytes(data, sign, pubBytes)
}

// EncryptForByte 以bytes方式加密
func (rc rsaCrypto) EncryptForByte(plainText, publicKey []byte) (cipherText []byte, err error) {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				_ = fmt.Errorf("runtime err=%v,Check that the key or text is correct", err)
			default:
				_ = fmt.Errorf("error=%v,check the cipherText ", err)
			}
		}
	}()
	pub, err := x509.ParsePKCS1PublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	pubSize, plainTextSize := pub.Size(), len(plainText)
	// EncryptPKCS1v15 encrypts the given message with RSA and the padding
	// scheme from PKCS #1 v1.5.  The message must be no longer than the
	// length of the public modulus minus 11 bytes.
	//
	// The rand parameter is used as a source of entropy to ensure that
	// encrypting the same message twice doesn't result in the same
	// ciphertext.
	//
	// WARNING: use of this function to encrypt plaintexts other than
	// session keys is dangerous. Use RSA OAEP in new protocols.
	offSet, once := 0, pubSize-11
	buffer := bytes.Buffer{}
	for offSet < plainTextSize {
		endIndex := offSet + once
		if endIndex > plainTextSize {
			endIndex = plainTextSize
		}
		bytesOnce, err := rsa.EncryptPKCS1v15(rand.Reader, pub, plainText[offSet:endIndex])
		if err != nil {
			return nil, err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	cipherText = buffer.Bytes()
	return cipherText, nil
}

// DecryptForBytes 以bytes方式解密
func (rc rsaCrypto) DecryptForBytes(cipherText, privateKey []byte) (plainText []byte, err error) {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				_ = fmt.Errorf("runtime err=%v,Check that the key or text is correct", err)
			default:
				_ = fmt.Errorf("error=%v,check the cipherText ", err)
			}
		}
	}()
	pri, err := x509.ParsePKCS1PrivateKey(privateKey)
	if err != nil {
		return []byte{}, err
	}
	priSize, cipherTextSize := pri.Size(), len(cipherText)
	var offSet = 0
	var buffer = bytes.Buffer{}
	for offSet < cipherTextSize {
		endIndex := offSet + priSize
		if endIndex > cipherTextSize {
			endIndex = cipherTextSize
		}
		bytesOnce, err := rsa.DecryptPKCS1v15(rand.Reader, pri, cipherText[offSet:endIndex])
		if err != nil {
			return nil, err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	plainText = buffer.Bytes()
	return plainText, nil
}

// EncryptForBase64 以base64方式加密
func (rc rsaCrypto) EncryptForBase64(plainText []byte, base64PubKey string) (string, error) {
	pubBytes, err := base64.StdEncoding.DecodeString(base64PubKey)
	if err != nil {
		return "", err
	}
	cipherText, err := rc.EncryptForByte(plainText, pubBytes)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// DecryptForBase64 以base64方式解密
func (rc rsaCrypto) DecryptForBase64(cipherText string, base64PriKey string) ([]byte, error) {
	priBytes, err := base64.StdEncoding.DecodeString(base64PriKey)
	if err != nil {
		return nil, err
	}
	cipherTextBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return nil, err
	}
	plainText, err := rc.DecryptForBytes(cipherTextBytes, priBytes)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}

// EncryptForHex 以十六进制方式加密
func (rc rsaCrypto) EncryptForHex(plainText []byte, hexPubKey string) (string, error) {
	pubBytes, err := hex.DecodeString(hexPubKey)
	if err != nil {
		return "", err
	}
	cipherText, err := rc.EncryptForByte(plainText, pubBytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(cipherText), nil
}

// DecryptForHex 以十六进制方式解密
func (rc rsaCrypto) DecryptForHex(cipherText string, hexPriKey string) ([]byte, error) {
	priBytes, err := hex.DecodeString(hexPriKey)
	if err != nil {
		return nil, err
	}
	cipherTextBytes, err := hex.DecodeString(cipherText)
	if err != nil {
		return nil, err
	}
	plainText, err := rc.DecryptForBytes(cipherTextBytes, priBytes)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}
