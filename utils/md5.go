package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 用于 token 编码
func MD5(str []byte, b ...byte) string {
	hash := md5.New()
	hash.Write(str)
	return hex.EncodeToString(hash.Sum(b)) // md5.Sum 是直接将字节切片做md5运算
}
