package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tmpStr := h.Sum(nil)
	return hex.EncodeToString(tmpStr)
}

func MakePassword(pwd, salt string) string {
	return Md5Encode(pwd + salt)
}

func ValidPassword(pwd, salt string, password string) bool {
	return Md5Encode(pwd+salt) == password
}
