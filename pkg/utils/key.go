package utils

import (
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var letterRunesLen = len(letterRunes)

func RandomString(n int) string {
	r := make([]rune, n)
	for i := range r {
		r[i] = letterRunes[rand.Intn(letterRunesLen)]
	}
	return string(r)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
