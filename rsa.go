package helper

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path"
)

// GenRSA 生成RSA公钥私钥证书
// bits 秘钥长度
// outputPath 证书保存路径
func GenRSA(bits int, outputPath string) error {
	if outputPath == "" {
		outputPath = "./"
	}
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		err := os.Mkdir(outputPath, 0o777)
		if err != nil {
			return err
		}
	}
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create(path.Join(outputPath, "private.pem"))
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create(path.Join(outputPath, "public.pem"))
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}
