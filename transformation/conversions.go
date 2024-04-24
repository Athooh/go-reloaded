package transform

import (
	"fmt"
	delete "go-reloaded/deletion" // Importing deletion package
	"strconv"
)

// HexToDec converts hexadecimal strings to decimal strings
func HexToDec(s []string) []string {
	var result bool // Variable to track if an error occurred during conversion
	for i, v := range s {
		if v == "(hex)" && i != 0 {
			hexVal, err := strconv.ParseInt(s[i-1], 16, 64) // Convert hex string to decimal integer
			if err != nil {
				result = true
				fmt.Println("Error!", err) // Print error message if conversion fails
			}
			s[i-1] = strconv.Itoa(int(hexVal)) // Replace the original value with the converted decimal value
		}
	}
	if !result {
		for idx, val := range s {
			if val == "(hex)" {
				s = delete.RemoveAtIndex(s, idx) // Remove "(hex)" markers from the slice
			}
		}
	}
	return s // Return the slice with converted decimal values
}

// BinToDec converts binary strings to decimal strings
func BinToDec(s []string) []string {
	var result bool // Variable to track if an error occurred during conversion
	for i, v := range s {
		if v == "(bin)" && i != 0 {
			binVal, err := strconv.ParseInt(s[i-1], 2, 64) // Convert binary string to decimal integer
			if err != nil {
				result = true
				fmt.Println("Error!", err) // Print error message if conversion fails
			}
			s[i-1] = strconv.Itoa(int(binVal)) // Replace the original value with the converted decimal value
		}
	}
	if !result {
		for idx, val := range s {
			if val == "(bin)" {
				s = delete.RemoveAtIndex(s, idx) // Remove "(bin)" markers from the slice
			}
		}
	}
	return s // Return the slice with converted decimal values
}
