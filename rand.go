package helper

import (
	"math/rand"
	"time"
)

// RandString 生成随机字符串
func RandString(len int) string {
	bytes := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	seed := "_0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < len; i++ {
		bytes = append(bytes, seed[r.Intn(63)])
	}
	return string(bytes)
}
