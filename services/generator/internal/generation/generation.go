package generation

import (
	"errors"
	"math/rand"
)

func RandNumber(minimum int, maximum int) int {
	num := rand.Intn(maximum-minimum) + minimum
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
