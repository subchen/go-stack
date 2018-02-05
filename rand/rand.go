package rand

import (
	"math/rand"
	"time"
)

var numerics      = []rune("0123456789")
var alphas        = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var alphanumerics = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomInt(min, max int) int {
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

func RandomNumeric(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, length)
	for i := range b {
		b[i] = numerics[rand.Intn(10)]
	}
	return string(b)
}

func RandomAlpha(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, length)
	for i := range b {
		b[i] = alphas[rand.Intn(52)]
	}
	return string(b)
}

func RandomAlphaNumeric(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, length)
	for i := range b {
		b[i] = alphanumerics[rand.Intn(62)]
	}
	return string(b)
}
