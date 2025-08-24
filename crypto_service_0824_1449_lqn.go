// 代码生成时间: 2025-08-24 14:49:54
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "io"
    "log"
    "os"
)

// CryptoService is a struct that contains necessary information for encryption and decryption
type CryptoService struct {
    key []byte
}

// NewCryptoService creates a new instance of CryptoService
// It takes a passphrase as a parameter and initializes the service with the key
func NewCryptoService(passphrase string) *CryptoService {
    key := []byte(passphrase)
    return &CryptoService{key: key}
}

// Encrypt takes a plaintext string and encrypts it using AES-256
func (cs *CryptoService) Encrypt(plaintext string) (string, error) {
    block, err := aes.NewCipher(cs.key)
    if err != nil {
        return "", err
    }

    // PKCS7 padding
    blockSize := block.BlockSize()
    padding := blockSize - len(plaintext)%blockSize
    paddedPlaintext := plaintext + string(byte(padding))*padding

    // Encrypt the padded plaintext
    ciphertext := make([]byte, aes.BlockSize+len(paddedPlaintext))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }
    mode := cipher.NewCBCEncrypter(block, iv)
    mode.CryptBlocks(ciphertext[aes.BlockSize:], []byte(paddedPlaintext))

    // Return the base64-encoded ciphertext
    return hex.EncodeToString(ciphertext), nil
}

// Decrypt takes an encrypted string and decrypts it using AES-256
func (cs *CryptoService) Decrypt(ciphertext string) (string, error) {
    block, err := aes.NewCipher(cs.key)
    if err != nil {
        return "", err
    }

    cipherData, err := hex.DecodeString(ciphertext)
    if err != nil {
        return "", err
    }

    if len(cipherData) < aes.BlockSize {
        return "", fmt.Errorf("ciphertext too short")
    }
    iv := cipherData[:aes.BlockSize]
    cipherText := cipherData[aes.BlockSize:]
    mode := cipher.NewCBCDecrypter(block, iv)
    mode.CryptBlocks(cipherText, cipherText)

    // Unpad the plaintext
    unpaddedPlaintext := make([]byte, len(cipherText))
    copy(unpaddedPlaintext, cipherText)
    unpaddedPlaintext = pkcs7Unpad(unpaddedPlaintext)

    return string(unpaddedPlaintext), nil
}

// pkcs7Unpad removes PKCS7 padding from the ciphertext
func pkcs7Unpad(p []byte) []byte {
    length := len(p)
    unpadding := int(p[length-1])
    return p[:(length - unpadding)]
}

func main() {
    passphrase := "mysecretkey"
    service := NewCryptoService(passphrase)

    plaintext := "Hello, world!"
    encrypted, err := service.Encrypt(plaintext)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Encrypted: %s
", encrypted)

    decrypted, err := service.Decrypt(encrypted)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Decrypted: %s
", decrypted)
}
