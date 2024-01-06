package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestRandNumber(t *testing.T) {
	var tests = []struct {
		name     string
		inputMin int
		inputMax int
	}{
		{"value is in range", 1, 100},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := randNumber(test.inputMin, test.inputMax)
			assert.Check(t, actual >= test.inputMin && actual <= test.inputMax)
		})
	}
}

func TestGeneratePassword(t *testing.T) {
	var tests = []struct {
		name           string
		length         int
		expectedLength int
	}{
		{"test if lenght is correct", 100, 100},
		{"test length of one", 1, 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := generatePassword(test.length)
			assert.NilError(t, err)
			assert.Equal(t, len(actual), test.expectedLength)
		})
	}
}

func TestGeneratePasswordForErrors(t *testing.T) {
	var tests = []struct {
		name          string
		inputLength   int
		expectedError string
	}{
		{"error thrown with length of 0", 0, "length is too small"},
		{"error thrown with negative length", -100, "length is too small"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := generatePassword(test.inputLength)
			assert.Error(t, err, test.expectedError)
		})
	}
}
