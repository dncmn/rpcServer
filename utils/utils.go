package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"strings"
)

// 判断字符串是否为空
func IsStringEmpty(str string) bool {
	strSli := strings.Trim(str, " ")
	return len(strSli) == 0
}

// Encode string by md5.
func EncodeMD5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return hex.EncodeToString(h.Sum(nil))
}
