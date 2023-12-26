package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestGeneratePassword(t *testing.T) {
	// test table
	var tests = []struct {
		name           string
		length         int
		expectedLength int
	}{
		{"test if lenght is correct", 100, 100},
		{"test for highest length", 1, 1},
	}

	// loop through test table
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := generatePassword(tt.length)
			assert.NilError(t, err)
			assert.Equal(t, len(actual), tt.expectedLength)
		})
	}
}

func TestGeneratePasswordForErrors(t *testing.T) {
	// test table
	var tests = []struct {
		name          string
		inputLength   int
		expectedError string
	}{
		{"error thrown with length of 0", 0, "length is is too small"},
		{"error thrown with length of 0", -100, "length is is too small"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := generatePassword(tt.inputLength)
			assert.Error(t, err, tt.expectedError)
		})
	}
}
