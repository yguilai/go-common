package hash

import (
	"crypto/md5"
	"fmt"
)

func Md5(p []byte) []byte {
	digest := md5.New()
	digest.Write(p)
	return digest.Sum(nil)
}

func Md5Hex(p []byte) string {
	return fmt.Sprintf("%x", Md5(p))
}
