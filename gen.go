package main

import (
	"errors"
	"math/rand"
)

func randNumber(min int, max int) int {
	num := rand.Intn(max-min) + min
	return num
}

func generatePassword(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("length is too small")
	}

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	password := make([]byte, length)

	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	return string(password), nil
}
