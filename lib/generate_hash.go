package lib

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateHashWithSha256(url string) string {
	algo := sha256.New()
	algo.Write([]byte(url))
	return hex.EncodeToString(algo.Sum(nil))[:8]
}
