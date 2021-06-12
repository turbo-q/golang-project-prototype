package helper

import (
	"crypto/md5"
	"fmt"
)

func MD5Sum(str string) string {
	md5Str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5Str
}
