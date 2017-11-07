package util

import (
	"crypto/sha512"
	"encoding/hex"
	"math/rand"
	"strings"
	"time"
	"crypto/sha256"
)

var seedString = strings.Split("1234567890abcdefghijklmnopqrstuvwxyz", "")

func GetRandomString() string {
	result := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 8; i++ {
		result += seedString[rand.Intn(len(seedString))]
	}

	return result
}

func GetSha512Password(password string, salt string) string {
	hash := sha512.New()
	hash.Write([]byte(password + salt))
	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}

func GetSha256Password(password string, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(password + salt))
	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}
