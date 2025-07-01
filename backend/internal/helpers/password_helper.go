package helpers

import (
    "crypto/sha256"
    "encoding/hex"
)

func HashPassword(password string) (string, error) {
    hash := sha256.Sum256([]byte(password))
    return hex.EncodeToString(hash[:]), nil
}

func CheckPassword(password, hashedPassword string) bool {
    hash, _ := HashPassword(password)
    return hash == hashedPassword
}