package main

import (
	"materiel/src/util"
	"fmt"
	"encoding/hex"
)

func main() {
	aesEnc := util.AesEncrypt{}
	arrEncrypt, err := aesEnc.Encrypt("1")
	if err != nil {
		fmt.Println(arrEncrypt)
		return
	}
	fmt.Println(hex.EncodeToString(arrEncrypt))
	strMsg, err := aesEnc.Decrypt(arrEncrypt)
	if err != nil {
		fmt.Println(arrEncrypt)
		return
	}
	fmt.Println(strMsg)
}
