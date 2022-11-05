package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-(min+1))
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generate a random owner as a random string of lnegth 6 for testing pursposes
func RandomOwner() string {
	return RandomString(6)
}

// RandomBalance generate a random balance amount
func RandomBalance() int64 {
	return RandomInt(0, 3000)
}

// RandomCurrencies generates a random currency code
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "JPY", "AUD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
