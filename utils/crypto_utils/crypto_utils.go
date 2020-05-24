package crypto_utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetSHA256(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
