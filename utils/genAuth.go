package utils

import (
	"encoding/hex"
	"math/rand"
)

func GenerateAuthAcess() (string, error) {
	bs := make([]byte, 32)

	if _, err := rand.Read(bs); err != nil {
		return "", err
	}

	hex := hex.EncodeToString(bs)

	return hex, nil
}
