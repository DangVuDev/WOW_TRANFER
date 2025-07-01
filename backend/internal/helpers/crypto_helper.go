package helpers

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/hex"
    "errors"
    "io"
)

func EncryptPrivateKey(privateKey, key string) (string, error) {
    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        return "", err
    }
    ciphertext := make([]byte, aes.BlockSize+len(privateKey))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }
    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(privateKey))
    return hex.EncodeToString(ciphertext), nil
}

func DecryptPrivateKey(encryptedKey, key string) (string, error) {
    ciphertext, _ := hex.DecodeString(encryptedKey)
    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        return "", err
    }
    if len(ciphertext) < aes.BlockSize {
        return "", errors.New("dữ liệu mã hóa quá ngắn")
    }
    iv := ciphertext[:aes.BlockSize]
    ciphertext = ciphertext[aes.BlockSize:]
    stream := cipher.NewCFBDecrypter(block, iv)
    stream.XORKeyStream(ciphertext, ciphertext)
    return string(ciphertext), nil
}