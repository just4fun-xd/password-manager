package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	EncryptionKey = os.Getenv("ENCRYPTION_KEY")
)

const (
	// Настройки сервера
	ServerPort = ":8080"

	// Символы для генерации пароля
	LetterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	DigitBytes  = "0123456789"
	SymbolBytes = "!@#$%^&*()-_=+[]{}<>?"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Файл .env не найден. Используются стандартные переменные окружения.")
	}

	EncryptionKey = os.Getenv("ENCRYPTION_KEY")
	if len(EncryptionKey) != 16 && len(EncryptionKey) != 24 && len(EncryptionKey) != 32 {
		log.Fatalf("Неверная длина ключа шифрования: %d. Ключ должен быть 16, 24 или 32 байта.", len(EncryptionKey))
	}
}
