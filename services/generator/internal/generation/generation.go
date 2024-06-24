package generation

import (
	"errors"
	"math/rand"
)

func RandNumber(min int, max int) int {
	num := rand.Intn(max-min) + min
	return num
}

func GenerateRandString(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("length is too small")
	}

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	t := make([]byte, length)

	for i := range t {
		t[i] = charset[rand.Intn(len(charset))]
	}

	return string(t), nil
}
