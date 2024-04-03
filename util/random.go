package util

import (
	"backend-master-class/enums"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates random int between min and max
func RandomInt(min int64, max int64) int64 {
	return min + rand.Int63n(max-(min+1))
}

// RandomString generates random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	var randIdx int
	for i := 0; i < n; i++ {
		randIdx = rand.Intn(len(alphabet))
		sb.WriteByte(alphabet[randIdx])
	}
	return sb.String()
}

// RandomOwner generates random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates random currency code
func RandomCurrency() int {
	currency := []int{enums.EGP, enums.EUR, enums.USD}
	currencyCount := int64(len(currency))
	return currency[RandomInt(0, currencyCount)]
}

func RandomEmail() string {
	return fmt.Sprintf("%s.email.com", RandomString(6))
}
