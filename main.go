package main

import (
	"fmt"
	"math"

	"unicode"
)

func main() {
	input := "www.google.com"
	result := BitEntropy(input)
	fmt.Printf("Entropy of '%s' is: %d", input, result)
}

// Arbitrary bit entropy classification
// < 70 bits: Weak
// 70-100+ bits: Acceptable
// 100+ bits: Strong
func BitEntropy(value string) (bits int) {

	totalCharacterSet := 0

	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSymbol := false

	for _, r := range value {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasDigit = true
		default:
			hasSymbol = true
		}
	}

	if hasUpper {
		totalCharacterSet += 26
	}
	if hasLower {
		totalCharacterSet += 26
	}
	if hasDigit {
		totalCharacterSet += 9
	}
	if hasSymbol {
		totalCharacterSet += 23
	}
	//23 assumption, some older  sites  prohibit certain special characters or spaces due to improper data handling practices.

	//H≈Lxlog2(R)
	return int(float64(len(value)) * math.Log2(float64(totalCharacterSet)))
}
