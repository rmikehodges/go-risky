package util

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
)

func GenerateRandomToken() (string, error) {
	token := make([]byte, 64) // 64 bytes for sha512
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	hash := sha512.New()
	hash.Write(token)
	hashedToken := hash.Sum(nil)
	return hex.EncodeToString(hashedToken), nil
}
