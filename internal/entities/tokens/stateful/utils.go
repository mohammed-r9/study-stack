package stateful

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
)

func hashFromPlainText(plainText string) string {
	hash := sha256.Sum256([]byte(plainText))
	return hex.EncodeToString(hash[:])
}

// takes a recievedToken (plain text) and compares it to the stored token hash in the db
func CompareOpaqueTokens(recievedToken string, storedToken string) bool {
	receivedHash := hashFromPlainText(recievedToken)

	return subtle.ConstantTimeCompare([]byte(receivedHash), []byte(storedToken)) == 1
}

func newOpaque() (string, error) {
	randomBytes := make([]byte, 32)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(randomBytes), nil
}
