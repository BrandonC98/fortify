package main

import (
	"errors"
	"math/rand"
	"time"
)

func generatePassword(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("length is is too small")
	}

	rand.Seed(time.Now().UnixNano())

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)

	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result), nil
}
