package utils

import (
	"encoding/hex"
	"os"

	"crypto/md5"
	"crypto/rand"
)

func GenerateAuthAcess() (string, error) {
	f, err := os.OpenFile("auth.md5", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return "", err
	}

	auth := make([]byte, 32)

	if _, err := rand.Read(auth); err != nil {
		return "", err
	}

	h := md5.New()
	h.Write(auth)

	f.Write(h.Sum(nil))

	return hex.EncodeToString(h.Sum(nil)), nil
}
