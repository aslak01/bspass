package main

import (
	"math"
	"unicode"
)

func calculateCharsetSize(password string) int {
	hasLower := false
	hasUpper := false
	hasDigit := false
	hasSymbol := false

	for _, r := range password {
		if unicode.IsLower(r) {
			hasLower = true
		} else if unicode.IsUpper(r) {
			hasUpper = true
		} else if unicode.IsDigit(r) {
			hasDigit = true
		} else if unicode.IsSymbol(r) || unicode.IsPunct(r) {
			hasSymbol = true
		}
	}

	charsetSize := 0
	if hasLower {
		charsetSize += 26
	}
	if hasUpper {
		charsetSize += 26
	}
	if hasDigit {
		charsetSize += 10
	}
	if hasSymbol {
		// approximation
		charsetSize += 32
	}
	return charsetSize
}

func CalculatePasswordEntropy(password string) float64 {
	length := len(password)
	charsetSize := calculateCharsetSize(password)

	if length == 0 || charsetSize == 0 {
		return 0.0
	}

	entropy := float64(length) * math.Log2(float64(charsetSize))
	return entropy
}
