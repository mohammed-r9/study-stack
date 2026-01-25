package password

import (
	"crypto/rand"
	"crypto/subtle"
	"errors"
	"fmt"
	appErrors "study-stack/internal/shared/app_errors"

	"golang.org/x/crypto/argon2"
)

type Password struct {
	Hash []byte
	Salt []byte
}

const (
	t       = 3
	memory  = 64 * 1024
	threads = 4
	keyLen  = 32
	saltLen = 16
)

func generateSalt(length int) ([]byte, error) {
	if length <= 0 {
		return nil, fmt.Errorf("invalid salt length: %d", length)
	}

	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, fmt.Errorf("failed to generate salt: %w", err)
	}
	return salt, nil
}

func (p *Password) Set(plainText string) error {
	salt, err := generateSalt(saltLen)
	if err != nil {
		return err
	}
	hash := argon2.IDKey([]byte(plainText), salt, t, memory, threads, keyLen)
	p.Salt = salt
	p.Hash = hash

	return nil
}

func (p *Password) Matches(plainText string) error {
	if p.Salt == nil || p.Hash == nil {
		return errors.New("missing salt or hash")
	}

	hash := argon2.IDKey([]byte(plainText), p.Salt, t, memory, threads, keyLen)
	if subtle.ConstantTimeCompare(hash, p.Hash) == 1 {
		return nil
	}
	return appErrors.PasswordMismatch
}
