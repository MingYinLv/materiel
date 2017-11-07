package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"
)

var privateKey, publicKey []byte

func init() {
	pri, err := ioutil.ReadFile("private.pem")
	if err != nil {
		panic(err)
	}
	privateKey = pri

	pub, err := ioutil.ReadFile("public.pem")
	if err != nil {
		panic(err)
	}
	publicKey = pub
}

func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

func GenRsaKey(bits int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "private",
		Bytes: derStream,
	}
	file, err := os.Create("private.pem")
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
		Type:  "public",
		Bytes: derPkix,
	}
	file, err = os.Create("public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}

func RsaEncode(data interface{}) string {
	js, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	tokenByte, err := RsaEncrypt([]byte(string(js)))
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(tokenByte)
}

func RsaDecode(data string) string {
	tokenByte, err := hex.DecodeString(data)
	if err != nil {
		panic(err)
	}
	tokenByte1, err := RsaDecrypt(tokenByte)
	if err != nil {
		panic(err)
	}
	return string(tokenByte1)
}
