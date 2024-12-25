package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"log"
	mrand "math/rand"
	"password-manager/config"
	"time"
)

var encryptionKey = []byte(config.EncryptionKey)

func GeneratePassword(length int, useDigits, useSymbols bool) string {
	generator := mrand.New(mrand.NewSource(time.Now().UnixNano()))
	charset := config.LetterBytes
	if useDigits {
		charset += config.DigitBytes
	}
	if useSymbols {
		charset += config.SymbolBytes
	}

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[generator.Intn(len(charset))]
	}
	return string(result)
}

func Encrypt(text string) (string, error) {
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		log.Println("Ошибка создания шифра AES:", err)
		return "", err
	}

	plaintext := []byte(text)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	// Генерация случайного IV
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Println("Ошибка генерации IV:", err)
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	encrypted := base64.StdEncoding.EncodeToString(ciphertext)
	return encrypted, nil
}

func Decrypt(encrypted string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil
}
