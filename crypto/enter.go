package crypto

type Enter struct {
	MD5    md5Crypto
	Base64 base64Crypto
	AES    aesCrypto
	SHA    shaCrypto
}
