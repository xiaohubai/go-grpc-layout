package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/sha3"
)

var (
	letters    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // 52
	symbols    = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"                   // 32
	digits     = "0123456789"                                           // 10
	characters = letters + digits + symbols                             // 94
)

// Md5 加密
func Md5(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}

// Sha3 sha3-256加密
func Sha3(str []byte, b ...byte) string {
	h := sha3.New256()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}

// RandString 生成随机字符串 值可能为：
// abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ
// !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~
// 0123456789
func RandString(n int) string {
	b := []byte(characters)
	res := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		res = append(res, b[r.Intn(len(b))])
	}
	return string(res)
}

// RandInt 生成随机数
func RandInt(n int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(n)
}

// 生成uuid
func UUID() string {
	return uuid.NewString()
}
