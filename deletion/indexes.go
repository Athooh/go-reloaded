package delete

import (
	"fmt" // Importing the fmt package for printing messages
)

// RemoveAtIndex removes an element at the specified index from a slice of strings.
// It returns the modified slice after removal.
func RemoveAtIndex(s []string, i int) []string {
	// Checking if the index is out of range
	if i <= 0 || i > len(s) {
		fmt.Println("Index Out of range.") // Printing an error message if the index is invalid
	} else {
		// Removing the element at index i from the slice
		s = append(s[:i], s[i+1:]...)
	}
	return s // Returning the modified slice
}

// RemoveAtIndex2 removes an element at the specified index from a slice of strings
// and the next element after that index.
// It returns the modified slice after removal.
func RemoveAtIndex2(s []string, i int) []string {
	// Checking if the index is out of range
	if i <= 0 || i > len(s) {
		fmt.Println("Index Out of range.") // Printing an error message if the index is invalid
	} else {
		// Removing the element at index i and the next element from the slice
		s = append(s[:i], s[i+2:]...)
	}
	return s // Returning the modified slice
}
