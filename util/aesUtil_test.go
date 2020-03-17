package util

import (
	"fmt"
	"testing"
)

func TestAesGcm(t *testing.T) {
	key := GetRandomString(32)
	sec := AesGcmEncrypt(key, []byte("123456789"))
	res := AesGcmDecrypt(key, sec)
	fmt.Println(string(res))
}
