package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"os"
	"shira-chan-dev/config"
)

func ReadRSAPKCS1PrivateKey() (*rsa.PrivateKey, error) {
	// 读取文件
	path := config.Config.Rsakey
	context, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// pem解码
	pemBlock, _ := pem.Decode(context)
	// x509解码
	privateKey, err := x509.ParsePKCS8PrivateKey(pemBlock.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey.(*rsa.PrivateKey), err
}

func RSADecrypt(base64data string) (string, error) {
	// http://liuqh.icu/2021/06/20/go/package/17-rsa/
	// data反解base64
	decodeString, err := base64.StdEncoding.DecodeString(base64data)
	if err != nil {
		return "", err
	}
	// 读取密钥
	rsaPrivateKey, err := ReadRSAPKCS1PrivateKey()
	if err != nil {
		return "", err
	}
	// 解密
	decryptPKCS1v15, err := rsa.DecryptPKCS1v15(rand.Reader, rsaPrivateKey, decodeString)
	return string(decryptPKCS1v15), err
}
