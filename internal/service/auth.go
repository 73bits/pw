package service

import (
	"crypto/pbkdf2"
	"crypto/sha256"
)

func DeriveKey(password string, salt []byte) ([]byte, error) {
	return pbkdf2.Key(sha256.New, password, salt, 100000, 32)
}
