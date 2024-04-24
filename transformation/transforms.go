package transform

import (
	"fmt"
	delete "go-reloaded/deletion" // Importing the deletion package
	"strconv"
	"strings"
	"unicode"
)

// Lowercase converts strings to lowercase based on certain criteria.
func Lowercase(s []string) []string {
	var result bool // Variable to track errors
	for i, v := range s {
		if v == "(low)" && i != 0 { // Check if the token is (low) and not the first element
			var word string            // Placeholder for the modified word
			for _, c := range s[i-1] { // Iterate over the previous string
				c = unicode.ToLower(c) // Convert character to lowercase
				word += string(c)      // Concatenate the characters
			}
			s[i-1] = word                  // Update the previous string with the lowercase version
			s = delete.RemoveAtIndex(s, i) // Remove the token (low)
		} else if v == "(low," && i != 0 { // Check if the token is (low,) and not the first element
			strNumb := strings.TrimSuffix(s[i+1], ")") // Extract the number from the next string
			numb, err := strconv.Atoi(strNumb)         // Convert string to integer
			if err != nil {                            // Check for conversion error
				result = true              // Set result to true indicating an error
				fmt.Println("Error!", err) // Print the error
			}
			// Check for invalid number
			if numb > i {
				fmt.Printf("Error!: Number (%d) greater than length of string upto index of (low,: %d \n", numb, i)
				result = true
				break
			} else if numb < 0 {
				fmt.Println("Error!: Number provided is negative.")
				result = true
				break
			}
			// Convert selected strings to lowercase
			for x := i - numb; x < i; x++ {
				s[x] = strings.ToLower(s[x])
			}
		}
	}
	// Remove any remaining tokens (low,)
	if !result {
		for idx, v := range s {
			if v == "(low," {
				s = delete.RemoveAtIndex2(s, idx)
			}
		}
	}
	return s // Return the modified slice
}

// Uppercase converts strings to uppercase based on certain criteria.
func Uppercase(s []string) []string {
	var result bool // Variable to track errors
	for i, v := range s {
		if v == "(up)" && i != 0 { // Check if the token is (up) and not the first element
			var word string            // Placeholder for the modified word
			for _, c := range s[i-1] { // Iterate over the previous string
				c = unicode.ToUpper(c) // Convert character to uppercase
				word += string(c)      // Concatenate the characters
			}
			s[i-1] = word                  // Update the previous string with the uppercase version
			s = delete.RemoveAtIndex(s, i) // Remove the token (up)
		} else if v == "(up," && i != 0 { // Check if the token is (up,) and not the first element
			strNumb := strings.TrimSuffix(s[i+1], ")") // Extract the number from the next string
			numb, err := strconv.Atoi(strNumb)         // Convert string to integer
			if err != nil {                            // Check for conversion error
				result = true // Set result to true indicating an error
				fmt.Println("Error! Value provided is not a valid number")
			}
			// Check for invalid number
			if numb > i {
				fmt.Printf("Error!: Number (%d) greater than length of string upto index of (up,: %d \n", numb, i)
				result = true
				break
			} else if numb < 0 {
				fmt.Println("Error!: Number provided is negative.")
				result = true
				break
			}
			// Convert selected strings to uppercase
			for x := i - numb; x < i; x++ {
				s[x] = strings.ToUpper(s[x])
			}
		}
	}
	// Remove any remaining tokens (up,)
	if !result {
		for idx, v := range s {
			if v == "(up," {
				s = delete.RemoveAtIndex2(s, idx)
			}
		}
	}
	return s // Return the modified slice
}

// Capitalize capitalizes the first letter of strings based on certain criteria.
func Capitalize(s []string) []string {
	var result bool // Variable to track errors
	for i, v := range s {
		if v == "(cap)" && i != 0 { // Check if the token is (cap) and not the first element
			var word string              // Placeholder for the modified word
			for idx, c := range s[i-1] { // Iterate over the previous string
				if idx == 0 {
					c = unicode.ToTitle(c) // Capitalize the first character
					word += string(c)      // Concatenate the characters
				} else {
					word += string(c) // Concatenate the characters
				}
			}
			s[i-1] = word                  // Update the previous string with the capitalized version
			s = delete.RemoveAtIndex(s, i) // Remove the token (cap)
		} else if v == "(cap," && i != 0 { // Check if the token is (cap,) and not the first element
			strNumb := strings.TrimSuffix(s[i+1], ")") // Extract the number from the next string
			numb, err := strconv.Atoi(strNumb)         // Convert string to integer
			if err != nil {                            // Check for conversion error
				result = true // Set result to true indicating an error
				fmt.Println("Error! Value provided is not a valid number")
			}
			// Check for invalid number
			if numb > i {
				fmt.Printf("Error!: Number (%d) greater than length of string upto index of (cap,: %d \n", numb, i)
				result = true
				break
			} else if numb < 0 {
				fmt.Println("Error!: Number provided is negative.")
				result = true
				break
			}
			// Capitalize the first letter of selected strings
			for x := i - numb; x < i; x++ {
				var word string              // Placeholder for the modified word
				for idx, val := range s[x] { // Iterate over the string
					if idx == 0 {
						val = unicode.ToTitle(val) // Capitalize the first character
						word += string(val)        // Concatenate the characters
					} else {
						word += string(val) // Concatenate the characters
					}
				}
				s[x] = word // Update the string with the capitalized version
			}
		}
	}
	// Remove any remaining tokens (cap,)
	if !result {
		for idx, v := range s {
			if v == "(cap," {
				s = delete.RemoveAtIndex2(s, idx)
			}
		}
	}
	return s // Return the modified slice
}
