package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/uuid"
)

// GenerateToken 生成明文 token + hash
func GenerateToken() (string, string) {
	rawToken := uuid.NewString()
	hashed := sha256.Sum256([]byte(rawToken))
	return rawToken, hex.EncodeToString(hashed[:])
}

// HashToken 对用户传回的 token 进行 hash
func HashToken(token string) string {
	hashed := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hashed[:])
}
