package main

import (
	"fmt"
	"strings"
)

func SlurredTalk(words *string) {
	// Mengubah huruf yang sulit dieja menjadi 'L'
	result := strings.ReplaceAll(*words, "S", "L")
	result = strings.ReplaceAll(result, "s", "l")
	result = strings.ReplaceAll(result, "R", "L")
	result = strings.ReplaceAll(result, "r", "l")
	result = strings.ReplaceAll(result, "Z", "L")
	result = strings.ReplaceAll(result, "z", "l")
	
	// Mengupdate pointer dengan string yang sudah dimodifikasi
	*words = result
}

func main() {
	// Menguji beberapa test case
	testCases := []string{
		"Steven",
		"Saya Steven",
		"Saya Steven, saya suka menggoreng telur dan suka hewan zebra",
		"",
	}

	for _, words := range testCases {
		fmt.Printf("Input: \"%s\"\n", words)
		SlurredTalk(&words)
		fmt.Printf("Output: \"%s\"\n\n", words)
	}
}
