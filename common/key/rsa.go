package key

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

// 生成RSA密钥对
func GenerateRSAKeys() (privatePEM, publicPEM string, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", "", err
	}

	// 私钥PEM编码
	privASN1 := x509.MarshalPKCS1PrivateKey(privateKey)
	privateBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privASN1,
	}
	privatePEM = string(pem.EncodeToMemory(privateBlock))

	// 公钥PEM编码
	pubASN1, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", "", err
	}
	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubASN1,
	}
	publicPEM = string(pem.EncodeToMemory(publicBlock))
	return
}

// 生成AES密钥
func GenerateAESKey() (string, error) {
	key := make([]byte, 32) // 256位
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(key), nil
}
