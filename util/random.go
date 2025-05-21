package util

import (
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init () {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max - min + 1)
}

func RandomString(n int) string {
	letters := []rune(alphabet)
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return int64(RandomInt(0, 1000))
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[RandomInt(0, n - 1)]
}