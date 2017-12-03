package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func Rand(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func MD5(s string) string {
	data := []byte(s)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return strings.ToUpper(md5str)
}
