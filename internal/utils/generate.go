package utils

import (
	"crypto/rand"
	"math/big"
)

func GenerateAccountNumber(length int) string {
	validChars := []rune("0123456789")
	accountNumber := make([]rune, length)

	for i := range accountNumber {
		// Generate a random integer in the range 0-9
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(validChars))))
		if err != nil {
			panic(err) // Handle potential errors from crypto/rand
		}
		index := n.Int64()
		accountNumber[i] = validChars[index]
	}

	return string(accountNumber)
}
