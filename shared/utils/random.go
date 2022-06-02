package utils

import (
	"crypto/rand"
	"encoding/hex"
)

// RandomHex returns a random hexadecimal string of n length
func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes)[0:n], nil
}
