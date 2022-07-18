package crypto

type Enter struct {
	MD5    md5Crypto    // MD5
	Base64 base64Crypto // Base64
	AES    aesCrypto    // AES
	SHA    shaCrypto    // SHA
	RSA    rsaCrypto    // RSA
}
