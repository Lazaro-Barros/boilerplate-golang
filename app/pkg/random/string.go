package random

import (
	"math/rand"
	"strings"
	"time"
)

func String(length int) string {
	// Inicializa o gerador de números aleatórios com a semente atual
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Define o conjunto de caracteres que podem ser usados na string
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var sb strings.Builder
	for i := 0; i < length; i++ {
		// Gera um índice aleatório dentro do charset
		randomIndex := rand.Intn(len(charset))
		sb.WriteByte(charset[randomIndex])
	}

	return sb.String()
}
