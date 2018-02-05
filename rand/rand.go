package rand

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var digits  = []rune("0123456789")

func RandInt(min, max int) int {
	if min >= max {
		return min
	}
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(max + 1)
	if n < min {
		return Rand(min, max)
	}
	return n
}

func RandString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(52)]
	}
	return string(b)
}

func RandDigits(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = digits[rand.Intn(10)]
	}
	return string(b)
}
